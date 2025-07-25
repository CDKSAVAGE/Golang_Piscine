package main

import "fmt"

func Itoa(n int)string {
    if n==0{
        return "0"
    }
    sign:="" 
    if n<0{
        sign="-"
        n=-n
    }
    result:="" 
    for n>0{
      digit:=n%10
      result=string(digit+'0')+result 
      n=n/10
    }
    return result + sign
}


func main (){
    fmt.Println(Itoa(123))
}
