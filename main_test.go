package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestSplitLinesPreserveEndings(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected [][]byte
	}{
		{"empty", []byte{}, nil},
		{"single line no newline", []byte("hello"), [][]byte{[]byte("hello\n")}},
		{"single line with LF", []byte("hello\n"), [][]byte{[]byte("hello\n")}},
		{"multiple lines LF", []byte("a\nb\nc\n"), [][]byte{[]byte("a\n"), []byte("b\n"), []byte("c\n")}},
		{"multiple lines no final newline", []byte("a\nb\nc"), [][]byte{[]byte("a\n"), []byte("b\n"), []byte("c\n")}},
		{"CRLF endings", []byte("a\r\nb\r\n"), [][]byte{[]byte("a\r\n"), []byte("b\r\n")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := splitLinesPreserveEndings(tt.input)
			if len(result) != len(tt.expected) {
				t.Errorf("splitLinesPreserveEndings() length = %d, want %d", len(result), len(tt.expected))
				return
			}
			for i, line := range result {
				if string(line) != string(tt.expected[i]) {
					t.Errorf("splitLinesPreserveEndings()[%d] = %q, want %q", i, string(line), string(tt.expected[i]))
				}
			}
		})
	}
}

func TestReverse(t *testing.T) {
	// Test with string slice
	strs := []string{"a", "b", "c"}
	reverse(strs)
	expected := []string{"c", "b", "a"}
	for i, s := range strs {
		if s != expected[i] {
			t.Errorf("reverse(strings) = %v, want %v", strs, expected)
			break
		}
	}

	// Test with byte slice slice  
	bytes := [][]byte{[]byte("first"), []byte("second")}
	reverse(bytes)
	if string(bytes[0]) != "second" || string(bytes[1]) != "first" {
		t.Errorf("reverse([][]byte) failed")
	}
}

func TestStripLastBck(t *testing.T) {
	tests := []struct {
		input, expected string
	}{
		{"file.txt.bck", "file.txt"},
		{"file.txt", "file.txt"},
		{"file.bck.bck", "file.bck"},
		{"/path/to/file.txt.BCK", "/path/to/file.txt"},
		{"just.bck", "just"},
	}

	for _, tt := range tests {
		if result := stripLastBck(tt.input); result != tt.expected {
			t.Errorf("stripLastBck(%q) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}

func TestEncodeDecode(t *testing.T) {
	tempDir := t.TempDir()
	
	tests := []struct {
		name string
		content string
	}{
		{"with newline", "line1\nline2\nline3\n"},
		{"no newline", "line1\nline2\nline3"},
		{"single line with newline", "hello world\n"},
		{"single line no newline", "hello world"},
		{"empty", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create test file
			testFile := filepath.Join(tempDir, tt.name+".txt")
			if err := os.WriteFile(testFile, []byte(tt.content), 0644); err != nil {
				t.Fatal(err)
			}

			// Encode
			if err := encode(testFile); err != nil {
				t.Fatalf("encode failed: %v", err)
			}

			// Check .bck file exists and contains marker if needed
			bckFile := testFile + ".bck"
			bckContent, err := os.ReadFile(bckFile)
			if err != nil {
				t.Fatal("failed to read .bck file")
			}

			// Check marker presence
			hasMarker := strings.HasPrefix(string(bckContent), "##BCKL.NNL##\n")
			shouldHaveMarker := len(tt.content) > 0 && !strings.HasSuffix(tt.content, "\n")
			if hasMarker != shouldHaveMarker {
				t.Errorf("marker presence mismatch: has=%v, should=%v", hasMarker, shouldHaveMarker)
			}

			// Decode (to different name to avoid overwrite prompt)
			os.Remove(testFile) // Remove original so decode creates clean copy
			if err := decode(bckFile); err != nil {
				t.Fatalf("decode failed: %v", err)
			}

			// Read decoded content
			decoded, err := os.ReadFile(testFile)
			if err != nil {
				t.Fatal(err)
			}

			// With the marker feature, content should round-trip perfectly
			if string(decoded) != tt.content {
				t.Errorf("encode/decode cycle failed:\noriginal: %q\ndecoded:  %q", 
					tt.content, string(decoded))
			}
		})
	}
}

func TestFileExists(t *testing.T) {
	tempDir := t.TempDir()
	
	// Create a test file
	testFile := filepath.Join(tempDir, "exists.txt")
	os.WriteFile(testFile, []byte("test"), 0644)
	
	if !fileExists(testFile) {
		t.Error("fileExists() should return true for existing file")
	}
	
	if fileExists(filepath.Join(tempDir, "nonexistent.txt")) {
		t.Error("fileExists() should return false for nonexistent file")
	}
}