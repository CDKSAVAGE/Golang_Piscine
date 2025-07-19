package main

import "fmt"

func first_word(s string)string {
    word:="" 
    for i:=0;i<len(s);i++{
        if s[i]==' ' {
            break 
        }
        word+=string(s[i])
    }
    return word
}


func main(){
    result:="Hello World" 
    fmt.Println(first_word(result))
}
