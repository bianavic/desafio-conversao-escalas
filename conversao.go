package main

import "fmt"

const ebulicaoK = 373

func main() {
	tempK := ebulicaoK
	tempC := (tempK - 273)
	fmt.Printf("A conversão da temperatura de ebulição da água de %vK para celsius resulta em %v graus C", tempK, tempC)
}
