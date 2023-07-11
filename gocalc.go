package main
import("fmt")
func main() {
var number1 float64
fmt.Scanln(&number1)
var operator string
fmt.Scanln(&operator)
var number2 float64
fmt.Scanln(&number2)
switch operator {
  case "+":
    fmt.Println(number1 + number2)
  case "*" , "x" , ".":
  fmt.Println(number1 * number2)
  case "/", ":":
  fmt.Println(number1/number2)
case "-":
  fmt.Println(number1 - number2)
default:
  fmt.Println("No valid operator provided.")
}
}
