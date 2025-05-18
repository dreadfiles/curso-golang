package main

import (
	"bufio"
	"fmt"
	"os"
)

type Task struct {
	title       string
	description string
	finished    bool
}

type TaskList struct {
	tasks []Task
}

func (taskList *TaskList) addTask(task Task) {
	taskList.tasks = append(taskList.tasks, task)
}

func (taskList *TaskList) delTask(index int) {
	taskList.tasks = append(taskList.tasks[:index], taskList.tasks[index+1:]...)
}

func (taskList *TaskList) updTask(index int, task Task) {
	taskList.tasks[index] = task
}

func (taskList *TaskList) endTask(index int) {
	taskList.tasks[index].finished = true
}

func (taskList *TaskList) showList() {
	fmt.Println("List of tasks: ", len(taskList.tasks))
	for index, task := range taskList.tasks {
		fmt.Println("***")
		fmt.Println("Index: ", index)
		fmt.Print("Title: ", task.title)
		fmt.Print("Description: ", task.description)
		fmt.Println("Finished: ", task.finished)
		fmt.Println("***")
	}
}

func main() {
	//Tasklist instance
	taskList := TaskList{}
	//IO instance
	ioInput := bufio.NewReader(os.Stdin)
	//Input option variable
	var option int

	for {
		fmt.Println("*** TASK ADMIN ***\n",
			"1. Add Task\n",
			"2. Update Task\n",
			"3. Delete Task\n",
			"4. Finish Task\n",
			"5. Show Tasks\n",
			"6. Exit")

		fmt.Print("Choose an option: ")
		fmt.Scanln(&option)

		switch option {
		case 1:
			var task Task
			fmt.Println("Title: ")
			task.title, _ = ioInput.ReadString('\n')
			fmt.Println("Description: ")
			task.description, _ = ioInput.ReadString('\n')
			taskList.addTask(task)
			fmt.Println("Added Successfully!")
		case 2:
			var index int
			var task Task
			fmt.Println("Index: ")
			fmt.Scanln(&index)
			fmt.Println("Title: ")
			task.title, _ = ioInput.ReadString('\n')
			fmt.Println("Description: ")
			task.description, _ = ioInput.ReadString('\n')
			taskList.updTask(index, task)
			fmt.Println("Updated Successfully!")
		case 3:
			var index int
			fmt.Println("Index: ")
			fmt.Scanln(&index)
			taskList.delTask(index)
			fmt.Println("Deleted Successfully!")
		case 4:
			var index int
			fmt.Println("Index: ")
			fmt.Scanln(&index)
			taskList.endTask(index)
			fmt.Println("Finished Successfully!")
		case 5:
			taskList.showList()
		case 6:
			fmt.Println("See you soon!")
			return
		default:
			fmt.Println("Invalid option!")
		}
	}
}
