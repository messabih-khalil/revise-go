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

func main(){
	
}