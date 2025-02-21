package main 


import "fmt"

func SliceRemove(slice []int , num int) []int {
for index, element :=range slice{
    if element ==num{
        return append (slice[:index], slice [index+1:]...)
    }
}
return slice 
}


func main (){
    my_slice:=[]int{1,2,3,4,5}
    fmt.Println("before ", my_slice)
    val:=5
    results:=(SliceRemove(my_slice, val))
    fmt.Println(results )
}
