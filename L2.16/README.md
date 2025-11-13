## L2.16

![wget banner](assets/banner.png)

<h3 align="center">A simplified implementation of the UNIX wget utility in Go, supporting recursive downloads, multiple resource types, and offline browsing.</h3>

<br>

## Supported Flags

-r, --recursive — enable recursive download.

-l N, --level N — recursion depth.

If no flags are provided, only the initial HTML page (and robots.txt if present) will be downloaded.

<br>

## Installation and usage

1) Optionally, edit [the config file](config.yaml) to customize your preferences, then build the project using the Makefile command:

```bash
make
```

2) Run the utility:
```bash
./wget [OPTION]... [URL]...
```

<br>

## Cool features

* Recursive page traversal within the specified domain.

* Asynchronous downloads with parallelism limits to prevent excessive requests.

* Automatic path localization.

* Support for multiple resource types: HTML, CSS, JS, images, videos, audio, fonts, favicons, and manifests.

* Robots.txt support (enabled by default; can be disabled in the configuration).

* Detailed statistics: total files, total size, elapsed time, and average speed.

<br>

## Testing & Linting

Run tests and ensure code quality:

```bash
make test        # Unit tests
make lint        # Linting checks
```

For live scraping tests, a good website to use is [Books to Scrape](https://books.toscrape.com/), but it’s better not to start scraping from the homepage, as it contains the list of categories and therefore the most links.