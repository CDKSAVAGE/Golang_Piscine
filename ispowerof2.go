package main

import "fmt"

func ispowerof2(n int)bool{
    return n>0&&(n&(n-1))==0
}
