package main 

import "fmt"

func SliceAdd(slice []int , num int) []int {
   return append(slice, num)
}

func main (){
    my_val:=[]int{1,2,3,4,5} 
    x:=5
    results:=SliceAdd(my_val,x)
    fmt.Println(results )
    
 
}
