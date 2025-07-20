package main

import "fmt"

func countword(s string)int{
    runes:=[]rune(s)
    n:=len(runes)
    inwords:=false 
    count:=0
    for i:=0;i<n;i++{
        if runes[i]!=' ' && !inwords{
            inwords=true 
            count++
        }else if runes[i]==' ' {
            inwords=false 
        }
    }
    return count 
}


func main(){
    result:="Hello World" 
    fmt.Println(countword(result))
}
