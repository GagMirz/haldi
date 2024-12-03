# Apply

## Apply

### Command
```bash
haldi apply
```
Applies haldi.json manifest, adding it to user manifests.

#### Arguments
1. Path: (-p, --path) path to the project.  
  Default: ```./```.  
  __Note__: manifest should be named ```haldi.json```, if file name was not provided it will be automatically extended.  
  Example:
```bash
haldi apply -p ~/projects/my-awesome-project/haldi.json
```