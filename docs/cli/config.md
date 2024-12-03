## Config

### Command
```bash
haldi config
```

List or Modify haldi cli user configurations

## Subcommands
### show
lists config value pairs
```bash
haldi config show
```
Result:
```bash
Shell:  zsh
Config2:  value
Config3:  other-value
```
#### Arguments
  1. Attribute: (-a, --attribute) shows only specified attribute (Case insensitive)
```bash
haldi config show -a shell
```
Result:
```bash
zsh
```

### set
sets config value for attribute
```bash
haldi config set
```
#### Arguments
  1. Attribute: (-a, --attribute) attribute to set value (Case insensitive)
```bash
haldi config set -a shell bash
```
<!-- # Supported shells
 - zsh
 - bash
 - dash
 - ksh
 - sh -->
