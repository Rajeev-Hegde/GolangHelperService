package src

import ("fmt"
		"strconv"
		"math/rand"
		)



func convertNumbersToString(a,b int) string {

	return ""+strconv.Itoa(a)+" "+strconv.Itoa(b);
}


func main(){
	fmt.Println("Hello!!!")
	fmt.Println(rand.Intn(100))

	var a =2
	fmt.Println(a)

	b := &a

	fmt.Println(b)
	fmt.Println(*b**b)
	fmt.Println(convertNumbersToString(a,*b))
	//fmt.Println("Sum of 2, 3: ", Add2Integers(2,3))
}

