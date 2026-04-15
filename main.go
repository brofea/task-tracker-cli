package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
type TaskList struct {
	NextID int    `json:"next_id"`
	Tasks  []Task `json:"tasks"`
}

var taskList TaskList

func openFile() {
	// 检查 tasks.json 是否存在，如果不存在则创建
	if _, err := os.Stat("tasks.json"); os.IsNotExist(err) {
		emptyList := TaskList{NextID: 1, Tasks: []Task{}}
		data, _ := json.Marshal(emptyList)
		os.WriteFile("tasks.json", data, 0644)
	}
	// 读取并解析 tasks.json 文件
	rawData, _ := os.ReadFile("tasks.json")
	// 尝试解析 JSON 数据，如果失败则报错
	if err := json.Unmarshal(rawData, &taskList); err != nil {
		panic(err)
	}
}

func saveFile() {
	// 将更新后的任务列表写回 tasks.json 文件
	data, _ := json.Marshal(taskList)
	os.WriteFile("tasks.json", data, 0644)
}

func commandHelp() {
	println("Usage:")
	println("  help - Show this help message")
	println("  add <description> - Add a new task")
	println("  update <id> <description> - Update the description of a task")
	println("  delete <id> - Delete a task")
	println("  mark-done <id> - Mark a task as done")
	println("  mark-todo <id> - Mark a task as todo")
	println("  mark-in-progress <id> - Mark a task as in progress")
	println("  list [done|todo|in-progress] - List tasks")
	println("  list - List all tasks")
}

func listTasks(filter string) {
	if filter != "all" && filter != "done" && filter != "todo" && filter != "in-progress" {
		println("Invalid filter. Use 'done', 'todo', 'in-progress', or 'all'.")
		return
	}
	openFile()
	for _, task := range taskList.Tasks {
		if filter == "all" || task.Status == filter {
			fmt.Printf("[%d] %s\t%s\tCreated: %s\tUpdated: %s\n",
				task.ID, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
		}
	}
	if len(taskList.Tasks) == 0 {
		println("No tasks found")
	}
}

func addTask(description string) {
	if description == "" {
		println("Description cannot be empty")
		return
	}
	openFile()
	now := time.Now().Format("2006-01-01 00:00:00")
	task := Task{
		ID:          taskList.NextID,
		Description: description,
		Status:      "todo",
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	taskList.Tasks = append(taskList.Tasks, task)
	taskList.NextID++
	saveFile()
}

func foundTaskByID(id int) int {
	var foundIndex int = -1
	for i, task := range taskList.Tasks {
		if task.ID == id {
			foundIndex = i
			break
		}
	}
	return foundIndex
}

func updateTask(idStr, description string) {
	if description == "" {
		println("Description cannot be empty")
		return
	}
	openFile()
	var id int
	fmt.Sscanf(idStr, "%d", &id)
	var index int = foundTaskByID(id)
	if index == -1 {
		println("Task not found")
		return
	}
	taskList.Tasks[index].Description = description
	taskList.Tasks[index].UpdatedAt = time.Now().Format("2006-01-01 00:00:00")
	saveFile()
}

func deleteTask(idStr string) {
	if idStr == "" {
		println("ID cannot be empty")
		return
	}
	openFile()
	var id int
	fmt.Sscanf(idStr, "%d", &id)
	var index int = foundTaskByID(id)
	if index == -1 {
		println("Task not found")
		return
	}
	taskList.Tasks = append(taskList.Tasks[:index], taskList.Tasks[index+1:]...)
	saveFile()
}

func markTask(idStr, status string) {
	if idStr == "" {
		println("ID cannot be empty")
		return
	}
	var validId bool = true
	for _, c := range idStr {
		if c < '0' || c > '9' {
			validId = false
			break
		}
	}
	if !validId {
		println("ID must be a number")
		return
	}
	openFile()
	var id int
	fmt.Sscanf(idStr, "%d", &id)
	var index int = foundTaskByID(id)
	if index == -1 {
		println("Task not found")
		return
	}
	taskList.Tasks[index].Status = status
	taskList.Tasks[index].UpdatedAt = time.Now().Format("2006-01-01 00:00:00")
	saveFile()
}

func main() {
	// 解析命令行参数
	if len(os.Args) < 2 {
		listTasks("all")
		return
	}
	switch os.Args[1] {
	case "help":
		commandHelp()
	case "add":
		if len(os.Args) != 3 {
			println("Usage: add <description>")
			break
		}
		addTask(os.Args[2])
	case "update":
		if len(os.Args) != 4 {
			println("Usage: update <id> <description>")
			break
		}
		updateTask(os.Args[2], os.Args[3])
	case "delete":
		if len(os.Args) != 3 {
			println("Usage: delete <id>")
			break
		}
		deleteTask(os.Args[2])
	case "mark-done":
		if len(os.Args) != 3 {
			println("Usage: mark-done <id>")
			break
		}
		markTask(os.Args[2], "done")
	case "mark-todo":
		if len(os.Args) != 3 {
			println("Usage: mark-todo <id>")
			break
		}
		markTask(os.Args[2], "todo")
	case "mark-in-progress":
		if len(os.Args) != 3 {
			println("Usage: mark-in-progress <id>")
			break
		}
		markTask(os.Args[2], "in-progress")
	case "list":
		if len(os.Args) > 3 {
			println("Usage: list [done|todo|in-progress]")
			break
		}
		if len(os.Args) == 3 {
			listTasks(os.Args[2])
		} else {
			listTasks("all")
		}
	default:
		println("Unknown command")
	}
}
