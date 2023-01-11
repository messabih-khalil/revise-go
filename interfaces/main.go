package main

import "fmt"

type Modifier interface{
	health(h int)
}

type Player struct{
	playerHealth int
}
type Monster struct{
	monsterHealth int
}


func (p * Player) health(h int){
	p.playerHealth = p.playerHealth - h
}

func (m * Monster) health(h int){
	m.monsterHealth = m.monsterHealth - h
}


func execute(k Modifier){
	k.health(15)
}

func main(){
	man := Player{20}
	dog := Monster{30}

	execute(&man)
	execute(&dog)

	fmt.Println(dog.monsterHealth)
	fmt.Println(man.playerHealth)
}