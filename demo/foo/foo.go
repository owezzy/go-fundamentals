package foo

import "fmt"

func Greet() {
	fmt.Println("Hello Word")
}

func IsOdd(n int) bool {
	return n%2 == 1
}

type User struct {
	Name string
	Age  int
}

func GetUser(id int) (*User, error) {
	return nil, fmt.Errorf("Not Implemented")
}
