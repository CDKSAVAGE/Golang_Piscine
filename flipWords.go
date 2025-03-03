// THIS PROGRAM FLIPWORDS OF A STRING, THE LAST STRING BECOMES FIRST AND SO ON

package main
import ("fmt"
        "strings")
        
func flipword(str string)string{
    word:=strings.Fields(str)
    if len(word)==0{
        return "invalid output "
    }
    for i, j:=0,len(word)-1;i<j;i,j=i+1,j-1{
        word[i],word[j]=word[j],word[i]
    }
    return strings.Join(word, " ")
}
func main() {
  fmt.Println(flipword("CDK is my name and I love coding ") )
}
