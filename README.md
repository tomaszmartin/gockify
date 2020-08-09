# Gockify

CLI App for Clockify written in GO. 
Cleverly named Gockify :P

## Usage

First You need to initialize the app.

```
$ glockify init
Insert API key: ********************
Initialized. Available workspaces:
    (1) First
    (2) Second
Choose default workspace: 1
Workspace [First] set as defualt.
```

This app is extremely simple. It allows to start a task by passing project and task.
It tries to match project and task based on name and if they don't exists it creates them.

```
$ glockify start [project] [task]
Created [task] in [project].
Task [task] running for 00:05. Total 02:10 tracked today.
Press e to end task.
```