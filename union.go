package main

import "fmt"

func union(s1, s2 string)string {
    seen:=make(map[rune]bool)
    result:="" 
    for _, char:=range s1+s2{
        if!seen{
            seen[char]=true
            result+=string(char)
        }
    }
    return result 
}
