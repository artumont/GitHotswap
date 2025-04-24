# 🔄 GitHotswap

A quick tool to switch between Git user profiles. Perfect for managing multiple identities (like work and personal) without editing .gitconfig manually.

## ✨ Features

- Easily switch between multiple Git user profiles
- Create, edit, and manage Git user profiles
- List all available profiles
- Quick swap between profiles using menu or hotswap
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

### Swap
Every operation that is related to the swap of git profiles.

```bash
git-hotswap swap              # Swap to a profile using the active mode (menu or hotswap)
git-hotswap swap to <profile> # Swap to a specific profile
git-hotswap swap mode <menu|hotswap> # Change the swap mode
```

### Help
Shows help for all commands.

```bash
git-hotswap help           # Show minimized help information for all commands
git-hotswap help <command> # Show detailed help information for a specific command
```

### Profile
Every operation that is related to the user profile.

```bash
git-hotswap profile create <profile> # Creates a new profile
git-hotswap profile edit <profile>   # Edits a profile
git-hotswap profile delete <profile> # Deletes a profile
git-hotswap profile list            # Lists all profiles
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/amazing-feature`
3. Commit your changes: `git commit -m 'Add amazing feature'`
4. Push to the branch: `git push origin feature/amazing-feature`
5. Open a pull request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
