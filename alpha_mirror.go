package main

import "fmt"

func alpha_mirror(s string) string{
    b:=[]rune(s)
    for i,char:=range b{
        if char>='a' && char<='z'{
            b[i]='z'-(char-'a' )
        }else if char>='A' && char<='Z' {
            b[i]='Z' - (char-'A')
        }
    }
    return string(b)
}

func main() {
    result:="My horse is Amazing" 
    fmt.Println(alpha_mirror(result))
}
