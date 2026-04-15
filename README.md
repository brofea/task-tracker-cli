# Task Tracker CLI

[English] | [中文](README.zh-CN.md)

A simple command-line application to track tasks, built with Go. It allows you to add, list, and manage your tasks efficiently.

This is an backend exercise project from [roadmap.sh](https://roadmap.sh/projects/task-tracker).

## How to Use

```bash
# Adding a new task
./task add "Buy groceries"
# Output: Task added successfully (ID: 1)

# Updating and deleting tasks
./task update 1 "Buy groceries and cook dinner"
./task delete 1

# Marking a task as in progress or done
./task mark-in-progress 1
./task mark-done 1

# Listing all tasks
./task list

# Listing tasks by status
./task list done
./task list todo
./task list in-progress

```

If you use Windows, replace `./task` with `task.exe`.

If you use macOS and the binary from the release, replace `./task` with `./task-darwin` or rename the file.

## Quick Start

1. `git clone` the repository
2. Run `go build -o task` to build the application
3. Use the commands as shown above to manage your tasks