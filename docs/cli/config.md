## Config

### Command
```bash
anvil config [SUBCOMMAND]
```

List or Modify anvil cli user configurations

## Available configurations

- shell: sets which shell should be used for running anvil aliases
  - Available values are zsh, bash, dash, ksh, sh

## Subcommands
### show
```bash
anvil config show [ATTRIBUTE]
```
If [ATTRIBUTE] is not specified lists all anvil cli configurations attributes with their values  
Result:
```bash
Shell:  zsh
Config2:  value
Config3:  other-value
```
if [ATTRIBUTE] is specified shows only attributes value  
F.E. ```anvil config show shell```)  
Result:
```bash
zsh
```

### set
```bash
anvil config set [ATTRIBUTE] [VALUE]
```
Updates [ATTRIBUTE] configuration with [VALUE]
