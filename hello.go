package main

import "fmt"
import "errors"

//import "os"

type Point struct {
	X, Y int
	Name string
}

func (p Point) Invert() Point {
	return Point{p.Y, p.X, p.Name}
}

func pointTest() {
	point := Point{5, 6, "starting point"}
	fmt.Println(point)

	points := make(map[string]Point)
	points["top"] = Point{1, 2, "starting point"}
	points["bottom"] = Point{3, 4, "finishing point"}

	value, present := points["top"]
	fmt.Println(value, present)
	fmt.Println(value.Invert())
}

var (
	hello = "Hello"
	world = "World"
)

func helloWorld() {
	fmt.Println(swap(hello, world))
}

func swap(a, b string) (string, string) {
	return b, a
}

func fib(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func fibTest() {
	fmt.Println(fib(10))
}

func arrayTest() {
	ints := []int{0, 1, 2, 3, 4}
	fmt.Println(ints)

	for i, v := range ints {
		fmt.Printf("i = %d, v = %d\n", i, v)
	}
}

func boolTest() {
	if true {
		fmt.Println("no doubt")
	}
}

func switchTest() {
	username := "scott"
	switch username {
	case "admin", "player number one":
		fmt.Println("is admin")
	default:
		fmt.Println("is not admin")
	}

	switch {
	case 1 > 0:
		fmt.Println("1 > 0")
	case 1 < 0:
		fmt.Println("1 < 0")
	}
}

func Salary(username string) (salary int, err error) {
	salary = 666
	err = errors.New("Error description")
	return
}

func foodMaker(tags []string) func(string) {
	return func(name string) {
		fmt.Println("Named with " + name)
		fmt.Println(tags) // closure
	}
}

func callMe(name string, maker func(string)) {
	maker(name)
}

func functionTest() {
	bananaBreadMaker := foodMaker([]string{"flour", "soda", "salt", "butter", "eggs", "banana", "vanilla"})

	callMe("Banana bread", bananaBreadMaker)

	s, e := Salary("patrick")
	fmt.Println(s, e)
}

func concurrencyTest() {
	tasks := make(chan string, 2)
	results := make(chan string, 2)
	go func(t chan string) {
		for {
			task := <-t
			results <- task + " done"
		}
	}(tasks)

	tasks <- "task1"
	tasks <- "task2"
	fmt.Println(<-results)
	fmt.Println(<-results)
}

func ioTest() {
	/*
		file, err := os.Open("test")
		defer file.close()
	*/
}

func main() {
	helloWorld()
	fibTest()
	pointTest()
	arrayTest()
	boolTest()
	switchTest()
	functionTest()
	ioTest()
}
