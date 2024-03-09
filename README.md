# RepoSort

RepoSort is a command-line tool written in Go that organizes repositories into language-specific directories. It is designed to help developers maintain a clean workspace by automatically sorting repositories based on the programming language they are primarily written in.

## Installation

Before installing RepoSort, ensure you have Go installed on your system. You can then install RepoSort by cloning the repository and building the binary.

```bash
git clone https://github.com/marcusziade/reposort.git
cd reposort/cmd
go build
```

This will generate a `reposort` executable in your current directory.

## Usage

To use RepoSort, simply provide the directory path containing the repositories as an argument to the executable:

```bash
./reposort /path/to/repositories
```

RepoSort will then process each repository in the provided directory and sort them into language-specific subdirectories.

## Requirements

RepoSort requires the following:

-   Go version 1.x or higher
-   Access to the directory containing the repositories

## Contributing

Contributions to RepoSort are welcome. If you have suggestions for improvements or bug fixes, please open an issue or submit a pull request.

## License

RepoSort is released under the MIT License. See the LICENSE file for more details.

## Support

If you encounter any problems or have any questions about using RepoSort, please open an issue on the GitHub repository, and I will be happy to assist you.
