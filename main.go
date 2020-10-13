package main

import "./add"

import ("fmt"
		"math/rand"
		"io/ioutil"
		"os"
		"sort"
		"time")
		//"net/http")

// structure
type Books struct {
	title string
	author string
	subject string
	book_id int
}

// Array Declaration
var array = [6] int {78, 34, 11, 23, 89, 199}
// slice
var sliceone = [] string {"Golang", "By", "Google"}
var slicetwo = array[1:5]
var slicethree = make([]int, 4, 7) 
// defer example function
// func display(a int) {  
//     fmt.Println(a)
// }

func rangeExample() {
	sliceRange := []int{8,9,0}
	sum := 0
	for _, val := range sliceRange {
		sum = sum + val
	}
	fmt.Println("slice added ", sum)

	for i, val := range sliceRange {
		if val == 0 {
			fmt.Println("Index of 0 is ", i)
		}
	}

	kvs := map[string]string{"a": "Go", "b": "Lang"}
    for k, v := range kvs {
        fmt.Println(k, v)
	}
	
	for k := range kvs {
        fmt.Println("key:", k)
	}

	for _, v := range kvs {
        fmt.Println("values:", v)
	}
	
	for i, c := range "go" {
        fmt.Println("unicode", i, c)
    }
}

func getMap() {
	Map2 := map[string]int{"key1": 23, "key2": 89}
	for k,v := range Map2 {
		fmt.Println(k,v)
	}
	fmt.Println("Map - ", Map2)
	delete(Map2, "key2")
	fmt.Println("After deletion", Map2)
}
 
 func printBook( book Books ) {
	 fmt.Printf( "Book title : %s\n", book.title);
	 fmt.Printf( "Book author : %s\n", book.author);
	 fmt.Printf( "Book subject : %s\n", book.subject);
	 fmt.Printf( "Book book_id : %d\n", book.book_id);
 }

// channel example for multi-tasking
func displayChannel(ch chan int) {
	// time.Sleep(5 * time.Second)
	fmt.Println("Inside display()")
	ch <- 1234
	close(ch)
}

// main function which gets executed
func main () {
	fmt.Println("Any number from 1 - 100", rand.Intn(100))

	// defer display(15) // gets called in the last when the execution of main function is done.
	// defer display(22)
	// defer display(33)
	// defer display(44)

	fmt.Println("Defer will be called now!")

	ch := make(chan int) // creating a channel
	go displayChannel(ch)
	x := <-ch // assigning values of ch to x
	fmt.Println("Inside main()")
	fmt.Println("Printing x in main() after taking from channel:", x)
	
	var Book1 Books
	var Book2 Books 

	/* book 1 specification */
	Book1.title = "Go Programming"
	Book1.author = "Mahesh Kumar"
	Book1.subject = "Go Programming Tutorial"
	Book1.book_id = 6495407

	/* book 2 specification */
	Book2.title = "Telecom Billing"
	Book2.author = "Zara Ali"
	Book2.subject = "Telecom Billing Tutorial"
	Book2.book_id = 6495700	

	printBook(Book1)
	printBook(Book2)

	// Reading data.txt file 
	data, err := ioutil.ReadFile("data.txt")
	if err != nil {
		fmt.Println("File reading error", err)
        return
	} 
	fmt.Println("Contents of Data.txt:", string(data))

	// Writing to a file 
	f, err := os.Create("write-file.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    l, err := f.WriteString("Write Line one")
    if err != nil {
        fmt.Println(err)
        f.Close()
        return
    }
    fmt.Println(l, "bytes written")
	err = f.Close()
    if err != nil {
        fmt.Println(err)
        return
	}

	// array examples
	fmt.Println("Array is -> ", array)
	fmt.Println("Length of Array is -> ", len(array))
	otherArray := [5] int {1,2,3,4,5}
	fmt.Println("Other array is -> ", otherArray)
	twoDimensionalArray := [2][3] string {{"a","b","c"},{"d","e","f"}}
	fmt.Println("Two dimensional Array -> ", twoDimensionalArray)
	structArray := [2] Books {Book1, Book2}
	fmt.Println("struct Array -> " ,structArray)

	// slices
	fmt.Println("slicesss 1", sliceone)
	fmt.Println("slicess 2", slicetwo)
	fmt.Println("slicess 3 ", slicethree)
	// appending sliceone
	sliceone = append(sliceone, "Is", "Great!")
	fmt.Println("after appending sliceone ", sliceone)

	sort.Strings(sliceone)
	sort.Ints(slicetwo)

	fmt.Println("\nAfter sorting:") 
    fmt.Println("Slice 1: ", sliceone) 
	fmt.Println("Slice 2: ", slicetwo) 

	fmt.Println("Printing from other package ", add.Morning)
	getMap()
	rangeExample()
	Time := time.now()
	fmt.Println("Printing time -> ", Time)
}