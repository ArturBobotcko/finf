# finf

`finf` is a simple command-line tool that provides detailed information about files, including type detection, permissions, and metadata.

## Features
- **File type detection**: Given a file, it detects its extension and outputs a description of the file type (e.g., `.go` as a "Go source file").
- **File Permissions**: It describes the file's permissions (read, write, execute) using symbolic notation (e.g., `rw-r--r--`).
- **File Metadata**: Displays the file's size, last modified time, and access rights.

## Installation

### Pre-built Binaries

You can download the latest pre-built release of `finf` for your platform directly from the GitHub [Releases page](https://github.com/ArturBobotcko/finf/releases).

- For **Linux**: `finf_Linux_x86_64.tar.gz`
- For **macOS**: `finf_Darwin_x86_64.tar.gz`
- For **Windows**: `finf_Windows_x86_64.zip`

### Manual Installation

If you prefer building `finf` from source, you can follow these steps:

1. Clone the repository:
   
   ```bash
   git clone https://github.com/ArturBobotcko/finf.git
   cd finf
   ```
   
2. Install Go (if not already installed): [Go installation guide.](https://go.dev/doc/install)
   
3. Build the application:
   
   ```
   go build -o finf
   ```
   
4. Make the binary executable and move it to a directory in your `PATH`, for example:
 
   ```
   chmod +x finf
   mv finf /usr/local/bin/
   ```

### Usage

Once installed, you can use finf from the command line:

```
finf <file-path>
```

### Example:

```
finf main.go
```

This will output the file's type, permissions, size, and metadata, e.g.:

```
Name: main.go
File type: "Go source code file"
Size: 228 Bytes
Access: Regular file; User can read; User can write; Group can read; Others can read; 
Modified: 2024-11-06 01:23:22.68627158 +0300 MSK
```

## Contributing

We welcome contributions to `finf`. To contribute, follow these steps:

1. Fork the repository on GitHub.
   
2. Create a new branch (`git checkout -b feature-name`).
   
3. Make your changes and commit them (`git commit -am 'Add new feature'`).
   
4. Push the branch to your fork (`git push origin feature-name`).
   
5. Submit a pull request to the main repository.
