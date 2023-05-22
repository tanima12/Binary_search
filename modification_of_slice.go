package main 
import "fmt"
 
func main(){
	arr := [5]int{1,2,3,4,5}
	slice:= arr[ :3]
	fmt.Println(arr)
	fmt.Println(slice)
	slice[2]=200
	fmt.Println(arr)
	fmt.Println(slice)
}
    	
