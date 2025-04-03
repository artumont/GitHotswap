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

### Help Command
Get help information for all commands or a specific command:
```bash
git-hotswap help           # Show minimized help information for all commands
git-hotswap help <command> # Show detailed help information for a specific command
```

### Managing Profiles
```bash
git-hotswap profile create <profile>  # Create a new profile
git-hotswap profile delete <profile>  # Delete a profile
git-hotswap profile edit <profile>    # Edit a profile
git-hotswap profile current          # Get the current profile
git-hotswap profile list            # List all profiles
```

### Switching Profiles
```bash
git-hotswap swap          # Swap to a profile depending on the active one
git-hotswap swap menu     # Swap to a profile using the menu
git-hotswap swap hotswap  # Swap to a profile using hotswap
git-hotswap swap to <profile>  # Swap to a specific profile
```

### Configuration Management
```bash
git-hotswap config show           # Show the current configuration
git-hotswap config reset          # Reset the configuration file to default
git-hotswap config open           # Open the configuration file in the default editor
git-hotswap config backup <path>  # Backup the configuration file
git-hotswap config restore <path> # Restore the configuration file from backup
git-hotswap config swap_method <method> # Set the swap method (menu or hotswap)
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/amazing-feature`
3. Commit your changes: `git commit -m 'Add amazing feature'`
4. Push to the branch: `git push origin feature/amazing-feature`
5. Open a pull request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
