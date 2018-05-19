package model

// TodoList TodoList struct
type TodoList struct {
	UID  string
	Todo *[]Todo
}

// Todo Todo struct
type Todo struct {
	Name   string
	Status int
}
