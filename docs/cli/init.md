## Init

### Command
```bash
haldi init [MANIFEST_NAME]
```
Initialize haldi manifest in project (Creates haldi.json file in project). If [MANIFEST_NAME] was not specified haldi will try to use parent directory name.

#### Arguments
1. Path: (-p, --path) path to the project.  
  Default: ```./```.  
  Example:
```bash
haldi init -n hld -p ~/git/hld
```