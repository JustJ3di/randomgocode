package main

type label struct {
	key   string
	value string
	next  *label
}

func push(head **label, key string, value string) {

	new := &label{key, value, *head}

	*head = new

}

func pop(head **label) (string, string) {

	curr := *head

	*head = (*head).next

	return curr.key, curr.value
}

func get_tail(head **label) *label {

	curr := *head

	for curr != nil {
		curr = curr.next
	}

	return curr

}

func get_size(head **label) uint32 {

	curr := *head
	var count uint32 = 0
	for curr != nil {
		count++
		curr = curr.next

	}
	return count
}
