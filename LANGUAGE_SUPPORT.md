# Adding Language Support to backlang

This guide explains how to add support for new programming languages to the backlang router.

## Quick Reference

To add a new language, simply add a new `Language` struct to the `getSupportedLanguages()` function in `router.go`.

## Language Struct Fields

```go
type Language struct {
    Name       string    // Display name (e.g., "Python", "JavaScript")
    Extensions []string  // File extensions (e.g., []string{".py", ".pyw"})
    Shebangs   []string  // Shebang patterns to match
    Command    string    // Command to execute (e.g., "python3", "node")
    Args       []string  // Default arguments before filename
}
```

## Detection Priority

1. **Shebang first** - More specific than extension
2. **File extension** - Fallback if no shebang or shebang not recognized

## Examples

### Python (Current Implementation)
```go
{
    Name:       "Python",
    Extensions: []string{".py"},
    Shebangs:   []string{"#!/usr/bin/env python3", "#!/usr/bin/python3", "#!/usr/bin/env python", "#!/usr/bin/python"},
    Command:    "python3",
    Args:       []string{},
}
```

### JavaScript/Node.js (Example)
```go
{
    Name:       "JavaScript",
    Extensions: []string{".js"},
    Shebangs:   []string{"#!/usr/bin/env node", "#!/usr/bin/node"},
    Command:    "node",
    Args:       []string{},
}
```

### Go (Compiled Language Example)
```go
{
    Name:       "Go",
    Extensions: []string{".go"},
    Shebangs:   []string{}, // Go doesn't typically use shebangs
    Command:    "go",
    Args:       []string{"run"}, // go run filename.go
}
```

### Rust (Complex Compiled Example)
```go
{
    Name:       "Rust",
    Extensions: []string{".rs"},
    Shebangs:   []string{},
    Command:    "rustc",
    Args:       []string{"--edition", "2021"}, // rustc --edition 2021 filename.rs && ./filename
}
```

Note: Compiled languages might need special handling in `executeFile()` for compile+run workflow.

## Adding a New Language

1. **Identify the language characteristics**:
   - Common file extensions
   - Typical shebang patterns
   - Command to run files
   - Any required arguments

2. **Add to `getSupportedLanguages()`**:
   ```go
   {
       Name:       "YourLanguage",
       Extensions: []string{".ext"},
       Shebangs:   []string{"#!/path/to/interpreter"},
       Command:    "interpreter",
       Args:       []string{}, // or required flags
   },
   ```

3. **Test with a sample .bck file**:
   ```bash
   # Create test file
   echo -e "print('world')\nprint('hello')" > test.py
   backlang encode test.py
   backlang run test.py.bck
   ```

## Special Cases

### Compiled Languages
For languages that need compilation before execution, you might need to modify `executeFile()` to:
1. Compile first
2. Run the compiled binary
3. Clean up binaries (optional)

### Languages with Complex Arguments
Some languages might need environment-specific arguments or flags. Add them to the `Args` field.

### Multiple Extensions
Languages with multiple file extensions (e.g., Python: `.py`, `.pyw`) can list all in the `Extensions` slice.

## Testing New Languages

1. Create a simple test file in the target language
2. Encode it: `backlang encode test.ext`
3. Run it: `backlang run test.ext.bck`
4. Verify both the decoding and execution work correctly

## Future Improvements

- **Auto-detection improvements**: Content-based detection for files without shebangs
- **Compiler caching**: Cache compiled binaries for repeated runs
- **Arguments passing**: Allow passing arguments to the executed program
- **Environment handling**: Custom environment variables per language