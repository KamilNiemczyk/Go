package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func withoutChange(ilerund int) float64 {
	licznik := 0
	max := ilerund
	for i := 0; i < max; i++ {
		lista := [3]bool{}
		losowanie := rand.Intn(3)
		lista[losowanie] = true
		losowywybor := rand.Intn(3)
		if lista[losowywybor] {
			licznik++
		}
	}
	return float64(licznik) / float64(max)
}

func zezmiana(ilerund int) float64 {
	licznik := 0
	max := ilerund
	for i := 0; i < max; i++ {
		lista := [3]bool{}
		losowanie := rand.Intn(3)
		lista[losowanie] = true
		nowalista := lista[:] // convert array to slice
		losowywybor := rand.Intn(3)
		nowalista = append(nowalista[:losowywybor], nowalista[losowywybor+1:]...)
		if nowalista[0] != false || nowalista[1] != false {
			licznik++
		}
	}
	return float64(licznik) / float64(max)
}

func ileprocentwygranych(nowalista []bool, ilepudel int) float64 {
	ilosctruepudel := 0
	iloscfalsepudel := 0
	for i := 0; i < ilepudel; i++ {
		if nowalista[i] == true {
			ilosctruepudel++
		} else {
			iloscfalsepudel++
		}
	}
	return float64(ilosctruepudel) / float64(ilepudel)
}

func odrzucpudlo(nowalista []bool) []bool {
	for i := 0; i < len(nowalista)-1; i++ {
		if nowalista[i] == false {
			nowalista = append(nowalista[:i], nowalista[i+1:]...)
			break
		}
	}
	return nowalista
}

func pudla(ilepudel int) string {
	lista := make([]bool, ilepudel)
	gdziewygrana := rand.Intn(ilepudel)
	lista[gdziewygrana] = true
	wyborgracza := rand.Intn(ilepudel)
	if lista[wyborgracza] {
		return "Nie ma znaczenia ile pudeł otworzy prowadzacy, bo gracz od razu wygrał"
	}
	nowalista := lista[:] // convert array to slice
	nowalista = append(nowalista[:wyborgracza], nowalista[wyborgracza+1:]...)
	ileprowadzacymusiotworzyc := 0
	i := 0
	for i < len(nowalista) {
		if ileprocentwygranych(nowalista, len(nowalista)) > 0.5 {
			break
		} else {
			if nowalista[i] == false {
				nowalista = append(nowalista[:i], nowalista[i+1:]...)
				ileprowadzacymusiotworzyc++
			} else {
				i++
			}
		}
	}
	fmt.Println(nowalista)
	return "Prowadzacy musi otworzyć " + strconv.Itoa(ileprowadzacymusiotworzyc) + " żeby opłacało sie zmieniac wybór"
}

// func pudla(ilepudel int) string {
// 	lista := make([]bool, ilepudel)
// 	gdziewygrana := rand.Intn(ilepudel)
// 	lista[gdziewygrana] = true
// 	wyborgracza := rand.Intn(ilepudel)
// 	nowalista := lista[:] // convert array to slice
// 	nowalista = append(nowalista[:wyborgracza], nowalista[wyborgracza+1:]...)
// 	ileodrzuconych := 0
// 	for true {
// 		nowalista = odrzucpudlo(nowalista)
// 		ileodrzuconych++
// 		procentdlagracza := ileprocentwygranych(nowalista, len(nowalista))
// 		procentdlaotwierajacego := 1 - procentdlagracza
// 		fmt.Println("Lista: ", nowalista)
// 		fmt.Println("Procent dla gracza: ", procentdlagracza)
// 		fmt.Println("Procent dla otwierajacego: ", procentdlaotwierajacego)
// 		if ileprocentwygranych(nowalista, len(nowalista)) >= 0.5 {
// 			break
// 		}
// 	}
// 	if lista[wyborgracza] {
// 		return "Wygrałeś"
// 	}
// 	return "Przegrałeś"
// }

//szansa na wygraną przy zmianie wyboru pudla

func main() {
	argumenty := os.Args[1:]
	arg1, _ := strconv.Atoi(argumenty[0])
	arg2 := argumenty[1]
	if arg2 == "change" {
		fmt.Println(zezmiana(arg1))
	} else if arg2 == "stay" {
		fmt.Println(withoutChange(arg1))
	} else {
		fmt.Println(pudla(5))
	}
}
