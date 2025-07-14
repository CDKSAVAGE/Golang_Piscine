package main

import "fmt"



func first_rune(s string) rune{
    b:=[]rune(s)
    return b[0]
}
func main() {
    fmt.Println(first_rune("Hello"))
 
    fmt.Println("Hello World")

}
