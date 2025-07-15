package main

import ("fmt" 
"os")

func firstArg(s string) string {
    if len(os.Args)>1{
        return os.Args[1]
    }
    return "" 
}

func main(){
    result:="Hello World" 
    fmt.Println(firstArg(result))
}
