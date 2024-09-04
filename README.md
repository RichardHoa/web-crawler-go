# Go Web Crawler

This project is a simple web crawler written in Go. It allows you to crawl web pages starting from a specified URL, with options to set concurrency and limit the maximum number of pages to crawl.

## Features

- Crawl web pages starting from a given URL.
- Set the level of concurrency for crawling.
- Limit the maximum number of pages to be crawled.

## Installation

To get started with the Go web crawler, you need to have Go installed on your machine. Follow these steps to set up and run the project:

1. **Clone the repository**:
   ```bash
   git clone <repository-url>
   cd go-web-crawler
   ```

2. **Install dependencies** (if any):
   ```bash
   go mod tidy
   ```

## Running the Crawler

You can run the crawler in two ways: using `go run` or by building the application with `go build`.

### Option 1: Run Directly with `go run`

To run the crawler directly without building an executable, use:

```bash
go run . --url <URL> --concurrency <CONCURRENCY_LEVEL> --max-pages <MAX_PAGES>
```

- `--url`: The starting URL for the crawler.
- `--concurrency`: The number of concurrent workers to use (default is 1).
- `--max-pages`: The maximum number of pages to crawl.

### Option 2: Build and Run the Executable

1. **Build the application**:
   ```bash
   go build -o crawler
   ```

2. **Run the built executable**:
   ```bash
   ./crawler --url <URL> --concurrency <CONCURRENCY_LEVEL> --max-pages <MAX_PAGES>
   ```

## Example Usage

```bash
# Running directly
go run . --url "https://example.com" --concurrency 5 --max-pages 100

# Running the built executable
./crawler --url "https://example.com" --concurrency 5 --max-pages 100
```

In both cases, you must specify the `--url` to start crawling from, the `--concurrency` level to determine how many requests can run in parallel, and the `--max-pages` to limit the number of pages the crawler will visit.

