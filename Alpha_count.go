package main

import "fmt"

func alpha_count(s string) int{
    count:=0
    for _, char:=range s{
       if char>='a' && char<='z' || char>='A' && char <='Z'{
           count++
       }
    }
    return count
}

func main(){
    result:=" Hello " 
    fmt.Println(alpha_count(result))
}
