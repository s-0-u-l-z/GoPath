# GoPath - Fast Web Directory Scanner in Go

A fast and lightweight **web directory scanner** built in **Go**, inspired by `dirsearch` but optimized for speed.

## Features

* Multi-threaded scanning for high performance
* Supports **custom wordlists** and **file extensions**
* Supports scanning **multiple URLs from a file (-l)**
* HTTP status code filtering (200, 403, 500, etc.)
* Works with **proxy support** ([http://127.0.0.1:8080](http://127.0.0.1:8080))
* **Recursive scanning** (-r option)
* Supports **User-Agent customization**
* **Saves results** to an output file

---

## Installation

### Install via `go install`

Run this command to install **GoPath** directly from GitHub:

```sh
go install github.com/s-0-u-l-z/GoPath@latest
```

Then run it anywhere:

```sh
GoPath -u https://example.com
```

### Manual Build (Linux/macOS/Windows)

```sh
git clone https://github.com/s-0-u-l-z/GoPath.git
cd GoPath
go build -o GoPath main.go
sudo mv GoPath /usr/local/bin/
```

Now run:

```sh
GoPath -u https://example.com
```

---

## Usage

```sh
GoPath -u <target-url> -w <wordlist> -e <extensions> -t <threads> [options]
```

or

```sh
GoPath -l <urls-file> -w <wordlist> -t <threads> [options]
```

### Example Scans

#### Basic Scan (Single URL)

```sh
GoPath -u https://example.com
```

#### Scan Multiple URLs from a File

```sh
GoPath -l urls.txt -w wordlist.txt
```

#### Using a Custom Wordlist

```sh
GoPath -u https://example.com -w custom_wordlist.txt
```

#### Scanning with File Extensions

```sh
GoPath -u https://example.com -e php,html,js
```

#### Multi-threaded Scan (25 threads)

```sh
GoPath -u https://example.com -t 25
```

#### Recursive Scanning

```sh
GoPath -u https://example.com -r
```

#### Save Results to a File

```sh
GoPath -u https://example.com -o results.txt
```

#### Using a Proxy

```sh
GoPath -u https://example.com -p http://127.0.0.1:8080
```

---

## Options

| Option                 | Description                        |
| ---------------------- | ---------------------------------- |
| `-u, --url`            | Target URL                         |
| `-l, --urls-file`      | File containing multiple URLs      |
| `-w, --wordlist`       | Wordlist file                      |
| `-e, --extensions`     | Extensions to scan                 |
| `-t, --threads`        | Number of threads                  |
| `-r, --recursive`      | Enable recursive scanning          |
| `-i, --include-status` | Include specific HTTP status codes |
| `-x, --exclude-status` | Exclude specific HTTP status codes |
| `-o, --output`         | Save results to file               |
| `-p, --proxy`          | Use a proxy                        |
| `--random-agent`       | Random User-Agent                  |
| `--help`               | Show help                          |
| `--version`            | Show version                       |

---

## Performance Warning

GoPath is optimized for speed but can be resource-intensive.

If you see crashes or hangs:

* Lower your `-t` value
* Avoid recursive mode or huge wordlists
* Monitor system usage

---

### Coming Soon

| Flag             | Description                               |
| ---------------- | ----------------------------------------- |
| `--max-ram <MB>` | Automatically limit threads if RAM is low |
| `--performance`  | Live scan speed and memory usage          |

---

## Why Use GoPath?

* Optimized for speed
* Flexible with custom options
* Accurate status filtering
* Simple to use

---

## Contributing

Feel free to fork the repo or submit PRs:

```sh
git clone https://github.com/s-0-u-l-z/GoPath.git
```

---

## License

GoPath is open-source under **GPL-2.0**.
Any modifications must also be GPL-2.0.

---

## Contact

GitHub: `s-0-u-l-z`
Issues: GitHub Issues

---

## Performance Benchmarks

Comparison between **GoPath** and `dirsearch`.

### GoPath (10 runs)

```
GoPath run 1
took .076s
GoPath run 2
took .012s
GoPath run 3
took .012s
GoPath run 4
took .012s
GoPath run 5
took .012s
GoPath run 6
took .011s
GoPath run 7
took .011s
GoPath run 8
took .011s
GoPath run 9
took .011s
GoPath run 10
took .012s
```

### dirsearch (10 runs)

```
dirsearch run 1
took 8.479s
dirsearch run 2
took 9.337s
dirsearch run 3
took 9.995s
dirsearch run 4
took 8.049s
dirsearch run 5
took 7.738s
dirsearch run 6
took 7.893s
dirsearch run 7
took 6.271s
dirsearch run 8
took 6.625s
dirsearch run 9
took 6.890s
dirsearch run 10
took 9.448s
```

### Results

Average GoPath time: **0.018s**
Average dirsearch time: **8.072s**
GoPath is approximately **448x faster**.
