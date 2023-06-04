package main

import (
	"fmt"
	"strings"

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

	//  i, f, b, s := Values()

	// fmt.Printf("var i %T = %v\n", i, i)
	// fmt.Printf("var i %T = %v\n", f, f)
	// fmt.Printf("var i %T = %v\n", b, b)
	// fmt.Printf("var i %T = %v\n", s, s)

	foo.Greet()

	names := [4]string{"owen", "owezzy", "kunta", "kinte"}
	namesSlice := []string{"owen", "owezzy", "kunta", "kinte"}

	fmt.Printf("Array: %[1]T:%[1]v\n", names)
	fmt.Printf("Slice: %[1]T:%[1]v\n", namesSlice)

	var a [5]int
	var b [4]string
	var c [3]bool

	fmt.Printf("%#v\n", a)
	fmt.Printf("%#v\n", b)
	fmt.Printf("%#v\n", c)

	// slices

	var nn []string
	nn = append(nn, "oguyo")

	fmt.Println(nn)

	nn = append(nn, "ting'amalo", "manyoPesa")

	fmt.Println(nn)

	jina := []string{}

	fmt.Println("len:", len(jina))
	fmt.Println("cap:", cap(jina))

	jina = append(jina, "ting'amalo")

	fmt.Println("len:", len(jina))
	fmt.Println("cap:", cap(jina))

	jina = append(jina, "oguyo")

	fmt.Println("len:", len(jina))
	fmt.Println("cap:", cap(jina))

	d := []string{} // declare initialize
	var e []string  // declare only

	fmt.Println(d)
	fmt.Println(e)

	f := make([]string, 1, 3) // use make function to declare slice

	fmt.Printf("%#v\n", f)
	fmt.Println("len:", len(f))
	fmt.Println("cap:", cap(f))

	letters := []string{"a", "b", "c", "d", "e", "f", "g"}
	fmt.Println(letters)

	// get 3 elements start from 3rd
	// fmt.Println(letters[2:5])

	//functionaliy equivalent
	fmt.Println(letters[4:len(letters)])
	fmt.Println(letters[4:])

	//functionaliy equivalent
	fmt.Println(letters[0:4])
	fmt.Println(letters[:4])

	// subsets

	// subset := letters[:2 ]

	// fmt.Println(subset)

	// using copy function to create subset

	subset := make([]string, 3)

	copy(subset, letters[:2])

	fmt.Println(subset)

	for i, g := range subset {
		subset[i] = strings.ToUpper(g)
	}

	fmt.Println(subset)

	fmt.Println(letters)

	// convert array to slice
	slicesOnly(letters)
	slicesOnly(names[:])

	tt := [4]string{"owen", "owezzy", "kunta", "kinte"}

	// ITERATIONS
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// using both index and value from range
	for i, n := range tt {
		fmt.Printf("%d %s\n", i, n)
	}

	// when you just need the index
	for i := range tt {
		fmt.Printf("%d %s\n", i, tt[i])
	}

	// use continue to skip iteration
	for i := 0; i < 10; i++ {
		if i == 3 {
			continue
		}

		fmt.Println(i)
	}


	// use break to stop iteration
	for i := 0; i < 10; i++ {
		if i == 3 {
			break
		}

		fmt.Println(i)
	}
}

func Values() (int, float64, bool, string) {
	return 43, 3.14, true, "Hello World"

}

func slicesOnly(letters []string) {
	for _, letter := range letters {
		fmt.Println(letter)
	}
}
