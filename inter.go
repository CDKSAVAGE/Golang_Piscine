package main

import "fmt" 


func interc(s1, s2 string)string {
    s1_rune:=[]rune(s1)
    s2_rune:=[]rune(s2)
    result:="" 
    for i:=0;i<len(s1_rune);i++{
        found:=false 
        for j:=0;j<len(s2_rune);j++{
            if s1_rune[i]==s2_rune[j]{
                found=true 
                break 
            }
        }
        if found{
            inresult:=false 
            for k:=0;k<len(result);k++{
                if s1_rune[i]==rune(result[k]){
                    inresult=true
                    break 
                }
            }
            if!inresult{
                result+=string(s1_rune[i])
            }
        }
    }
    return result 
}

func main (){
    fmt.Println(interc("enter your name", "my name is cdk" ))
}
