package main

import "fmt"

type Name string

type Product struct{
	name string
	price int
	inStock bool
}

func main(){
	shopingList := make(map[Name]Product)

	shopingList["product 1"] = Product{"xbox" , 5000 , true}
	shopingList["product 2"] = Product{"ps5" , 5000 , false}
	shopingList["product 3"] = Product{"iphone 14" , 5000 , true}

	for k , v := range shopingList{
		fmt.Println(k , "\nname : " , v.name ,"\nprice : " , v.price , "\navailable : " , v.inStock)
	}
}