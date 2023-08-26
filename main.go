package main

import (
	"fmt"
	"time"
)

type Task struct {
	ID        int
	Title     string
	Assignee  string
	Deadline  time.Time
	Completed bool
}

type Scheduler struct {
	Tasks []Task
}

func (s *Scheduler) AddTask(title, assignee string, deadline time.Time) {
	taskID := len(s.Tasks) + 1
	task := Task{
		ID:        taskID,
		Title:     title,
		Assignee:  assignee,
		Deadline:  deadline,
		Completed: false,
	}
	s.Tasks = append(s.Tasks, task)
	fmt.Println("Task added:", task)
}

func (s *Scheduler) CompleteTask(taskID int) {
	for i, task := range s.Tasks {
		if task.ID == taskID {
			s.Tasks[i].Completed = true
			fmt.Println("Task completed:", task)
			return
		}
	}
	fmt.Println("Task not found")
}

func main() {
	scheduler := Scheduler{}

	scheduler.AddTask("Design UI", "Alice", time.Now().Add(48*time.Hour))
	scheduler.AddTask("Write API Docs", "Bob", time.Now().Add(72*time.Hour))
	scheduler.AddTask("Test Feature", "Charlie", time.Now().Add(96*time.Hour))

	fmt.Println("Current Tasks:")
	for _, task := range scheduler.Tasks {
		fmt.Printf("ID: %d | Title: %s | Assignee: %s | Deadline: %s | Completed: %v\n",
			task.ID, task.Title, task.Assignee, task.Deadline.Format("2006-01-02 15:04:05"), task.Completed)
	}

	scheduler.CompleteTask(2)

	fmt.Println("\nUpdated Tasks:")
	for _, task := range scheduler.Tasks {
		fmt.Printf("ID: %d | Title: %s | Assignee: %s | Deadline: %s | Completed: %v\n",
			task.ID, task.Title, task.Assignee, task.Deadline.Format("2006-01-02 15:04:05"), task.Completed)
	}
}
