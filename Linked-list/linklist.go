package main

import "fmt"

var pl = fmt.Println

type node struct {
	data string
	next *node
}
type link struct {
	head *node
	len  int
}

func Init() *link {
	return &link{}
}

func (l *link) insert(data string) {
	if l.head == nil {
		l.head = &node{data, nil}
		l.len++
		return
	}

	current := l.head
	for current.next != nil {
		pl(current.data)
		current = current.next
	}
	current.next = &node{data, nil}
	l.len++

}

func nil_checker(current *node) {
	if current == nil {
		pl("Head is empty!")
		return
	}
}
func (l *link) display() {
	current := l.head

	// Checks for nil
	nil_checker(current)

	for current != nil {
		pl(current.data)
		current = current.next
	}

}
func (l *link) remove_at(counter int) bool {

	current := l.head
	// Checks for nil
	nil_checker(current)
	count := 1
	var prev *node

	if counter == 0 {
		temp := current.next
		l.head = temp

		return true
	}
	if counter >= l.len {
		return false
	}

	for current != nil {
		prev = current
		if count == counter {
			break
		}
		current = current.next
		count++
	}

	if count < counter {
		return false
	} else {

		temp := prev.next
		prev.next = temp.next

		return true
	}
}

// Reverse reverses the linked list
func (l *link) Reverse() {
	var prev, current, next *node
	current = l.head
	prev = nil
	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}
	l.head = prev
}

func main() {
	list := Init()

	list.insert("Peter")
	list.insert("Osu")
	list.insert("Messy")
	list.insert("Ronson")

	pl()

	pl(list.remove_at(5))

	list.display()

	list.Reverse()

	pl()

	list.display()

}
