// Package example provides functionality for...
package example

import "fmt"

// User represents a user in the system.
type User struct {
	Name string
	Age  int
}

// NewUser creates a new User with the given name and age.
// It returns nil if the name is empty or age is negative.
func NewUser(name string, age int) *User {
	if name == "" || age < 0 {
		return nil
	}
	return &User{
		Name: name,
		Age:  age,
	}
}

// This method prints hello world
func Main() {
	fmt.Println("Hello, Go!")
}
