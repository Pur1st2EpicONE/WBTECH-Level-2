## L2.13

![cut banner](assets/banner.png)

<h3 align="center">A simplified implementation of the UNIX cut utility in Go, supporting field selection, custom delimiters, and input from multiple files or STDIN.</h3>

<br>

## Supported Flags

-f N — Specify which fields (columns) to print. Field numbers can be listed individually or as ranges.

-d N — Use a custom field delimiter. Defaults to tab (\t) if not specified. 

-s N — Only print lines that contain the delimiter. Lines without the delimiter are ignored.

Multiple flags can be combined (e.g., -f 2-4 -d ',' -s for selecting columns 2 to 4 from comma-separated lines, ignoring lines without commas).

<br>

## Installation and usage

1) Build the project:

```bash
make
```

2) Run the utility:
```bash
./cut OPTION... [FILE]...
```

<br>

## Cool features

* Field Selection & Error Handling: Supports individual columns and ranges. Invalid fields or ranges (e.g., 3-1, 0-3, 1-a) are safely handled, mimicking GNU cut behavior.

* Written in Go. That's cool, right?

<br>

## Testing & Linting

Run tests and ensure code quality:

```bash
make test        # Unit tests
make diff_test   # Compare output with GNU cut (tested on Linux; results may differ on macOS)
make lint        # Linting checks
```