## Installation

<!-- TODO: Add simpler way for installation -->

## Build files
You can see latest releases and build files available in [GitHub Releases](https://github.com/GagMirz/anvil/releases)

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
- Install anvil terminal tool in ```/usr/local/bin```
- On initial run, anvil will
    - Create anvil configuration directory in user home as ```~/.anvil```

### Dev requirements
- docsify (npm)

### Command
```bash
make docs-install
```
To host docsify run ```make docs```
