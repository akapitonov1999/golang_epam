package main

import "fmt"
import "sort"

var median float64
var L int
var sum int

func main(){
Median([]int{5, 3, 2, 1, 4, 6})}

func Median (i []int) float64 {
L = 0
sum = 0
sort.Ints(i)
fmt.Println(i)
l := len(i)
if l == 0 {
fmt.Println("No median")
} else if l%2 == 0 {
Sum(i)
median = (float64(sum) + float64(i[l/2]) + float64(i[l/2-1]))/float64((l+2))
fmt.Println(median)
} else {median = float64(i[l/2])
fmt.Println(median)}
return median
}

func Sum(i []int) int {
l := len(i)
for L:=0; L<l; L++{
sum = sum + i[L]
}
fmt.Println(sum)
return sum
}
