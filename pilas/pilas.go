package main

import "fmt"

type Stack struct {
	items []string
}

func (s *Stack) Push(url string) {
	s.items = append(s.items, url)
}

func (s *Stack) Pop() string {
	if len(s.items) == 0 {
		return ""
	}
	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return last
}

type Browser struct {
	current string
	history Stack
	forward Stack
}

func (b *Browser) Visit(url string) {
	if b.current != "" {
		// Limitar historial a 10
		if len(b.history.items) >= 10 {
			b.history.items = b.history.items[1:]
		}
		b.history.Push(b.current)
	}
	b.current = url
	// Al visitar nueva pagina se limpia el forward
	b.forward.items = []string{}
}

func (b *Browser) Back() {
	prev := b.history.Pop()
	if prev != "" {
		b.forward.Push(b.current)
		b.current = prev
	}
}

// Forward: ir hacia adelante
func (b *Browser) Forward() {
	next := b.forward.Pop()
	if next != "" {
		b.history.Push(b.current)
		b.current = next
	}
}

func main() {
	b := Browser{}
	b.Visit("google.com")
	b.Visit("github.com")
	b.Visit("stackoverflow.com")

	fmt.Println("Actual:", b.current)
	b.Back()
	fmt.Println("Despues de back:", b.current)
	b.Forward()
	fmt.Println("Despues de forward:", b.current)
}
