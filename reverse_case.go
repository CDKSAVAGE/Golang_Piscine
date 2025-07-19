package main

import "fmt"

func reverse_case(s string)string{
    b:=[]rune(s)
    for i,char:=range b{
        if char>='a' && char<='z' {
            b[i]=(char-32)
        }else if char>='A' && char<='Z' {
            b[i]=(char+32)
        }
    }
    return string(b)
}

func main(){
    result:="Hello World " 
    fmt.Println(reverse_case(result))
}
 
