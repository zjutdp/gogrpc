package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"os"
	"sort"
	"time"
)

var startTime time.Time

func init() {
	startTime = time.Now()
}

// 将排好序的内存数据打印输出，或者存文件
func main() {
	//generateFile("large.in", 100000000)
	largeMergeDemo()
}

func inMemDemo() {
	p := InMemSort(ArraySource(3, 2, 4, 5, 6, 8, 90, 34, 6, 7, 4))
	for v := range p {
		fmt.Println(v)
	}
}

func inMemMergeDemo() {
	p := Merge(
		InMemSort(ArraySource(3, 2, 4, 6, 7, 4)),
		InMemSort(ArraySource(8, 9, 10, 45, 4, 6, 0, 1)))

	for v := range p {
		fmt.Println(v)
	}
}

func smallMergeDemo() {
	p := CreatePipeline("small.in", 512, 4)

	writeToFile(p, "small.out")
	printFile("small.out")
}

func largeMergeDemo() {
	p := CreatePipeline("large.in", 800000000, 4)

	writeToFile(p, "large.out")
	printFile("large.out")
}

func generateFile(filename string, n int) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	p := RandomSource(n)
	WriterSink(file, p)
}

func writeToFile(p <-chan int, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	WriterSink(writer, p)
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	p := ReaderSource(file, -1)

	limit := 100
	for v := range p {
		fmt.Println(v)
		limit--
		if limit <= 0 {
			break
		}
	}
}

// ArraySource pipeline
func ArraySource(a ...int) chan int {
	out := make(chan int)
	go func() {
		for _, v := range a {
			out <- v
		}
		close(out)
	}()

	return out
}

// 将文件读取的数据输送到一个节点
// 该节点通过goroutine将数据输送到chan
func ReaderSource(reader io.Reader, chunkSize int) <-chan int {
	out := make(chan int, 1024)

	reader = bufio.NewReader(reader)
	go func() {
		buffer := make([]byte, 8)
		bytesRead := 0
		for {
			n, err := reader.Read(buffer)
			bytesRead += n
			if n > 0 {
				v := int(binary.BigEndian.Uint64(buffer))
				out <- v
			}
			if err != nil ||
				(chunkSize != -1 && bytesRead >= chunkSize) {
				break
			}
		}

		// 一定要close，close后，外面会用if或range取判断取失败
		// 数据量大的话，不关闭会很占内存
		close(out)
	}()

	return out
}

// 将上面ReaderSource返回的chan传进来读入内存
// 使用内部排序对读入内存的数据排序
// 然后通过goroutine输出到chan返回出去
// 参数in 只进不出，返回参数只出不进
func InMemSort(in <-chan int) <-chan int {
	out := make(chan int, 1024)

	go func() {
		// Read into memory
		a := []int{}
		for v := range in {
			a = append(a, v)
		}

		fmt.Println("Read Done: ", time.Now().Sub(startTime))
		// Sort
		sort.Ints(a)

		fmt.Println("InMem Sort Done: ", time.Now().Sub(startTime))
		// Output
		for _, v := range a {
			out <- v
		}

		// close
		close(out)
	}()

	return out
}

// 将排好序的多个节点通过2路归并排序
func MergeN(inputs ...<-chan int) <-chan int {
	if len(inputs) == 1 {
		return inputs[0]
	}

	m := len(inputs) / 2

	// merge inputs[0...m) and inputs [m...end)
	return Merge(MergeN(inputs[:m]...), MergeN(inputs[m:]...))
}

// 将排好序的2个节点归并归并
func Merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int, 1024)

	go func() {
		v1, ok1 := <-in1 // 没有元素ok1返回false
		v2, ok2 := <-in2
		for ok1 || ok2 {
			// v2没有元素就出v1; v1,v2都有数据，且v1 <= v2也出v1
			if !ok2 || (ok1 && v1 <= v2) {
				out <- v1
				v1, ok1 = <-in1
			} else {
				out <- v2
				v2, ok2 = <-in2
			}
		}

		// 关闭
		close(out)
		fmt.Println("Merge done: ", time.Now().Sub(startTime))
	}()

	return out
}

func WriterSink(writer io.Writer, in <-chan int) {
	bufWriter := bufio.NewWriter(writer)
	defer bufWriter.Flush()
	for v := range in {
		buffer := make([]byte, 8)
		binary.BigEndian.PutUint64(
			buffer, uint64(v))
		bufWriter.Write(buffer)
	}
}

func RandomSource(count int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Int()
		}

		close(out)
	}()

	return out
}

func CreatePipeline(
	filename string,
	fileSize, chunkCount int) <-chan int {

	chunkSize := fileSize / chunkCount

	sortResults := []<-chan int{}
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}

		file.Seek(int64(i*chunkSize), 0)

		source := ReaderSource(file, chunkSize)

		sortResults = append(sortResults, InMemSort(source))
	}
	return MergeN(sortResults...)
}
