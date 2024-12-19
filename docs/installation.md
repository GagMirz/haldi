## Installation

<!-- TODO: Add simpler way for installation -->

## Build files
You can see latest releases and build files available in [GitHub Releases](https://github.com/GagMirz/haldi/releases)

## From source files

### Pre requirements
- Golang
- Makefile

### Command
```bash
make build && make install
```

These commands would
- Build binaries from source
- Install haldi terminal tool in ```/usr/local/bin```
- On initial run, haldi will
    - Create haldi configuration directory in user home as ```~/.haldi```

### Dev requirements
- docsify (npm)

### Command
```bash
make docs-install
```
To host docsify run ```make docs```
