package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Language represents a supported programming language
type Language struct {
	Name       string
	Extensions []string
	Shebangs   []string
	Command    string
	Args       []string
}

// getSupportedLanguages returns the list of supported languages
func getSupportedLanguages() []Language {
	return []Language{
		{
			Name:       "Python",
			Extensions: []string{".py"},
			Shebangs:   []string{"#!/usr/bin/env python3", "#!/usr/bin/python3", "#!/usr/bin/env python", "#!/usr/bin/python"},
			Command:    "python3",
			Args:       []string{}, // Will append filename
		},
		// Future languages can be added here
		// {
		//     Name:       "JavaScript",
		//     Extensions: []string{".js"},
		//     Shebangs:   []string{"#!/usr/bin/env node", "#!/usr/bin/node"},
		//     Command:    "node",
		//     Args:       []string{},
		// },
	}
}

// run decodes a .bck file and executes it with the appropriate interpreter
func run(inPath string) error {
	// Validate input is a .bck file
	if !strings.HasSuffix(strings.ToLower(inPath), ".bck") {
		return fmt.Errorf("Error: run command only accepts .bck files")
	}

	// Decode the file
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

	// Create output file path (remove .bck extension)
	outPath := stripLastBck(inPath)
	
	// Handle file collisions with numbering
	if fileExists(outPath) {
		outPath = nextAvailableName(outPath)
	}

	// Write decoded content
	if err := os.WriteFile(outPath, join(lines), 0o666); err != nil {
		return wrapPathErr(err, outPath)
	}

	fmt.Printf("Decoded '%s' → '%s'\n", filepath.Base(inPath), filepath.Base(outPath))

	// Detect language and run
	lang, err := detectLanguage(outPath)
	if err != nil {
		return err
	}

	fmt.Printf("Detected %s, running with %s...\n", lang.Name, lang.Command)
	return executeFile(lang, outPath)
}

// detectLanguage determines the programming language based on shebang and extension
func detectLanguage(filePath string) (*Language, error) {
	languages := getSupportedLanguages()
	
	// Read first line to check for shebang
	file, err := os.Open(filePath)
	if err != nil {
		return nil, wrapPathErr(err, filePath)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var firstLine string
	if scanner.Scan() {
		firstLine = strings.TrimSpace(scanner.Text())
	}

	// Check shebang first (more specific)
	if strings.HasPrefix(firstLine, "#!") {
		for _, lang := range languages {
			for _, shebang := range lang.Shebangs {
				if strings.HasPrefix(firstLine, shebang) {
					return &lang, nil
				}
			}
		}
	}

	// Check file extension
	ext := strings.ToLower(filepath.Ext(filePath))
	for _, lang := range languages {
		for _, langExt := range lang.Extensions {
			if ext == langExt {
				return &lang, nil
			}
		}
	}

	return nil, fmt.Errorf("Error: No interpreter found for '%s'", filepath.Base(filePath))
}

// executeFile runs the decoded file with the appropriate interpreter
func executeFile(lang *Language, filePath string) error {
	// Prepare command
	args := append(lang.Args, filePath)
	cmd := exec.Command(lang.Command, args...)
	
	// Connect stdin, stdout, stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	// Execute
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("Error: Failed to execute with %s: %v", lang.Command, err)
	}
	
	return nil
}