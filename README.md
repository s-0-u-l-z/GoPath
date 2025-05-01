# 🚀 GoPath - Fast Web Directory Scanner in Go

![GoPath](https://img.shields.io/badge/GoPath-Scanner-blue?style=for-the-badge)  
A fast and lightweight **web directory scanner** built in **Go**, inspired by `dirsearch` but optimized for speed.  

## 🔥 Features  
✅ Multi-threaded scanning for high performance  
✅ Supports **custom wordlists** and **file extensions**  
✅ Supports scanning **multiple URLs from a file (`-l`)**  
✅ HTTP status code filtering (`200, 403, 500, etc.`)  
✅ Works with **proxy support** (`http://127.0.0.1:8080`)  
✅ **Recursive scanning** (`-r` option)  
✅ Supports **User-Agent customization**  
✅ **Saves results** to an output file  

---

## ⚡ **Installation**  
### ✅ Install via `go install`
Run this command to install **GoPath** directly from GitHub:  
```sh
go install github.com/s-0-u-l-z/GoPath@latest
```
Then, you can run it anywhere using:  
```sh
GoPath -u https://example.com
```

### ✅ Manual Build (Linux/macOS/Windows)
If you want to **build from source**, clone the repo and compile it:  
```sh
git clone https://github.com/s-0-u-l-z/GoPath.git
cd GoPath
go build -o GoPath main.go
sudo mv GoPath /usr/local/bin/  # Move to global path
```

Now, run:  
```sh
GoPath -u https://example.com
```

---

## 🎯 **Usage**
```sh
GoPath -u <target-url> -w <wordlist> -e <extensions> -t <threads> [options]
```
or  
```sh
GoPath -l <urls-file> -w <wordlist> -t <threads> [options]
```

### **Example Scans**
#### 🔹 **Basic Scan (Single URL)**  
```sh
GoPath -u https://example.com
```
#### 🔹 **Scan Multiple URLs from a File**  
```sh
GoPath -l urls.txt -w wordlist.txt
```
(📌 `urls.txt` should contain multiple URLs, one per line.)  

#### 🔹 **Using a Custom Wordlist**  
```sh
GoPath -u https://example.com -w custom_wordlist.txt
```
#### 🔹 **Scanning with File Extensions**  
```sh
GoPath -u https://example.com -e php,html,js
```
#### 🔹 **Multi-threaded Fast Scan (25 threads)**  
```sh
GoPath -u https://example.com -t 25
```
#### 🔹 **Recursive Scanning**  
```sh
GoPath -u https://example.com -r
```
#### 🔹 **Save Results to a File**  
```sh
GoPath -u https://example.com -o results.txt
```
#### 🔹 **Using a Proxy**  
```sh
GoPath -u https://example.com -p http://127.0.0.1:8080
```

---

## ⚙️ **Options**
| Option | Description |
|--------|-------------|
| `-u, --url` | Target URL (required for single scan) |
| `-l, --urls-file` | File containing multiple URLs to scan |
| `-w, --wordlist` | Wordlist file (default: `wordlist.txt`) |
| `-e, --extensions` | Extensions to scan (e.g. `php,html,js`) |
| `-t, --threads` | Number of threads (default: `10`) |
| `-r, --recursive` | Enable recursive scanning |
| `-i, --include-status` | Include specific HTTP status codes |
| `-x, --exclude-status` | Exclude specific HTTP status codes |
| `-o, --output` | Save results to a file |
| `-p, --proxy` | Use a proxy (`http://127.0.0.1:8080`) |
| `--random-agent` | Use a random User-Agent |
| `--help` | Show help menu |
| `--version` | Show GoPath version |

---

## ⚠️ **Performance Warning**
GoPath is optimized for **maximum speed**, but it can be **very resource-intensive**, especially in `--mode fast`.

> If you're seeing crashes, timeouts, or hangs:
> - Lower your `-t` value (e.g. `-t 25`)
> - Avoid recursive mode or massive wordlists
> - Monitor system usage while running

---

### ⚙️ **Coming Soon**
| Flag | Description |
|------|-------------|
| `--max-ram <MB>` | (Coming Soon) Automatically limits number of threads if RAM is below the threshold |
| `--performance` | (Coming Soon) Shows live scan speed, memory usage, and estimated system load |

> 🚧 These features are in development and will help avoid system overload, especially on lower-end machines.

---

## 💡 **Why Use GoPath?**
🚀 **Optimized for Speed** → Multi-threaded & fast  
🌐 **Flexible** → Custom wordlists, extensions, proxies  
🔍 **Accurate** → Status code filtering for real results  
🛠 **Simple to Use** → No unnecessary features, just works  

---

## ❤️ **Contributing**
Want to improve GoPath? Feel free to fork this repo, submit PRs, or report issues!  
```sh
git clone https://github.com/s-0-u-l-z/GoPath.git
```
Let's make **GoPath** even better together! 🎯🔥  

---

## 📜 **License**
GoPath is open-source under the **GNU General Public License v2.0 (GPL-2.0)**.  
You may modify and distribute it, but **any changes must also be open-source under GPL-2.0**.  

📖 **Full License Text:** [GNU GPL v2.0](https://www.gnu.org/licenses/old-licenses/gpl-2.0.html)  

---

## 📞 **Contact**
📧 **GitHub:** [s-0-u-l-z](https://github.com/s-0-u-l-z)  
📌 **Issues & Bugs?** [Report here](https://github.com/s-0-u-l-z/GoPath/issues)  

---

## 📊 **Performance Benchmarks**

We ran a head-to-head performance test between **GoPath** and `dirsearch` to see how they compare:

### 🚀 GoPath vs dirsearch

**Test Setup**: 10 runs of each tool on the same target system.

```
🚀 Running GoPath 10 times from source...
⚙️  GoPath run 1
⏱️  took .076s
⚙️  GoPath run 2
⏱️  took .012s
⚙️  GoPath run 3
⏱️  took .012s
⚙️  GoPath run 4
⏱️  took .012s
⚙️  GoPath run 5
⏱️  took .012s
⚙️  GoPath run 6
⏱️  took .011s
⚙️  GoPath run 7
⏱️  took .011s
⚙️  GoPath run 8
⏱️  took .011s
⚙️  GoPath run 9
⏱️  took .011s
⚙️  GoPath run 10
⏱️  took .012s

🔍 Running dirsearch 10 times...
🧭 dirsearch run 1
⏱️  took 8.479s
🧭 dirsearch run 2
⏱️  took 9.337s
🧭 dirsearch run 3
⏱️  took 9.995s
🧭 dirsearch run 4
⏱️  took 8.049s
🧭 dirsearch run 5
⏱️  took 7.738s
🧭 dirsearch run 6
⏱️  took 7.893s
🧭 dirsearch run 7
⏱️  took 6.271s
🧭 dirsearch run 8
⏱️  took 6.625s
🧭 dirsearch run 9
⏱️  took 6.890s
🧭 dirsearch run 10
⏱️  took 9.448s
```

### ✨ ============================= ✨  
🎯 **Average GoPath time**: **0.018s**  
🎯 **Average dirsearch time**: **8.072s**  
⚡️ **GoPath is approximately _448x faster_ than dirsearch**  
✨ ============================= ✨
