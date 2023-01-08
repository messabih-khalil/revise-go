package main

import "fmt"


type Student struct{
	name string
	highSchool bool
}

type ClassRoom struct{
	students [3]Student
}


func main(){
	aldin := Student{"aldin" , true}

	var (
		tars = Student{"tars" , false}
		james = Student{"james" , true}
	)

	class := ClassRoom{}

	class.students[0] = aldin
	class.students[1] = tars
	class.students[2] = james

	// show details 

	for _ , v := range class.students {
		if v.highSchool {
			fmt.Println(v.name , "Is on high school")
		}
	}
}