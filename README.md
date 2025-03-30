# 🔄 GitHotswap

A quick tool to switch between Git user profiles. Perfect for managing multiple identities (like work and personal) without editing .gitconfig manually.

## ✨ Features

- Easily switch between multiple Git user profiles
- Create, edit, and manage Git user profiles
- List all available profiles
- Quick swap between two commonly used profiles
- Works with local Git repositories

## 🛠️ Technologies

- Go 1.22+
- Git integration

## 📋 Prerequisites

- Git installed and configured
- Go 1.22 or later
- Windows environment (currently Windows-only)

## 📦 Installation

### Option 1: Download binary

Download the latest release from the [Releases page](https://github.com/artumont/GitHotswap/releases).

### Option 2: Build from source

1. Clone the repository:
```bash
git clone https://github.com/artumont/GitHotswap
cd GitHotswap
```

2. Build the application:
```bash
go build -o ./bin/git-hotswap.exe ./src/
```

Or use the provided build script:
```bash
build.cmd
```

## 🚀 Usage

### Managing Profiles

Add a new profile:
```bash
git-hotswap profile add --key work --name "John Doe" --email "john.doe@company.com"
```

List all profiles:
```bash
git-hotswap profile list
```

Remove a profile:
```bash
git-hotswap profile remove --key work
```

Edit a profile:
```bash
git-hotswap profile edit --key work --name "John Smith" --email "john.smith@company.com"
```

Rename a profile:
```bash
git-hotswap profile rename --key work --new company
```

### Switching Profiles

Switch to a specific profile:
```bash
git-hotswap swap work
```

Quick swap between two profiles (when only two profiles exist):
```bash
git-hotswap swap
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/amazing-feature`
3. Commit your changes: `git commit -m 'Add amazing feature'`
4. Push to the branch: `git push origin feature/amazing-feature`
5. Open a pull request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
