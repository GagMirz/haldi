## Config

### Command
```bash
haldi config [SUBCOMMAND]
```

List or Modify haldi cli user configurations

## Available configurations

- shell: sets which shell should be used for running haldi aliases
  - Available values are zsh, bash, dash, ksh, sh

## Subcommands
### show
```bash
haldi config show [ATTRIBUTE]
```
If [ATTRIBUTE] is not specified lists all haldi cli configurations attributes with their values  
Result:
```bash
Shell:  zsh
Config2:  value
Config3:  other-value
```
if [ATTRIBUTE] is specified shows only attributes value  
F.E. ```haldi config show shell```)  
Result:
```bash
zsh
```

### set
```bash
haldi config set [ATTRIBUTE] [VALUE]
```
Updates [ATTRIBUTE] configuration with [VALUE]
