package Sleep

import "testing"
import "fmt"
import "time"


func Test_Sleep1(t* testing.T){	
	
	t1 := time.Now().Second()
	fmt.Println(t1)

	Sleep(5)
    
	t2 :=  time.Now().Second()
	fmt.Println(t2)

	if(t2-t1 < 5){
	 t.Error("The function of Sleep didn't work as expected")
	 fmt.Println("The function of Sleep didn't work as expected")
	}else{
	t.Log("1st Sleep test has passed")
	fmt.Println("1st Sleep test has passed")
	}		
}

func Test_Sleep2(t* testing.T){
	t1 := time.Now().Second()
	fmt.Println(t1)

	Sleep(25)
    
	t2 :=  time.Now().Second()
	fmt.Println(t2)

	if(t2-t1  < 25){
	 t.Error("The function of Sleep didn't work as expected")
	 fmt.Println("The function of Sleep didn't work as expected")
	}else{
	t.Log("2nd Sleep test has passed")
	fmt.Println("2nd Sleep test has passed")
	}
}

