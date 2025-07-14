package main

import "fmt"

func is_alpha(c rune) bool {
    return(c>='a' && c<='z' )||(c>='A' && c<='Z')
}

func alpha_count(c rune) int{
    if!is_alpha(c){
        return 0
    }
    if c>='a' && c<='z' {
        return int(c-'a' +1)
    }
    return int(c-'a'+1)
}

func main(){
    result:='z' 
    fmt.Println(alpha_count(result))
}

