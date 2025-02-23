# Downloader-go

Downloader-go is a simple command-line tool written in Go for downloading files with specific extensions from a provided URL.

## Features

- Scrape a given URL to find all links with specified extensions (e.g., `.zip`, `.mp3`, `.mp4`).
- Download the identified files to the default Downloads folder for both Windows and Linux systems.
- Shows a progress bar while downloading each file.

## Requirements

- Go 1.24+ (for building the project)
- Access to a terminal/command prompt

## Installation

To build and install the application, follow these steps:

1. Clone the repository:

    ```bash
    git clone https://github.com/your-username/downloader-go.git
    cd downloader-go
    ```

2. Build the project:

    ```bash
    go build
    ```

This will generate an executable named `downloader-go` (or `downloader-go.exe` on Windows).

## Usage

### Command-line flags

- `-u` or `--url`: URL of the page to scrape for links.
- `-e` or `--extension`: Extensions to filter links (comma-separated, e.g., `.zip, .mp3, .mp4`).

### Example

To download `.zip` files from a specific website, run the following command:

```bash
./downloader-go -u https://example.com -e .zip,.mp3
