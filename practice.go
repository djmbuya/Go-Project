// package main

// import (
// 	"fmt"
// 	"os"
// )

//"fmt"

// //firstparam
//
//	func main() {
//		args := os.Args[1]
//		for _, v := range args {
//			z01.PrintRune(v)
//		}
//		z01.PrintRune('\n')
//	}
//
// //lastparam
//
//	func main() {
//		args := os.Args[len(os.Args)-1]
//		for _, char := range args{
//			z01.PrintRune(char)
//		}
//		z01.PrintRune('\n')
//	}
// func main() {
// 	for i := 'a'; i <= 'z'; i++ {
// 		if i%2 == 0 {
// 			z01.PrintRune(i-('a' - 'A'))
// 		} else {
// 			z01.PrintRune(i)
// 		}
// 	}
// 	z01.PrintRune('\n')
// }
// param count
// func main() {
// 	args := os.Args[1:]
// 	count := 0
// 	for range args {
// 		count++
// 	}

//		remainder:=0
//		var str string
//		   for count>0{
//			remainder=count%10
//			str=string(rune(remainder)+'0')+str
//			count=count/10
//		}
//		for _,ch:=range str{
//			z01.PrintRune(ch)
//		}
//		z01.PrintRune('\n')
//	}
//
//	func main() {
//		args := os.Args[1:]
//		count := 0
//		for i := '0'; i < rune(len(args)); i++ {
//			count++
//			z01.PrintRune(rune(count) + '0')
//		}
//	}
//
// max
//
//	func Max(a []int) int {
//		max := 0
//		for _,num := range a{
//			if num > max {
//				max = num
//			}
//		}
//		return max
//	}
//
//	func main() {
//		a := []int{23, 123, 1, 11, 55, 93}
//		max := Max(a)
//		fmt.Println(max)
//	}
//
// strlen
// func StrLen(s string) int {
// 	return len([]rune(s))
// }

//	func main() {
//		l := StrLen("Hello World!")
//		fmt.Println(l)
//	}
// func FirstRune(s string) rune {
// 	myRune := []rune(s)
// 	index:=len(myRune)-1
// 	return myRune[index]
// }

// func main() {
// 	z01.PrintRune(FirstRune("Hello!"))
// 	z01.PrintRune(FirstRune("Salut!"))
// 	z01.PrintRune(FirstRune("Ola!"))
// 	z01.PrintRune('\n')
// }

//	for i := 0; i < len(myRune); i++ {
//		if i == 0 {
//			return myRune[i]
//		}
//	}
//
// return 0
// func wdmatch(arg1, arg2 string) string {
// 	s := ""
// 	arg2rune := []rune(arg2)
// 	for _, v := range arg1 {
// 		for i, v1 := range arg2rune {
// 			if v == v1 {
// 				s += string(v)
// 				arg2rune = arg2rune[i+1:]
// 				break
// 			}
// 		}
// 	}
// 	if s != arg1 {
// 		return ""
// 	}
// 	return s
// }

// func main() {
// 	fmt.Println(wdmatch(os.Args[1], os.Args[2]))
// }