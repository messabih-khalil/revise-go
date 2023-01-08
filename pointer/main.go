package main

import "fmt"

type Item struct{
	name string
	price float64
}

type Product struct{
	item Item
	security bool
}

func (p * Product) statusChanger(){
	p.security = !p.security
}

func checkOut(ps []Product)  {
	for i := 0; i < len(ps); i++ {
		ps[i].statusChanger()
	}
}

func printData(ps []Product){
	for i := 0; i < len(ps); i++ {
		fmt.Println(ps[i])
	}
}

const (
	inactive = false
	active = true
)

func main(){
	items := make([]Product , 4)

	items[0] = Product{
		item: Item{
			name : "xbox",
			price: 2000.0,
		},
		security: active,
	}
	items[1] = Product{
		item: Item{
			name : "ps5",
			price: 2000.0,
		},
		security: active,
	}
	items[0] = Product{
		item: Item{
			name : "iphone",
			price: 2000.0,
		},
		security: active,
	}
	items[0] = Product{
		item: Item{
			name : "mac",
			price: 2000.0,
		},
		security: active,
	}

	checkOut(items)
}