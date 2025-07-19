package main

import "fmt"

func last_word(s string)string {
    word:="" 
    for i:=len(s)-1;i>=0;i--{
        if s[i]==' ' && word!="" {
            break 
        }
        if s[i]!=' '{
            word=string(s[i])+word
        }
         
    }
    return word 
}

func main(){
    result:="Hello World" 
    fmt.Println(last_word(result))
}
