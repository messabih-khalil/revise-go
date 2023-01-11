package main

type Modifier interface{
	health(h int)
	shield(s int)
}

type Player struct{
	playerHealth int
}
type Monster struct{
	monsterHealth int
}


func (p * Player) health(h int){

}

func main(){

}