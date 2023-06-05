package main

import (
	"fmt"
	"os"
	"sort"
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

	// MAPS
	fmt.Println("MAPS")

	// declare and initialize map is recommended way
	user := map[string]string{"owe": "owe@gmail.com", "que": "que@gmail.com"}

	fmt.Println(user)

	watu := map[User]string{}

	mtu := User{ID: 1, Name: "Owezzy"}
	watu[mtu] = "owezzy@example.com"

	fmt.Println(watu)

	// interation order is random
	for key, value := range user {
		fmt.Println(key, value)
	}

	// when you just need the key
	for key := range user {
		fmt.Println(key, user[key])
	}

	// delete value using its key

	delete(user, "owe")
	fmt.Println("after delete", user)

	// to check value exits and hanlde err

	data := map[int]string{}

	data[1] = "Hello, Word"

	value, ok := data[1]

	if !ok {
		fmt.Printf("key %d not found\n", 10)
		os.Exit(1)
	}

	fmt.Printf("%q\n", value)

	// using Zero value in maps
	counts := map[string]int{}

	sentence := "The quick fox jumps over the lazy dog"

	words := strings.Fields(strings.ToLower(sentence))

	for _, w := range words {
		// counts[w] = counts[w] + 1 or
		counts[w]++
	}

	fmt.Println(counts)

	months := map[int]string{
		1: "Jan",
		2: "Feb",
		3: "Mar",
		4: "Apr",
		5: "May",
		6: "June",
	}

	keys := make([]int, 0, len(months))

	for k := range months {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	fmt.Printf("keys: %+v\n", keys)

	// if statements
	greet := true

	if greet {
		fmt.Println("Hello")
	} else {
		fmt.Println("Goodbye")
	}

	if 1 == 1 {
		fmt.Println("1 == 1")
	}

	if 1 != 2 {
		fmt.Println("1 != 2")
	}
	if 2 > 1 {
		fmt.Println("2 > 1")
	}
	if 1 < 2 {
		fmt.Println("1 < 2")
	}
	if 1 >= 1 {
		fmt.Println("1 >= 1")
	}
	if 1 <= 1 {
		fmt.Println("1 <= 1")
	}

	// switch statement
	mwezi := 3

	switch mwezi {
	case  1:
		fmt.Println("Jan")
	case  2:
		fmt.Println("Feb")
	case  3:
		fmt.Println("March")
	case 4:
		fmt.Println("April")
	case  5:
		fmt.Println("May")

	}

	recommendActivity(19)
	recommendActivity(45)
	recommendActivity(90)
}

type User struct {
	ID   int
	Name string
}

func Values() (int, float64, bool, string) {
	return 43, 3.14, true, "Hello World"

}

func slicesOnly(letters []string) {
	for _, letter := range letters {
		fmt.Println(letter)
	}
}

func recommendActivity(temp int){
	fmt.Printf("It is %d degrees out. You could", temp)

	switch{
	case temp <= 32:
		fmt.Print(" go ice skating")
		fallthrough
	case temp >=45 && temp <= 90:
		fmt.Print(" go jogging")
		fallthrough
	case temp >= 80:
		fmt.Print(" go swimming")
		fallthrough
	default:
		fmt.Print(" or just stay at home. \n")
	}


}
