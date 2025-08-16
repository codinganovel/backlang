package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const usageText = "Usage: backlang <encode|decode|run> <file>\n"

func main() {
	if len(os.Args) != 3 {
		fmt.Fprint(os.Stderr, usageText)
		os.Exit(2)
	}

	cmd := os.Args[1]
	inPath := os.Args[2]

	switch cmd {
	case "encode":
		if err := encode(inPath); err != nil {
			printErr(err)
			os.Exit(1)
		}
	case "decode":
		if !strings.HasSuffix(strings.ToLower(inPath), ".bck") {
			fmt.Fprintln(os.Stderr, "Error: decode command only accepts .bck files")
			os.Exit(2)
		}
		if err := decode(inPath); err != nil {
			printErr(err)
			os.Exit(1)
		}
	case "run":
		if err := run(inPath); err != nil {
			printErr(err)
			os.Exit(1)
		}
	default:
		fmt.Fprint(os.Stderr, usageText)
		os.Exit(2)
	}
}

func encode(inPath string) error {
	data, err := os.ReadFile(inPath)
	if err != nil {
		return wrapPathErr(err, inPath)
	}

	// Check if original file lacks trailing newline
	hasTrailingNewline := len(data) > 0 && (data[len(data)-1] == '\n')
	
	lines := splitLinesPreserveEndings(data) // each slice includes its original newline (if any)
	reverse(lines)
	
	// Add marker if original had no trailing newline
	if !hasTrailingNewline && len(data) > 0 {
		marker := []byte("##BCKL.NNL##\n")
		lines = append([][]byte{marker}, lines...)
	}

	outPath := inPath + ".bck"
	if err := os.WriteFile(outPath, join(lines), 0o666); err != nil {
		return wrapPathErr(err, outPath)
	}

	fmt.Printf("Encoded '%s' → '%s'\n", filepath.Base(inPath), filepath.Base(outPath))
	return nil
}

func decode(inPath string) error {
	data, err := os.ReadFile(inPath)
	if err != nil {
		return wrapPathErr(err, inPath)
	}

	lines := splitLinesPreserveEndings(data)
	
	// Check for marker at the beginning
	hasMarker := false
	if len(lines) > 0 && string(lines[0]) == "##BCKL.NNL##\n" {
		hasMarker = true
		lines = lines[1:] // Remove marker
	}
	
	reverse(lines)
	
	// If marker was present, remove the trailing newline we added during encode
	if hasMarker && len(lines) > 0 {
		lastLine := lines[len(lines)-1]
		if len(lastLine) > 0 && lastLine[len(lastLine)-1] == '\n' {
			lines[len(lines)-1] = lastLine[:len(lastLine)-1]
		}
	}

	outPath := stripLastBck(inPath)
	// If target exists, prompt and either overwrite or auto-increment.
	if fileExists(outPath) {
		overwrite, err := promptOverwrite(outPath)
		if err != nil {
			return err
		}
		if !overwrite {
			outPath = nextAvailableName(outPath)
		}
	}

	if err := os.WriteFile(outPath, join(lines), 0o666); err != nil {
		return wrapPathErr(err, outPath)
	}

	fmt.Printf("Decoded '%s' → '%s'\n", filepath.Base(inPath), filepath.Base(outPath))
	return nil
}

// --- helpers ---

// splitLinesPreserveEndings splits into records where each element includes its original
// newline sequence (LF or CRLF) if present. The last element may not end with a newline.
func splitLinesPreserveEndings(b []byte) [][]byte {
	var lines [][]byte
	start := 0
	for i := 0; i < len(b); i++ {
		if b[i] == '\n' {
			// include CR if present
			end := i + 1
			lines = append(lines, b[start:end])
			start = end
		}
	}
	if start < len(b) {
		// trailing line without newline - add a newline to prevent concatenation
		lastLine := make([]byte, len(b[start:])+1)
		copy(lastLine, b[start:])
		lastLine[len(lastLine)-1] = '\n'
		lines = append(lines, lastLine)
	} else if len(b) > 0 && (b[len(b)-1] == '\n') {
		// file ends with newline: already included in last record; nothing to add
	}
	return lines
}

func reverse[T any](s []T) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func join(chunks [][]byte) []byte {
	if len(chunks) == 0 {
		return nil
	}
	total := 0
	for _, c := range chunks {
		total += len(c)
	}
	out := make([]byte, 0, total)
	for _, c := range chunks {
		out = append(out, c...)
	}
	return out
}

func stripLastBck(path string) string {
	// remove only the final ".bck" (case-insensitive)
	dir := filepath.Dir(path)
	base := filepath.Base(path)
	// find last ".bck" ignoring case
	idx := strings.LastIndex(strings.ToLower(base), ".bck")
	if idx >= 0 && idx == len(base)-4 {
		base = base[:idx]
	}
	if dir == "." {
		return base
	}
	return filepath.Join(dir, base)
}

func nextAvailableName(path string) string {
	dir := filepath.Dir(path)
	base := filepath.Base(path)
	ext := filepath.Ext(base)
	name := strings.TrimSuffix(base, ext)

	for i := 1; ; i++ {
		candidate := filepath.Join(dir, fmt.Sprintf("%s_%d%s", name, i, ext))
		if !fileExists(candidate) {
			return candidate
		}
	}
}

func fileExists(p string) bool {
	_, err := os.Stat(p)
	return err == nil
}

func promptOverwrite(target string) (bool, error) {
	fmt.Printf("File '%s' exists. Overwrite? (y/n): ", filepath.Base(target))
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil && !errors.Is(err, io.EOF) {
		return false, err
	}
	line = strings.TrimSpace(strings.ToLower(line))
	return line == "y" || line == "yes", nil
}

func wrapPathErr(err error, path string) error {
	if errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("Error: File '%s' not found", filepath.Base(path))
	}
	if errors.Is(err, os.ErrPermission) {
		return fmt.Errorf("Error: Permission denied accessing '%s'", filepath.Base(path))
	}
	// fallback with original message
	return err
}

func printErr(err error) {
	fmt.Fprintln(os.Stderr, err.Error())
}
