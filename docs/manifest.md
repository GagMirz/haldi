## Manifest file

### Example manifest
With this example manifest yuo would be able to run npm in your projects backend directory regardless where are you in your terminal session by running ```haldi run haldi be```

```json
{
    "name": "haldi",
    "aliases": [
        {
            "command": "npm",
            "name": "be",
            "path": "/backend"
        }
    ]
}
```
- name is project name, which would be used to refer manifest via cli
- aliases is array of commands you can later run
- alias:command command which will be triggered, after command any argument provided would be passed to command
- alias:name name of command which would be used to refer command via cli
- alias:path relative path to manifest file, so command can be run in subdirectory of project
