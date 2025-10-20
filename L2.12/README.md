## L2.12

![grep banner](assets/banner.png)

<h3 align="center">A simplified implementation of the UNIX grep utility in Go, supporting various flags, syntax highlighting, and input from multiple files or STDIN.</h3>

<br>

## Supported Flags

-A N — Print N lines after each matching line (trailing context).  

-B N — Print N lines before each matching line (leading context).  

-C N — Print N lines around each matching line (both before and after; equivalent to -A N -B N).  

-c — Print only the count of matching lines instead of the lines themselves.  

-i — Ignore case when matching lines.  

-v — Invert match: print lines that do not match the pattern.  

-F — Treat the pattern as a fixed string, not a regular expression.  

-n — Print line numbers before each matching line.  

Multiple flags can be combined (e.g., -C 2 -n -i for context, line numbers, and case-insensitive search).  

<br>

## Installation and usage

1) Build the project:

```bash
make
```

2) Run the utility:
```bash
./grep [OPTION]... PATTERN [FILE]...
```

<br>

## Cool features

* Colored Output: Highlights text in different colors depending on the match type and flags, similar to GNU grep.  

* Flexible Flags: Supports combinations of context lines, counts, line numbers, inverted match, and fixed-string mode.  

* GNU-like Error Handling: Mimics GNU grep exit codes and error messages for invalid input or missing files.  

* Multiple Files: Handles multiple input files, showing filenames when needed, or reads from STDIN if none are provided.

* Circular Buffer is used to store lines when flag -B is enabled.  

* Circular Affer is not used to store lines when flag -A is enabled.  

<br>

## Testing & Linting

Run tests and ensure code quality:

```bash
make test        # Unit tests
make diff_test   # Differential tests comparing output with GNU grep (tested on Linux; results may differ on macOS)
make lint        # Linting checks
```