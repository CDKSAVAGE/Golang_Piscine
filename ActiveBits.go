package main

import "fmt"

func ActiveBits(n int)int{
    count:=0
    for n!=0{
        if n&1==1{
            count++
        }
        n=n>>1
    }
    return count 
}


func main(){
    r:=7
    fmt.Println(ActiveBits(r))
}
 
