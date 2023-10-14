package main

import "fmt"

const ebulicaoF = 212.0

func main() {
	
	tempF := ebulicaoF
	tempC := (tempF - 32) * 5/9

	fmt.Println("Temperatura de ebulição da àgua em ºF = ", tempF)
	fmt.Println("Temperatura de ebulição da àgua em ºC = ", tempC)

}

