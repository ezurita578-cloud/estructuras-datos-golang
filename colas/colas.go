package main

import "fmt"

type Client struct {
	Name  string
	IsVIP bool
}

type Queue struct {
	items []Client
}

func (q *Queue) Enqueue(c Client) {
	if c.IsVIP {
		q.items = append([]Client{c}, q.items...)
	} else {
		q.items = append(q.items, c)
	}
}

func (q *Queue) Dequeue() Client {
	if len(q.items) == 0 {
		return Client{}
	}
	first := q.items[0]
	q.items = q.items[1:]
	return first
}

func (q *Queue) Peek() Client {
	if len(q.items) == 0 {
		return Client{}
	}
	return q.items[0]
}

type Bank struct {
	line Queue
}

func (b *Bank) Arrive(name string, isVIP bool) {
	b.line.Enqueue(Client{Name: name, IsVIP: isVIP})
}

func (b *Bank) Attend() {
	client := b.line.Dequeue()
	if client.Name != "" {
		fmt.Println("Atendiendo a:", client.Name)
	} else {
		fmt.Println("No hay clientes")
	}
}

func (b *Bank) Next() {
	client := b.line.Peek()
	if client.Name != "" {
		fmt.Println("Siguiente en fila:", client.Name)
	} else {
		fmt.Println("No hay clientes en fila")
	}
}

func main() {
	bank := Bank{}
	bank.Arrive("Ana", false)
	bank.Arrive("Luis", false)
	bank.Arrive("Pedro", true)
	bank.Next()
	bank.Attend()
	bank.Attend()
	bank.Attend()
}
