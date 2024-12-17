package main

import (
	"fmt"
	"math/rand"
)

type Match struct{
	win bool
	loss bool
}

var(
	winCount float64 = 0
	lossCount float64  = 0
	gameCount float64 = 0
)


func main(){
	game(selectDifficulty())
}

func generateRandomNumber(max int)int{
	return rand.Intn(max)
}

func selectDifficulty()int{
	var diff int
	var max int = 10
	fmt.Println("Seleziona la difficoltà selezionando il numero correlato:\n",
	"1.Facile	2.Media		3.Difficile")
	fmt.Scanln(&diff)

	switch diff{
	case 1:
		max = 10
	case 2:
		max = 100
	case 3:
		max = 1000
	default:
		fmt.Printf("Hai scelto %v. Opzione non riconosciuta.\n",diff)
		selectDifficulty()
	}
	return max
}

func game (max int){
	gameCount++
	var (
		randNum int
		choice int
		keep int
	)
	fmt.Println("Scegli un numero compreso tra 0 e", max)
	fmt.Scanln(&choice)

	randNum = generateRandomNumber(max)
	if 0 > choice || choice > max{
		fmt.Printf("Hai scelto %v. Sei andato leggermente fuori range.",choice)
		gameCount--
		game(max)
	}else if choice == randNum{
		winCount++
		
		fmt.Println("Congratulazioni, hai vinto! Ora voglio la rivincita! Continuiamo?\n",
		"1.Sì	2.No	3.Cambia difficoltà		4.Mostrami lo storico delle partite")
		fmt.Scanln(&keep)
	} else {
		lossCount++
		
		fmt.Println("Mi dispiace, hai sbagliato, avevo scelto",randNum,". Vuoi riprovare?\n1.Sì\t2.No\t3.Cambia difficoltà\t4.Mostrami lo storico delle partite")
		fmt.Scanln(&keep)
		
	}

	switch keep{
	case 1:
		game(max)
	case 2:
		break
	case 3:
		max=selectDifficulty()
		game(max)
	case 4:
		gameData()
		fmt.Println("Vuoi continuare?\n1.Sì\t2.No\t3.Cambia difficoltà")
		fmt.Scanln(&keep)
	}
}

func gameData(){
	var winRate float64 = (winCount/gameCount)*100
	fmt.Println("\nHai giocato:\n",
	gameCount, "partite in totale",
	"\nVinte: ", winCount,"\tPerse: ", lossCount,
	"\nHai una percentuale di vittorie del: ", winRate,"%\n")
}