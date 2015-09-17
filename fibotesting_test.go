package fibonacci

import "testing"
import "fmt"

func Test_Fibonacci1(t* testing.T){
	if(fib(5) != 5){
	 t.Error("The function of Fibonacci didn't work as expected")
	 fmt.Println("The function of Fibonacci didn't work as expected")
	}else{
	t.Log("1st test has passed")
	fmt.Println("1st test has passed")
	}
}


func Test_Fibonacci2(t* testing.T){
	if(fib(2) != 1){
	 t.Error("The function of Fibonacci didn't work as expected")
	 fmt.Println("The function of Fibonacci didn't work as expected")
	}else{
	t.Log("2nd test has passed")
	fmt.Println("2nd test has passed")
	}
}