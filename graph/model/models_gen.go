// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Todo struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type TodoInput struct {
	Title string `json:"title"`
}
