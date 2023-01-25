package main

import (
	"errors"
	"fmt"
	"strings"
)

type VowelFinder interface {
	FindVowelUsingInterface() string
}

type Mystring string

//var Mysstrings string

func UpperCase(in chan string, out chan string) {
	flag := false
	input := <-in
	output := ""
	for i := 0; i < len(input); i++ {
		if input[i] == ' ' {
			output += string(input[i])
			continue
		}
		if flag {
			output += string(strings.ToUpper(string(input[i])))
			flag = false
		} else {
			output += string(input[i])
			flag = true
		}
	}
	out <- output
}

func howAreYou() {
	chOutput := make(chan string)
	ch := make(chan string)
	go UpperCase(ch, chOutput)
	ch <- "how are you"
	output := <-chOutput
	fmt.Println(output)
}

func (vowel Mystring) FindVowelUsingInterface() string {
	var output string
	for i := 0; i < len(vowel); i++ {
		switch vowel[i] {
		case 'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u':
			output += string(vowel[i])
		}
	}
	return output
}

func bubbleSort(array []int) {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}

}

func binarySearch(arr []int, n, left, right int) bool {
	if left > right {
		return false
	}
	mid := (left + right) / 2
	fmt.Println(" Mid value is ", arr[mid])
	if arr[mid] == n {
		return true
	} else if n < arr[mid] {
		return binarySearch(arr, n, left, mid-1)
	} else {
		return binarySearch(arr, n, mid+1, right)
	}
}

func factorial(n int) int {
	if n == 1 {
		return n
	}
	if n == 2 {
		return n
	}
	return (n * factorial(n-1))
}

func main() {
	fmt.Println("start the programm...")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println("key is ", os.Getenv("HELLO_KEY"))

	howAreYou()
	var temp VowelFinder = Mystring("My name is Sams")
	vowel := temp.FindVowelUsingInterface()
	fmt.Println("vowels is ", vowel)
	array := []int{5, 4, 3, 6, 7, 1, 9, 2}
	fmt.Println(" Before Sort ", array)
	bubbleSort(array)
	fmt.Println(" After Sort ", array)
	left := 0
	right := len(array) - 1
	if binarySearch(array, 3, left, right) {
		fmt.Println("Number are present ")
	} else {
		fmt.Println("Number are not present ")
	}
	fmt.Println("factorial number is ", factorial(10))
}

func main() {
	fmt.Println("hello")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	// go hell(cancel)
	fmt.Println(hello(ctx))
}

func hell(c context.CancelFunc) {
	time.Sleep(5 * time.Second)
	c()
}

func hello(ctx context.Context) string {

	select {
	case <-ctx.Done():
		return "context cancelled"
	}

}

func worker(jobs <-chan int, wg *sync.WaitGroup) {
	for j := range jobs {
		time.Sleep(time.Second)
		fmt.Println("worker", j)
	}
	wg.Done()

}

func main() {
	const numJobs = 5
	jobs := make(chan int)
	var wg sync.WaitGroup
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(jobs, &wg)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)
	wg.Wait()
	time.Sleep(2 * time.Second)
}

func Hello(s string) (string, error) {
	if s == "" {
		return "", errors.New("name is empty")
	}
	return s, nil
}
