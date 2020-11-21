# lfl (Links for Lectures)
lfl (Links for Lectures) is a CLI which opens the link of the lecture of your choice
in zoom. You add the subjects' info in your config file and you are done.

## Requirements
- Go
- Git

You can remove them later.

## Install

### Basic installation
Run:
```
$ sudo sh -c "$(curl -fsSL https://raw.githubusercontent.com/Stavrospanakakis/lfl/main/scripts/install.sh)"
```

### Manual Installation
Run:
```
$ wget https://raw.githubusercontent.com/Stavrospanakakis/lfl/main/scripts/install.sh
$ sudo sh install.sh
```

## Usage
### Available options

**Manual of the CLI**
```
$ lfl --help 
```

**Show all the ids and open the lecture of your choice**
```
$ lfl
```

**Open the link of the lecture of your choice**
```
$ lfl -l <lecture-id> 
```

**Add your own configuration file**
```
$ lfl --config <config-file-location> 
```

## Configuration
Default configuration file is located at $HOME/.lfl.json. Add the lectures of your choice and you are done.

```
{
    "1" : {
        "name": "Example Lecture 1",
        "link": "https://zoom.us/j/<rest-of-the-link>"
    },
    "2": {
        "name": "Example Lecture 2",
        "link": "https://zoom.us/j/<rest-of-the-link>"
    },
    "totalSum":"2"
}
```


## Uninstall
Run:
```
$ wget https://raw.githubusercontent.com/Stavrospanakakis/lfl/main/scripts/uninstall.sh
$ sudo sh uninstall.sh
```
## Example addition