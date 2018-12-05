package main

import "os"
import "fmt"
import "time"
import "unsafe"
import "encoding/json"

type Circle struct {
	x  int   // kkk
	y int // jjj
	Radius int
}

type ArrayStruct struct{
	value [10]int
}

type SliceStruct struct{
	value []int
}

func main(){

	"abcdedddd"

	p := fmt.Println
	p(sum(1, 2, 3, 4))

	testTime()

	os.Exit(200)

	testSlice()
	testMap()
	fmt.Print("hello", ",", "world!\n")

	var c Circle = Circle{
		x : 100,
		y : 100,
		Radius : 50,
	}

	var d Circle = Circle{}
	var e Circle = Circle {100, 100, 50}
	var f = &e

	fmt.Printf("%+v\n", c)
	fmt.Printf("%+v\n", d)
	fmt.Printf("%+v\n", e)
	fmt.Printf("%+v\n", f)

	var as = ArrayStruct{[...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}
	var ss = SliceStruct{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}

	fmt.Println(unsafe.Sizeof(as), unsafe.Sizeof(ss))

	changeByValue(e)
	fmt.Printf("%+v\n", e)

	ChangeByPointer(f)
	fmt.Printf("%+v\n", f)
}

func changeByValue(c Circle){
	c.Radius *= 2
}

func ChangeByPointer(cf *Circle){
	cf.Radius *= 2
}

type Something struct{
	Values []int
}
// Nil slice vs. empty slice
func testSlice() {
	var s1 []int
	var s2 = []int{}
	var s3 = make([]int, 0)
	var s4 = *new([]int)

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)

	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)
	fmt.Println(s3 == nil)
	fmt.Println(s4 == nil)

	var a1 = *(*[3]int)(unsafe.Pointer(&s1))
	var a2 = *(*[3]int)(unsafe.Pointer(&s2))
	var a3 = *(*[3]int)(unsafe.Pointer(&s3))
	var a4 = *(*[3]int)(unsafe.Pointer(&s4))

	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)

	var sth1 = Something{}
	var sth2 = Something{[]int{}}

	bs1, _ := json.Marshal(sth1)
	bs2, _ := json.Marshal(sth2)

	fmt.Println(string(bs1))
	fmt.Println(string(bs2))

	var TestSlice = []string{"A", "B", "C"}

	
	fmt.Println("len & cap of slice:", len(TestSlice), cap(TestSlice))

	fmt.Println("Iterate and print every slice element using []")
	for i := 0; i<len(TestSlice); i++{
		fmt.Println(TestSlice[i])
	}

	TestSlice2 := []string{}

	// _ must be used, otherwise compiler will warn below statement "evaluated but not used"
	_ = append(TestSlice, "D")
	TestSlice2 = append(TestSlice, "D")
	fmt.Println("len & cap of slice2:", len(TestSlice2), cap(TestSlice2))
	fmt.Println(TestSlice[1])
	fmt.Println("Iterate and print every slice element using range")
	for index, value := range TestSlice2{
		fmt.Println(index, ". ", value)
	}

	fmt.Println("----------")
}

func testMap(){
	var m1 = map[string]int{
		"A" : 90,
		"B" : 80,
		"C" : 60,
	}
	fmt.Println(m1)

	// size of the map pointer
	fmt.Println(unsafe.Sizeof(m1))

	// length of the map
	fmt.Println(len(m1))
	
	for key, value := range m1 {
		fmt.Println(key, " = ", value)
	}

	// delete key/value pair
	delete(m1, "A")
	fmt.Println(m1)

	// tell if the key exists
	value, ok := m1["B"]
	fmt.Println(ok, value)
}

func testTime(){

	p := fmt.Println

	now := time.Now()
	p(now)

	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	p(then.Weekday())

	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	diff := now.Sub(then)
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(then.Add(diff))
	p(then.Add(-diff))
}

func sum(nums ...int) int{
	total := 0
	for num := range nums{
		total += num
	}

	return total
}