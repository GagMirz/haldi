## Init

### Command
```bash
haldi init
```
Initialize haldi manifest in project (Creates haldi.json manifest in project).

#### Arguments
1. Name: (-n, --name) add manifest name, this will be used for later manifest usage.  
  Default: project directory name.  
  Example:
```bash
haldi init -n hld
```

2. Path: (-p, --path) path to the project.  
  Default: ```./```.  
  Example:
```bash
haldi init -n hld -p ~/git/hld
```