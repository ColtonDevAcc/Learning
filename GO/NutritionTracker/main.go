package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	ns := GetNutritionalScore(NutritionalData{	
		Energy: EnergyFromKcal(100),
		Sugars: SugarsFromGram(10),
		SaturatedFattyAcids: SaturatedFattyAcidsFromGram(10),
		Sodium: SodiumFromMg(10),
		Fruits: FruitsPercent(10),
		Fiber:	FielderFromGram(10),
		Protein:	ProteinFromGram(10),
		}, Food: true,
	),

	fmt.Println(ns.Value)
}
