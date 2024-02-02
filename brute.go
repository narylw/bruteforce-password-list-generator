package main

import	(
	
	"fmt"
	"os"
	
	"strings"

	
	"strconv"

)
type List []interface{}
func GenCombinations(n, r int) <-chan []int {

	if r > n {
		panic("Invalid arguments")
	}

	ch := make(chan []int)

	go func() {
		result := make([]int, r)
		for i := range result {
			result[i] = i
		}

		temp := make([]int, r)
		copy(temp, result) 
		ch <- temp

		for {
			for i := r - 1; i >= 0; i-- {
				if result[i] < i+n-r {
					result[i]++
					for j := 1; j < r-i; j++ {
						result[i+j] = result[i] + j
					}
					temp := make([]int, r)
					copy(temp, result) 
					ch <- temp
					break
				}
			}
			if result[0] >= n-r {
				break
			}
		}
		close(ch)

	}()
	return ch
}



func CombinationsInt(iterable []int, r int) chan []int {

	ch := make(chan []int)

	go func() {

		length := len(iterable)

		for comb := range GenCombinations(length, r) {
			result := make([]int, r)
			for i, val := range comb {
				result[i] = iterable[val]
			}
			ch <- result
		}

		close(ch)
	}()
	return ch
}



func CombinationsStr(iterable []string, r int) chan []string {

	ch := make(chan []string)

	go func() {

		length := len(iterable)

		for comb := range GenCombinations(length, r) {
			result := make([]string, r)
			for i, val := range comb {
				result[i] = iterable[val]
			}
			ch <- result
		}

		close(ch)
	}()
	return ch
}





func CombinationsList(iterable List, r int) chan List {

	ch := make(chan List)

	go func() {

		length := len(iterable)

		for comb := range GenCombinations(length, r) {
			result := make(List, r)
			for i, val := range comb {
				result[i] = iterable[val]
			}
			ch <- result
		}

		close(ch)
	}()
	return ch
}
func GenPermutations(n int) <-chan []int {
	if n < 0 {
		panic("Invalid argument")
	}

	ch := make(chan []int)

	go func() {
		var finished bool

		result := make([]int, n)

		for i := range result {
			result[i] = i
		}

		temp := make([]int, n)
		copy(temp, result) 
		ch <- temp

		for {
			finished = true

			for i := n - 1; i > 0; i-- {

				if result[i] > result[i-1] {
					finished = false

					minGreaterIndex := i
					for j := i + 1; j < n; j++ {
						if result[j] < result[minGreaterIndex] && result[j] > result[i-1] {
							minGreaterIndex = j
						}

					}

					result[i-1], result[minGreaterIndex] = result[minGreaterIndex], result[i-1]

					
					for j := i; j < n; j++ {
						for k := j + 1; k < n; k++ {
							if result[j] > result[k] {
								result[j], result[k] = result[k], result[j]
							}

						}
					}
					break
				}
			}

			if finished {
				close(ch)
				break
			}
			temp := make([]int, n)
			copy(temp, result) 
			ch <- temp

		}

	}()
	return ch
}



func PermutationsInt(iterable []int, r int) chan []int {

	ch := make(chan []int)

	go func() {

		length := len(iterable)

		for comb := range GenCombinations(length, r) {
			for perm := range GenPermutations(r) {
				result := make([]int, r)
				for i := 0; i < r; i++ {
					result[i] = iterable[comb[perm[i]]]
				}
				ch <- result
			}
		}

		close(ch)
	}()
	return ch
}



func PermutationsStr(iterable []string, r int) chan []string {

	ch := make(chan []string)

	go func() {

		length := len(iterable)

		for comb := range GenCombinations(length, r) {
			for perm := range GenPermutations(r) {
				result := make([]string, r)
				for i := 0; i < r; i++ {
					result[i] = iterable[comb[perm[i]]]
				}
				ch <- result
			}
		}

		close(ch)
	}()
	return ch
}





func PermutationsList(iterable List, r int) chan List {

	ch := make(chan List)

	go func() {

		length := len(iterable)

		for comb := range GenCombinations(length, r) {
			for perm := range GenPermutations(r) {
				result := make(List, r)
				for i := 0; i < r; i++ {
					result[i] = iterable[comb[perm[i]]]
				}
				ch <- result
			}
		}

		close(ch)
	}()
	return ch
}

func main()  {


	
		var passwordLength string;
			fmt.Println(`
 ____   _  _,             _,  _,  ____, ____,  _  _,  __,    _   _,
(-|__) (-\_/     ____,   (-|\ |  (-/_| (-|__) (-\_/  (-|    (-|  | 
 _|__)   _|,    (         _| \|, _/  |, _|  \,  _|,   _|__,  _|/\|,
(       (                (      (      (       (     (      (      
[ Instagram ] = @ Narylw
	  `)
	fmt.Println("Please enter the PasswordLength you want:")
	fmt.Scan(&passwordLength)

	
	characters := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()+-./"


	
	passwordLengthList := strings.Split(passwordLength, ",")
	
	
	charactersList := strings.Split(characters, "")


	
	for _, passLen := range passwordLengthList {
	

		
		passLenInt, err := strconv.Atoi(passLen)
		
		
		if err != nil {
			panic(err)
		}

		var str string;
		
		for v := range PermutationsStr(charactersList, passLenInt) {
			
			
			
			str += strings.Join(v, "") + "\n";
		}
		  file, errs := os.Create("PassworsList.txt")
   if errs != nil {
      fmt.Println("Failed to create file:", errs)
      return
   }
   defer file.Close()
   
   _, errs = file.WriteString(str)
   if errs != nil {
      fmt.Println("Failed to write to file:", errs) 
      return
   }
   fmt.Println("Wrote to file 'PassworsList.txt'.") 
		
	}
}