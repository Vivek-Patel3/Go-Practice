## Installation

From the project root (`command-line/`):

```bash
go install ./cmd/motadata
```
Make sure Go’s binary directory is in your PATH:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```
    Tip: Add the line above to your ~/.bashrc or ~/.zshrc to make the change permanent.

## Usage

Run the CLI from anywhere:

```bash
motadata --name Vivek
```

Example Output

    Hello, Vivek

    Environment Info:
    OS: linux
    Architecture: amd64
    Go version: go1.22.0
    User: vivek

---