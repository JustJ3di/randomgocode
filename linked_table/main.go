package main

import (
	"fmt"
)

func main() {
	var head *label

	head = nil

	push(&head, "ciao", "value")

	push(&head, "ciao", "valss")

	push(&head, "olaaa", "sss1")

	key, value := pop(&head)
	push(&head, "olasss", "dasa")

	fmt.Println(head.key, head.value)

	fmt.Println(key, value)

	fmt.Println(get_size(&head))

}
