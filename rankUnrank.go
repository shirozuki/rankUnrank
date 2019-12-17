package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"sort"
	"time"
)

func fact(n int) int {
	var f int

	f = 1

	for n >= 1 {
		f *= n
		n--
	}

	return f
}

func cpyArr(arr []int) []int {

	tmpArr := make([]int, len(arr))
	copy(tmpArr, arr)

	return tmpArr
}

func sortArr(arr []int) []int {

	stdArr := cpyArr(arr)

	sort.Ints(stdArr)
	return stdArr
}

func genArr(n int) []int {

	var arr []int
	var exs bool

	for i := 0; i < n; {
		exs = false
		tmp := rand.Intn(n) + 1

		for j := 0; j < i; j++ {
			if tmp == arr[j] {
				exs = true
				break
			}
		}

		if exs == true {
			continue
		} else {
			arr = append(arr, tmp)
			i++
		}
	}

	return arr
}

func findSmaller(arr []int, low int, high int) int {
	var cr int
	cr = 0

	for i := low + 1; i <= high; i++ {
		if arr[i] < arr[low] {
			cr++
		}

	}

	return cr

}

func findRank(arr []int) int {

	var ln, mul, r, cr int

	ln = len(arr)
	mul = fact(ln)

	r = 1

	for i := 0; i < ln; i++ {
		mul /= ln - i
		cr = findSmaller(arr, i, ln-1)

		r += cr * mul
	}

	return r
}

func unrank(arr []int, n int, r int) []int {

	var fct, q int

	fct = 1
	for i := 2; i < n; i++ {
		fct *= i
	}

	r--

	var unrPrm []int

	for i := 0; i < n; i++ {
		q = int(math.Floor(float64((r / fct))))
		r = r % fct
		unrPrm = append(unrPrm, arr[q])

		arr = arr[:q+copy(arr[q:], arr[q+1:])]

		if i != n-1 {
			fct = int(math.Floor(float64(fct / (n - 1 - i))))
		}
	}

	return unrPrm
}

func listAll(arr []int, max int, ln int, rank int) {

	if rank > max {
		os.Exit(0)
	}

	newArr := cpyArr(arr)

	fmt.Println(rank, ":", unrank(newArr, ln, rank))
	listAll(arr, max, ln, rank+1)
}

func main() {

	rand.Seed(time.Now().UnixNano())

	var ln int

	fmt.Print("Długość permutacji: ")
	_, err := fmt.Scanf("%d", &ln)

	if err != nil {
		fmt.Println("Err")
	}

	var inpArr, stdArr, befArr, aftArr, rankArr, recArr []int

	inpArr = genArr(ln)
	stdArr = sortArr(inpArr)
	rankArr = cpyArr(inpArr)
	befArr = cpyArr(stdArr)
	aftArr = cpyArr(stdArr)
	recArr = cpyArr(stdArr)

	fmt.Println("Permutacja:", inpArr)

	rank := findRank(rankArr)

	fmt.Println("Rank:", rank)

	fmt.Println("Len:", ln)

	fct := fact(ln)
	fmt.Println("Max:", fct)

	if rank <= 0 || rank > fct {
		fmt.Println("Perm (", rank, ") nie istnieje.")
	} else {
		fmt.Println("Perm (", rank, "): ", unrank(stdArr, ln, rank))
		if rank-1 <= 0 {
			fmt.Println("Poprzednik nie istnieje.")
		} else {
			fmt.Println(" Poprzednik: ", unrank(befArr, ln, rank-1))
		}
		if rank+1 > fct {
			fmt.Println("Następnik nie istnieje.")
		} else {
			fmt.Println("  Następnik: ", unrank(aftArr, ln, rank+1))
		}
	}

	time.Sleep(10 * time.Second)
	listAll(recArr, fct, ln, 1)

	/*
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("\nWypisać wszystkie permutacje? (T/N) >> ")
		listPerms, _ := reader.ReadString('\n')

		if runtime.GOOS == "windows" {
			time.Sleep(2 * time.Second)
			//listPerms = strings.Replace(listPerms, "\r\n", "", -1)
		} else {
			listPerms = strings.Replace(listPerms, "\n", "", -1)
		}

		if strings.Compare("T", listPerms) == 0 || strings.Compare("t", listPerms) == 0 {
			listAll(recArr, fct, ln, 1)
		} else {
			os.Exit(0)
		}
	*/

}
