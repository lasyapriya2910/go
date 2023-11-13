package main
import "fmt"
// start the function main ()
func main() {
fmt.Println("Golang program to check for palindrome")
// declare the variables
var number,rem,temporary int
var reverse int = 0
// initialize the number variable
fmt.Println("enter a string:")
fmt.Scan(&number)
temporary=number
// For Loop used
for{
rem = number%10
reverse = reverse*10 + rem
number /= 10
if(number==0){
break // Break Statement used to exit from loop
}
}
if(temporary==reverse){
fmt.Printf("Number %d is a Palindrome",temporary)
}else{
fmt.Printf("Number %d is not a Palindrome",temporary)
}
// print the result using fmt.Printf () function
}