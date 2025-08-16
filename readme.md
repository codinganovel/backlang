# backlang - Finally, Programming Can Be Hard Enough

**The universal backwards programming language that makes every language harder to write.**

> *"Why make programming easier when you can make it unnecessarily difficult?"*

---

## 🎯 The Problem

Programming languages today are just too easy. You write code from top to bottom, functions flow logically, and variables make sense in context. Where's the challenge? Where's the mental gymnastics? Where's the suffering that builds character?

After decades of languages trying to be "readable" and "maintainable," we finally have a solution.

## ✨ The Solution

**backlang** - where every program is written backwards, line by line. Take any programming language, write it in reverse order, and suddenly you have the most challenging way to code ever invented.

Want to use Rust but feel like it's not quite painful enough? **Write it backwards.**  
Think Python is too simple? **Write it backwards.**  
Feel like JavaScript doesn't cause enough mental anguish? **Write it backwards.**

Finally, a "language" that works with every language and makes them all equally traumatic to write.

---

## 📦 Installation

you obviously only build it from source, do you know how to do that right? :

```bash
git clone https://github.com/codinganovel/backlang
cd backlang
go build -o backlang
```

---

## 🚀 Quick Start

### Encode: Make Any Language Backwards
```bash
# Take your beautiful, readable code
backlang encode hello.py

# Creates hello.py.bck - now it's properly backwards and unreadable
```

### Decode: Convert Back to Boring Normal Code
```bash
# When you need to actually run your backwards masterpiece
backlang decode hello.py.bck

# Creates hello.py - back to the boring, sensible version
```

### Run: Because Even Backwards Code Should Execute
```bash
# For the truly committed - decode AND execute in one command
backlang run hello.py.bck

# Automatically detects it's Python, decodes, and runs it
# Your backwards code is now actually executable. What have we done?
```

---

## 💡 Example: Python → backlang → Python

**Original Python (boring, readable):**
```python
print("Hello")
print("World")
print("!")
```

**After `backlang encode hello.py` (artistic, challenging):**
```python
print("!")
print("World") 
print("Hello")
```

**The workflow:**
1. Write your program backwards in any language
2. Save it as a `.bck` file 
3. Use `backlang decode` to get runnable code (or `backlang run` to decode and execute automatically)
4. Impress your colleagues with your dedication to unnecessary complexity

**What happens when you run it:**
```bash
$ backlang run hello.py.bck
Decoded 'hello.py.bck' → 'hello_1.py'
Detected Python, running with python3...
Hello
World
!
```

---

## 🎪 Features

- **Universal compatibility** - Works with Python, JavaScript, Rust, Go, C++, or any text-based language
- **Bidirectional translation** - Encode normal code to backlang, decode backlang to normal
- **Direct execution** - Decode and run backwards code in one command (Python, JavaScript, Bash supported)
- **100% reversible** - Perfect round-trip preservation of your original file, including trailing newline handling
- **Line-perfect preservation** - Every character, space, and tab exactly where you left it
- **File conflict protection** - Won't accidentally overwrite your backwards masterpieces

---

## 🤔 Use Cases

### For the Overengineering Enthusiast
```bash
# Make your JavaScript even more confusing
backlang encode app.js
# Edit app.js.bck (writing backwards, obviously)
backlang run app.js.bck  # Decode and execute in one go
```

### For the Code Golf Masochist  
```bash
# Challenge mode: write Rust backwards
backlang encode main.rs
# Now try to implement a linked list in reverse order
```

### For the Interview Preparation Sadist
```bash
# Practice writing algorithms backwards
# If you can implement quicksort in backlang, you can handle any interview
```

---

## 📚 Commands

| Command | What It Does | File Requirements |
|---------|--------------|-------------------|
| `backlang encode <file>` | Converts normal code to backlang | Any text file |
| `backlang decode <file>` | Converts backlang back to normal | Must be a `.bck` file |
| `backlang run <file>` | Decodes and executes backwards code | Must be a `.bck` file (Python, JS, Bash) |

### Advanced Workflows

```bash
# Double-backwards (why not?)
backlang encode script.py        # → script.py.bck
backlang encode script.py.bck    # → script.py.bck.bck
backlang decode script.py.bck.bck # → script.py.bck  
backlang decode script.py.bck    # → script.py (back to start)
```

---

## 🎨 Philosophy

In an industry obsessed with "developer experience," "readability," and "maintainability," backlang stands as a monument to the opposite philosophy:

- **Why be readable when you can be cryptic?**
- **Why flow naturally when you can flow backwards?**  
- **Why make sense when you can make developers suffer?**

backlang is a critique of our endless quest to make programming "easier." Sometimes, things should be hard. Sometimes, we should have to think backwards. Sometimes, we should question why we're doing what we're doing.

Also, it's kind of fun in a weird way.

---

## 🔧 Technical Details

- **Algorithm:** Simple line reversal (first line becomes last, last becomes first)
- **File format:** `.bck` files are plain text, editable in any editor
- **Special marker `##BCKL.NNL##`:** Because your text files are special snowflakes that don't need trailing newlines. This marker appears at the top of `.bck` files when the original was too cool for standard line endings. Don't worry, we'll preserve your artisanal, hand-crafted lack of newlines.
- **Language detection:** Automatically detects Python, JavaScript, and Bash via shebangs (`#!/usr/bin/env python3`) or file extensions (`.py`, `.js`, `.sh`)
- **Auto-execution:** Decodes `.bck` files to temporary numbered files, then routes to the appropriate interpreter
- **Encoding preservation:** Maintains original file encoding (UTF-8, ASCII, etc.)
- **Error handling:** Graceful failures with helpful error messages
- **Cross-platform:** Works on Linux, macOS, Windows

---

## ⚠️ Warnings

- **Do not use in production** (unless you hate your future self)
- **May cause confusion, frustration, and existential crisis**  
- **Your IDE's syntax highlighting will have no idea what's happening**
- **Code reviews will be... interesting**
- **Debugging backwards code is exactly as fun as it sounds**
- **The `run` command will actually execute your backwards code** - because apparently we've made this thing too convenient

---

## 🤝 Contributing

Think backlang isn't painful enough? Have ideas for making programming even more unnecessarily difficult? PRs welcome!

Areas for contribution:
- More creative ways to make code unreadable
- Documentation for writing specific languages backwards
- Support for right-to-left languages (because why not?)
- Integration with popular masochistic development workflows

---

## 🎭 Real Talk

Use it for:
- **Don't**

---

## 📄 License

under ☕️, check out [the-coffee-license](https://github.com/codinganovel/The-Coffee-License)

I've included both licenses with the repo, do what you know is right. The licensing works by assuming you're operating under good faith.

---

## 🎯 Final Thoughts

After years of languages trying to be more intuitive, more readable, and more developer-friendly, we finally have a tool that goes in the complete opposite direction.

backlang: Because programming wasn't hard enough already.

---

*"I can finally write Rust backwards. My suffering is complete." - A developer, probably*

*"This is either genius or the dumbest thing I've ever seen. I'm not sure which." - Another developer, definitely*