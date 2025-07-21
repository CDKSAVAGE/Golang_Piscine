package main

import "fmt"

func trim(s string)string{
    start:=0
    end:=len(s)-1
    for start<=end && s[start]==' '{
        start++
       
    }
    for end>=start && s[end]==' ' {
        end--
    }
    return s[start:end+1]
}


func lastword(s string)string{
    word:=trim(s)
    end:=len(word)-1
    start:=end
    for start>=0 && word[start]!=' ' {
        start--
    }
    return word[start+1:]
}

func main(){
    result:=" Hello World cdk " 
    fmt.Println(lastword(result))
}
 
