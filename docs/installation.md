# Installation

<!-- TODO: Add simpler way for installation -->

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
- Create haldi configuration directory in ```~/```

### Dev requirements
- docsify (npm)

### Command
```bash
make docs-install
```
To host docsify run ```make docs```
