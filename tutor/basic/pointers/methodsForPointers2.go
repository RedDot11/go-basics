package main

import "fmt"

type person struct{
	name string
	age int
}

func(p *person) updateAge(newAge int){
	(*p).age=newAge
}

func main(){
	var tom=person{name:"Tom",age:23}
	fmt.Println("before:",tom.age)
	tom.updateAge(33)
	fmt.Println("after:",tom.age)
}
