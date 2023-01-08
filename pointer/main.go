package main

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

// func checkOut(ps * []Product)  {
// 	for _ , v := range Product{
// 		v.
// 	}
// }

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
}