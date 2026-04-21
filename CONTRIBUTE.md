# 🤝 Contributing to Gunter

> "Pull requests are welcome, but please don't bring actual malware." - Probably not a maintainer

## 🎉 How to Contribute

We love contributions! Gunter is a learning project, so we're excited to see what you bring to the table.
Whether it's fixing a typo, adding a new feature, or just improving the documentation, your help is appreciated.

### 📝 What We're Looking For

- **Bug fixes**: Found something broken? Let's fix it!
- **Feature ideas**: Especially things that move us toward our AV engine goals
- **Documentation improvements**: Make it funnier, clearer, or more complete
- **Code quality**: Better patterns, more comments, etc.
- **Tests**: When we start writing them (contributions here will be especially valued)

### 🚫 What We're Not Looking For

- Actual malware samples (please don't send us viruses)
- Off-topic spam or advertising
- Changes that break the build without fixing something important

## 🔧 Getting Started

1. **Fork the repository** (click the Fork button at the top right)
2. **Clone your fork**:
   ```bash
   git clone https://github.com/your-username/gunter.git
   cd gunter
   ```
3. **Create a branch** for your changes:
   ```bash
   git checkout -b feature/your-feature-name
   # or
   git checkout -b fix/your-bug-fix
   ```
4. **Make your changes** (please follow the existing code style)
5. **Test your changes** (if applicable):
   ```bash
   go run cmd/main.go --help   # Should show help without panicking
   ```
6. **Commit your changes**:
   ```bash
   git commit -m "feat: add something awesome"
   # or
   git commit -m "fix: resolve nil pointer in scanner"
   ```
7. **Push to your fork**:
   ```bash
   git push origin feature/your-feature-name
   ```
8. **Open a Pull Request** against the `main` branch of this repository

## 📏 Coding Style

- We follow [Go's formatting standards](https://golang.org/doc/effective_go.html#formatting) (run `go fmt ./...`)
- Keep lines to a reasonable length (80-100 characters)
- Write clear, concise commit messages
- Comment complex logic, but don't over-comment obvious things
- Error messages should be informative and user-friendly

## 🐛 Reporting Issues

Found a bug or have a feature request? Please open an issue:
1. Check if it's already been reported
2. Use a clear, descriptive title
3. Include steps to reproduce (for bugs)
4. Specify your Go version and OS
5. Add screenshots if relevant

## 📜 Licensing

By contributing, you agree that your contributions will be licensed under the MIT License (same as the project).

## 🙌 Recognition

Contributors will be acknowledged in the README (if they want to be) and in the git history.
We believe in giving credit where it's due!

## ❓ Questions?

Feel free to open an issue or reach out in discussions if you're unsure about anything.

Remember: The goal is to learn and have fun while building something useful (or at least entertaining).
Let's make Gunter the best darn hash-scanner-turned-AV-engine we can!

Happy hacking! 🐱‍💻