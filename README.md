# goggl

Manage your work time from the terminal !

## Get goggl

You can either:

* clone this repo and build from sources (Go required):

```
$ git clone https://github.com/eze-kiel/goggl.git
$ cd goggl
$ go build
```

* get it from the [releases](https://github.com/eze-kiel/goggl/releases) (soon)

## Usage

`goggl` helps you manage your work sessions from the terminal easily.
When a session is created, a JSON file is created under `$HOME/.goggl/running/`.
When the session is stopped, the file is moved under `$HOME/.goggl/history/`for later
usage.

```
Usage:
  goggl [command]

Available Commands:
  completion  Generate completion script
  help        Help about any command
  init        Initialize goggl files
  start       Start a new working session
  stop        Stop a running working session

Flags:
  -h, --help   help for goggl

Use "goggl [command] --help" for more information about a command.
```

### Initialization

Before using `goggl`, you need to intialize it:

```
$ goggl init
INFO[0000] creating goggl files at /home/ezekiel/.goggl 
INFO[0000] initialization has been successful !
```

### Create a work session

A session can be created by typing:

```
$ goggl start
INFO[0000] work session untagged_2021-01-27_18:26:39 created successfully ! GLHF
```

You can change the prefix of your sessions with the flag `-t` or `--tag`:

```
$ goggl start -t coding
INFO[0000] work session coding_2021-01-27_18:27:39 created successfully ! GLHF
```

If you look in `$HOME/.goggl/runnnig`, you'll see the running sessions:

```
$ tree .goggl/running/
.goggl/running/
└── untagged_2021-01-27_18:26:39.json
```

### Stop a work session

Once you're done with your task, you can stop and archive it:

```
$ goggl stop coding_2021-01-27_18:27:39.json 
INFO[0000] work session duration: 46s                   
INFO[0000] archived coding_2021-01-27_18:27:39 into history
```

You can now check the content of `$HOME/.goggl/history` to see the file:

```
$ tree .goggl/history/
.goggl/history/
└── 2021
    └── January
        └── coding_2021-01-27_18:27:39.json
```

## The content of the JSON file

The files that are created are really simples, and follow this structure:

```json
{
  "name": "coding_2021-01-27_18:27:39",
  "tag": "coding",
  "start_time": "2021-01-27_18:27:39",
  "end_time": "2021-01-27_18:28:25",
  "duration": "46s"
}
```