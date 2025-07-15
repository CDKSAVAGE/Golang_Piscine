package main

import ("fmt" 
)

func str_len(s string)int{
    count:=0
    if len(s)==0{
        fmt.Println("\n")
    }
    for i:=0;i<len(s);i++{
        count++
    }
    return count 
}

func main(){
    result:="Hello World " 
    fmt.Println (str_len(result))
}
 
