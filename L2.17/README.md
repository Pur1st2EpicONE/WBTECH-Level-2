## L2.17

![telnet banner](assets/banner.png)

<h3 align="center">A simplified implementation of the UNIX telnet utility in Go, supporting basic TCP connections, interactive input/output, and graceful shutdown handling.</h3>

<br>

## Supported Flags

--timeout DURATION — set the connection timeout (default: 10s).

Positional arguments: [host-name [port]] — the host to connect to and the optional TCP port (default: 23).

<br>

## Installation and usage

1) Build the project using the Makefile:

```bash
make
```

2) Run the telnet client:
```bash
./telnet [--timeout DURATION] [host-name [port]]
```
<br>

Example:

```bash
./telnet --timeout 15s mtrek.com 1701 ## (Star Trek–themed multiplayer RTS) 
```

<br>

## Features

* Interactive stdin/stdout I/O with real-time server output.  

* Graceful handling of EOF (Ctrl+D) and SIGINT/SIGTERM.

<br>

## Cool feature

* Mimics the standard telnet behavior, including error messages and exit codes.

<br>

## Testing & Linting

Run tests and ensure code quality:

```bash
make test        # Unit tests
make lint        # Linting checks
```

<br>

For live testing, start a local echo server:

```bash
make server
```

Then connect using the telnet client:

```bash
./telnet localhost 8080
```

<br>

Stop the server when done:

```bash
make stop
```