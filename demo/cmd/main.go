package main

import (
	"fmt"
	// "log"

	// "github.com/gobuffalo/flect"
	// "github.com/gofrs/uuid/v5"
	"github.com/owezzy/go-fundamentals/foo"
)

func main() {
	// s := "Hello World Owezzy"
	// d:= flect.Dasherize(s)

	// fmt.Println(s, d)

	// u, err := uuid.NewV4()

	// if err != nil{
	// log.Fatal(err)
	// }

	// fmt.Println(u)

	// var i int
	// var f float64
	// var b bool
	// var s string

	//  i = 43
	//  f = 3.14
	//  b = true
	//  s = "Hello World"

	//  i := 43
	//  f := 3.14
	//  b := true
	//  s := "Hello World"

	//  i, f, b, s := 43 ,3.14, true ,"Hello World"

	 i, f, b, s := Values()
 
	fmt.Printf("var i %T = %v\n", i, i)
	fmt.Printf("var i %T = %v\n", f, f)
	fmt.Printf("var i %T = %v\n", b, b)
	fmt.Printf("var i %T = %v\n", s, s)
	
	foo.Greet()


}

func Values()(int, float64, bool, string){
	return 	43 ,3.14, true ,"Hello World"

}