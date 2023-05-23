package main

import (
	"fmt"
	"log"

	// "github.com/gobuffalo/flect"
	"github.com/gofrs/uuid/v5"
	"github.com/owezzy/go-fundamentals/demo/foo"
	"foo"
)

func main(){
	// s := "Hello World Owezzy"
	// d:= flect.Dasherize(s)

	// fmt.Println(s, d)

	// u, err := uuid.NewV4()

	// if err != nil{
		// log.Fatal(err)
	// }

	// fmt.Println(u)

	foo.Greet()
}