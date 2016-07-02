package main
import (
  "fmt"
  "net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Hello!")
}

func InsertionSort(arr []int) []int {
  var size = len(arr)

  for i:=1; i < size; i++ {
    if arr[i] < arr[i-1] {
      var tmp = arr[i]
      arr[i]  = arr[i-1]

      for j := i-1; j >= 0; j-- {
        if j == 0 {
          arr[j] = tmp
          break
        }
        if tmp >= arr[j-1] {
          arr[j] = tmp
          break
        } else {
          arr[j] = arr[j-1]
        }
      }
    }
  }
  return arr
}

func BubbleSort(arr []int) []int {
  var size = len(arr)
  var swapped = true

  for ; swapped; {
    swapped = false
    for i:=0; i < size - 1; i++ {
      if arr[i] > arr[i+1] {
        var tmp = arr[i]
        arr[i] = arr[i+1]
        arr[i+1] = tmp
        swapped = true
        // fmt.Printf("Swapping %d and %d\n", arr[i], arr[i+1])
      }
    }
  }

  return arr
}

func main() {
  var a = 43
  // var h Hello
  var mixed  = []int{ 3, 5, 2, 65, 14, 7, 33, 12}
  var mixed2 = []int{ 3, 5, 2, 65, 14, 7, 33, 12}

  fmt.Printf("Sorting algorithms %d\n", a)
  fmt.Printf("Soring with BubbleSort %d\n", mixed)
  fmt.Printf("-> %d\n", BubbleSort(mixed))
  fmt.Printf("Soring with InsertionSort %d\n", mixed2)
  fmt.Printf("-> %d\n", InsertionSort(mixed2))

  // http.ListenAndServe("localhost:4000", h)
}

func print_range(start int, stop int) {
  if (start >= stop) {
    fmt.Print("Start must be strictly greater than stop\n")
    return
  }
  fmt.Printf("printing range from %d to %d.\n", start, stop)
  var all_the_numbers []int
  for i := start; i <= stop; i++ {
    all_the_numbers = append(all_the_numbers, i)
    fmt.Printf("Value %d\n", i)
    fmt.Printf("All: %d\n\n", all_the_numbers)
  }
  return
}
