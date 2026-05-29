# webpcompress

A fast, simple CLI tool written in Go for converting and compressing images to WebP format using Google's `cwebp` encoder.

## Prerequisites

`webpcompress` requires `cwebp` to be installed on your system.

**macOS**
```bash
brew install webp
```

**Ubuntu/Debian**
```bash
sudo apt install webp
```

**Windows**
```bash
choco install webp
```

Verify the installation with:
```bash
cwebp -version
```

## Installation

### Option 1 — Build from source

Requires [Go](https://go.dev/dl) 1.21 or later.

```bash
git clone https://github.com/khunnaball/webpcompress.git
cd webpcompress
go install
```

This builds the binary and adds it to your `$GOPATH/bin`, making `webpcompress` available anywhere on your system.

### Option 2 — Download a prebuilt binary

Download the latest binary for your platform from the [releases page](https://github.com/khunnaball/webpcompress/releases).

**macOS (Apple Silicon)**
```bash
chmod +x webpcompress-darwin-arm64
mv webpcompress-darwin-arm64 /usr/local/bin/webpcompress
```

**macOS (Intel)**
```bash
chmod +x webpcompress-darwin-amd64
mv webpcompress-darwin-amd64 /usr/local/bin/webpcompress
```

**Linux**
```bash
chmod +x webpcompress-linux-amd64
mv webpcompress-linux-amd64 /usr/local/bin/webpcompress
```

## Usage

### Convert a single file

```bash
webpcompress image.jpg
```

### Convert multiple files

```bash
webpcompress image.jpg photo.png banner.gif
```

### Convert all images in a directory

```bash
webpcompress --dir ./images
```

### Specify output directory and quality

```bash
webpcompress image.jpg --quality 90 --output ./converted
```

### Overwrite existing files

By default, `webpcompress` skips files that already exist in the output directory. Use `--force` to overwrite them:

```bash
webpcompress image.jpg --force
```

## Flags

| Flag | Shorthand | Default | Description |
|------|-----------|---------|-------------|
| `--quality` | `-q` | `80` | WebP quality, 1–100 |
| `--output` | `-o` | `./output` | Output directory |
| `--dir` | `-d` | | Convert all images in a directory |
| `--force` | `-f` | `false` | Overwrite existing output files |
| `--help` | `-h` | | Show help |

## Supported formats

`webpcompress` will attempt to convert the following file types:

- `.jpg` / `.jpeg`
- `.png`
- `.gif`
- `.tiff`
- `.bmp`
- `.webp` — existing WebP files can be recompressed at a lower quality. Note that each recompression pass introduces additional quality loss as WebP is a lossy format. Use with care.

All other file types are silently skipped.

## Output

After each run, `webpcompress` prints a summary of the results:

```
Done: 4 converted, 2 skipped, 0 error(s)
```

## Building for other platforms

```bash
# macOS Apple Silicon
GOOS=darwin GOARCH=arm64 go build -o webpcompress-darwin-arm64

# macOS Intel
GOOS=darwin GOARCH=amd64 go build -o webpcompress-darwin-amd64

# Linux
GOOS=linux GOARCH=amd64 go build -o webpcompress-linux-amd64

# Windows
GOOS=windows GOARCH=amd64 go build -o webpcompress.exe
```

## License

MIT — see [LICENSE](LICENSE) for details.
