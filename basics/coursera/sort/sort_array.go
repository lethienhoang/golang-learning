package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func sub_arr_sort(wg *sync.WaitGroup, arr []int) {
	fmt.Println(arr)
	sort.Ints(arr)
	wg.Done()
}

func sub_merge(wg *sync.WaitGroup, left, right []int) []int {
	slice := make([]int, len(left)+len(right))
	i, j := 0, 0
	count := 0

	defer func() {
		wg.Done()
	}()

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			slice[count] = left[i]
			count, i = count+1, i+1
		} else {
			slice[count] = right[j]
			count, j = count+1, j+1
		}
	}

	for i < len(left) {
		slice[count] = left[i]
		count, i = count+1, i+1
	}

	for j < len(right) {
		slice[count] = right[j]
		count, j = count+1, j+1
	}

	return slice
}

func merge(arr1, arr2, arr3, arr4 []int) []int {
	var wg sync.WaitGroup
	wg.Add(3)

	slice1 := sub_merge(&wg, arr1, arr2)
	slice2 := sub_merge(&wg, arr3, arr4)

	result := sub_merge(&wg, slice1, slice2)

	wg.Wait()
	return result
}

func main() {
	fmt.Println("Please input some integers to sort: ")
	br := bufio.NewReader(os.Stdin)
	a, _, _ := br.ReadLine()
	ns := strings.Split(string(a), " ")

	var num []int
	for _, value := range ns {
		n, _ := strconv.Atoi(value)
		num = append(num, n)
	}

	// ¼ of the array - partition the array into 4 parts
	size := len(num) / 4

	var wg sync.WaitGroup

	for i := 0; i < 4; i++ {
		wg.Add(1)
		if i == 3 {
			go sub_arr_sort(&wg, num[i*size:])
		} else {
			go sub_arr_sort(&wg, num[i*size:(i+1)*size])
		}
	}

	subArr1 := num[:size]
	subArr2 := num[size : 2*size]
	subArr3 := num[2*size : 3*size]
	subArr4 := num[3*size:]

	// go sub_arr_sort(&wg, subArr1)
	// go sub_arr_sort(&wg, subArr2)
	// go sub_arr_sort(&wg, subArr3)
	// go sub_arr_sort(&wg, subArr4)

	wg.Wait()

	sorted := merge(subArr1, subArr2, subArr3, subArr4)
	fmt.Println(sorted)
}
