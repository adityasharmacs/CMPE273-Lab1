package perimeter

import "testing"
import "fmt"


func Test_PerimeterRectangle(t* testing.T){
	r := rectangle{width: 5,height: 6}
	if(r.perimeter() != 22){
	 t.Error("The function of perimeter for rectangle didn't work as expected")
	 fmt.Println("The function of perimeter for rectangle didn't work as expected")
	}else{
	t.Log("1st test has passed")
	fmt.Println("1st test has passed for perimeter of rectangle")
	fmt.Println(r.perimeter())
	}
}

func Test_PerimeterCircle(t* testing.T){
	c := circle{radius: 6}
	if(c.perimeter() < 37){
	 t.Error("The function of perimeter for circle didn't work as expected")
	 fmt.Println("The function of perimeter for circle didn't work as expected")
	}else{
	t.Log("2nd test has passed")
	fmt.Println("2nd test has passed for circle perimeter calculation")
	fmt.Println(c.perimeter())
	}
}