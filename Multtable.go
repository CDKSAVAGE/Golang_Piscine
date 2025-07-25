package main

import "fmt"

func Atoi(s string )int {
    result:=0
    sign:=1
    start:=0
    if len(s)==0{
        return 0
    }
    if s[0]=='-'{
        sign=-1
        start=1
    }else if s[0]=='+'{
        start=1
    }
    for i:=start;i<len(s);i++{
        if s[i]<'0'||s[i]>'9'{
            break 
        }
        digit:=int(s[i]-'0')
        result=result*10+digit 
    }
    return result*sign
}


func main(){
    fmt.Println(Atoi("+123b"))
    
    input:=Atoi("9")
    for i:=1;i<9;i++{
        fmt.Printf("%d * %d =%d\n", i, input, i*input)
    }
}
