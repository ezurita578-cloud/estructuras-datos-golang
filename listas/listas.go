package main

import "fmt"

type Task struct {
	Title string
	Done  bool
}

type Node struct {
	value Task
	next  *Node
}

type TaskList struct {
	head *Node
}

func (l *TaskList) Add(title string) {
	newNode := &Node{value: Task{Title: title}}
	if l.head == nil {
		l.head = newNode
		return
	}
	temp := l.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = newNode
}

// Marcar tarea como completada
func (l *TaskList) Complete(title string) {
	temp := l.head
	for temp != nil {
		if temp.value.Title == title {
			temp.value.Done = true
			return
		}
		temp = temp.next
	}
}

// Eliminar tarea por nombre
func (l *TaskList) Remove(title string) {
	if l.head == nil {
		return
	}
	if l.head.value.Title == title {
		l.head = l.head.next
		return
	}
	temp := l.head
	for temp.next != nil {
		if temp.next.value.Title == title {
			temp.next = temp.next.next
			return
		}
		temp = temp.next
	}
}

// Contar tareas
func (l *TaskList) Count() int {
	count := 0
	temp := l.head
	for temp != nil {
		count++
		temp = temp.next
	}
	return count
}

func (l *TaskList) Print() {
	temp := l.head
	for temp != nil {
		fmt.Println("-", temp.value.Title, "| Done:", temp.value.Done)
		temp = temp.next
	}
}

func main() {
	list := TaskList{}
	list.Add("Aprender Go")
	list.Add("Hacer tarea")
	list.Add("Estudiar estructuras")

	list.Complete("Aprender Go")
	list.Remove("Hacer tarea")

	fmt.Println("Tareas:", list.Count())
	list.Print()
}
