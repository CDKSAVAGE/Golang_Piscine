package main 

import "fmt"

func main (){
    my_string:="find a"
    find_A:=false
    for _, char :=range my_string{
        if(char=='a'){
           find_A=true
           if(find_A){
               fmt.Printf("%c", char)
               fmt.Println()
           }else{
               fmt.Println()
           }
        }
    }
}
