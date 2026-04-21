# 🦠 Gunter: The Not-So-Scary Antivirus Engine

> "I come in peace... to scan your files for nasty stuff!" - Gunter, probably

## 🎯 Mission Statement

Gunter started as a simple hash-scanner with dreams of becoming a full-fledged antivirus engine. 
Our journey: **Signature-based detection → Heuristics → Behavior Analysis → Sandboxing → Machine Learning** 
*(because even antiviruses need to level up)*

## 🚀 Features (Current & Planned)

### 🔒 Current Implementation (v0.1 - The Hash Hunter)
- ✅ **Signature-based detection** using MD5, SHA1, and SHA256 hashes
- ✅ Recursive directory scanning for hash databases (`.txt` files)
- ✅ Single file and directory scanning modes
- ✅ Proper path sanitization (no more `../../../etc/passwd` surprises!)
- ✅ Hex-encoded hash comparisons (because `string([]byte{0xde,0xad,0xbe,0xef})` is sad)

### 🧪 Coming Soon (v0.2 - The Heuristic Hawk)
- 🔸 **Static heuristic analysis**: Suspicious section names, packer detection, entropy calculation
- 🔸 **Dynamic heuristic analysis**: API call monitoring, registry changes (Windows-focused initially)
- 🔸 YARA rule support (because regex is fun... said no one ever)

### 🛡️ Future Plans (v0.3+ - The ML Guardian)
- 🔸 **Feature extraction**: Opcodes, API sequences, byte n-grams, PE header features
- 🔸 **ML Models**: Random Forest for classification, Isolation Forest for anomaly detection
- 🔸 **Online learning**: Update models with new threat intelligence
- 🔸 **Ensemble methods**: Combine signature, heuristic, and ML results

## 📦 Installation

```bash
# Clone the repo (you're already here if you're reading this!)
git clone https://github.com/fr13nd230/gunter.git
cd gunter

# Build it (requires Go 1.16+)
make build

# Or run directly
go run cmd/main.go --help
```

## 🖥️ Usage

### Basic Scanning
```bash
# Scan a single file
gunter --file /path/to/suspicious.exe

# Scan a directory recursively
gunter --dir /path/to/sketchy/downloads
```

### Specify Custom Hash Database
```bash
# Default looks in ./database/
gunter --database /path/to/your/hashdb --file /tmp/malware.bin
```

### Example Output
```
========== GUNTER =========
-database 12453
-file /tmp/malware.bin
-dir 
Scanned files, 1 has been affected.
============================
```
*(In this case, "1 affected" means we found a hash match - time to quarantine!)*

## 🗂️ Project Structure

```
gunter/
├── cmd/                # Main application entry point
│   └── main.go         # Where the magic begins
├── database/           # Hash signature databases (one hash per line in .txt files)
│   ├── full-hash-md5-aa.txt
│   ├── full-hash-sha1-aa.txt
│   └── full-hash-sha256-aa.txt
├── build/              # Compiled binaries go here
├── Makefile            # Build automation (because typing go build is so 2010)
├── README.md           # You're reading it! (ﾉ◕ヮ◕)ﾉ*:･ﾟ✧
└── CONTRIBUTE.md       # How to join the fun (coming soon!)
```

## 🛠️ Development

### Prerequisites
- Go 1.16+ (we use `fmt.Sprintf("%x")` which is ancient by Go standards)
- A sense of humor (mandatory)

### Building
```bash
make build       # Creates binary in ./build/gunter
make clean       # Removes build artifacts
make test        # Runs tests (when we have them - contributions welcome!)
```

### Database Format
Hash databases are simple text files:
- One hash per line
- Supported formats: MD5 (32 hex chars), SHA1 (40 hex chars), SHA256 (64 hex chars)
- Mixed algorithms in same file? Sure, why not! We'll check all three.

Example (`database/md5-hashes.txt`):
```
d41d8cd98f00b204e9800998ecf8427e
098f6bcd4621d373cade4e832627b4f6
# Comments start with # (we skip them)
```

## 🧠 Why "Gunter"?

According to highly scientific research (aka Wikipedia), Gunter is:
1. A Chuukese word for "stonefish" (venomous and dangerous - like malware!)
2. A character from Arthurian legend (knightly duties = scanning files)
3. What you yell when you find a particularly nasty piece of malware

## 📚 Resources We Stole Ideas From (With Love)

- [ClamAV](https://www.clamav.net/) - The OG open-source AV
- [YARA](https://virustotal.github.io/yara/) - Pattern matching swiss army knife
- [PEfile](https://github.com/erocarrera/pefile) - PE format parser (for future heuristics)
- Various academic papers on ML for malware detection (we'll cite them when we get there)

## ⚠️ Disclaimer

> **Gunter is educational software.** 
> Don't rely on it for actual security protection until it grows up and gets a real job.
> We are not responsible for:
> - False sense of security
> - Actually finding malware and causing panic
> - Sudden urges to pursue a career in cybersecurity
> - Developing an unhealthy obsession with hash functions

## 🙌 Credits

Crafted with ☕️ and 💻 by **[fr13nd230](https://github.com/fr13nd230)**  
*(Human by day, malware hunter by night)*

## 📜 License

MIT License - see [LICENSE](LICENSE) file for details  
*(Basically: do what you want, but don't sue us if Gunter misses something)*

---

*"In a world full of trojans, be a Gunter."* - Probably not a real proverb, but it should be.