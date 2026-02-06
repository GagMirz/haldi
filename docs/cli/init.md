## Init

### Command
```bash
anvil init [MANIFEST_NAME]
```
Initialize anvil manifest in project (Creates anvil.json file in project). If [MANIFEST_NAME] was not specified anvil will try to use parent directory name.

#### Arguments
1. Path: (-p, --path) path to the project.  
  Default: ```./```.  
  Example:
```bash
anvil init -n hld -p ~/git/hld
```