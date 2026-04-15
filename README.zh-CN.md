# Task Tracker CLI

[English](README.md) | [中文]

一个简单的命令行任务跟踪程序，使用 Go 语言编写。它允许您高效地添加、列出和管理任务。

这是一个来自 [roadmap.sh](https://roadmap.sh/projects/task-tracker) 的后端练习项目。

## 功能

```bash
# 添加一个新任务
./task add "Buy groceries"
# 更新一个任务
./task update 1 "Buy groceries and cook dinner"
# 删除一个任务
./task delete 1
# 标记任务状态
./task mark-in-progress 1
./task mark-done 1
# 列出所有任务
./task list
# 列出特定状态的任务
./task list done
./task list todo
./task list in-progress
```

如果您使用 Windows 系统，请将上述命令中的 `./task` 替换为 `task.exe`

如果您使用 macOS 和 release 中的二进制文件，请将上述命令中的 `./task` 替换为 `./task-darwin` 或者重命名文件。

## 快速开始

1. 克隆仓库：`git clone`
2. 构建应用程序：`go build -o task`，在 Windows 上使用 `go build -o task.exe`
3. 使用上述命令管理您的任务