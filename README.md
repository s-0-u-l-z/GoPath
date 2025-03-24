# 🚀 GoPath - Fast Web Directory Scanner in Go

![GoPath](https://img.shields.io/badge/GoPath-Scanner-blue?style=for-the-badge)  
A fast and lightweight **web directory scanner** built in **Go**, inspired by `dirsearch` but optimized for speed.  

## 🔥 Features  
✅ Multi-threaded scanning for high performance  
✅ Supports **custom wordlists** and **file extensions**  
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

### **Example Scans**
#### 🔹 **Basic Scan**  
```sh
GoPath -u https://example.com
```
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
| `-u, --url` | Target URL (required) |
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
GoPath is open-source under the **MIT License**.  
Feel free to use and modify it!  

---

## 📞 **Contact**
📧 **GitHub:** [s-0-u-l-z](https://github.com/s-0-u-l-z)  
📌 **Issues & Bugs?** [Report here](https://github.com/s-0-u-l-z/GoPath/issues)  

---
