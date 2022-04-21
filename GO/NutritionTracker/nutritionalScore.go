package main

type ScoreType int

const (
	Food ScoreType = iota
	Beverage
	Water
	Cheese
)

type NutritionalScore struct {
	Value     int
	Positive  int
	Negative  int
	ScoreType ScoreType
}

type EnergyKJ float64
type SugarGram float64
type SaturatedFattyAcidsGram float64
type SodiumMg float64
type FiberGram float64
type FruitsPercent float64
type ProteinGram float64

type NutritionalData struct {
	Energy              EnergyKJ
	Sugars              SugarGram
	SaturatedFattyAcids SaturatedFattyAcidsGram
	Sodium              SodiumMg
	Fruits              FruitsPercent
	Fiber               FiberGram
	Protein             ProteinGram
	IsWater             bool
}

func (e EnergyKJ) GetPoints(st ScoreType) int {
	return 0
}

func (s SugarGram) GetPoints(st ScoreType) int {
	return 0
}

func (sf SaturatedFattyAcidsGram) GetPoints(st ScoreType) int {
	return 0
}
func (s SodiumMg) GetPoints(st ScoreType) int {
	return 0
}

func (fp FruitsPercent) GetPoints(st ScoreType) int {
	return 0
}

func (g FiberGram) GetPoints(st ScoreType) int {
	return 0
}

func (g ProteinGram) GetPoints(st ScoreType) int {
	return 0
}

func GetNutritionalScore(n NutritionalData, st ScoreType) NutritionalScore {
	value := 0
	positive := 0
	negative := 0

	if st != Water {
		fruitsPoints := n.Fruits.GetPoints(st)
		fiberpoints := n.Fiber.GetPoints(st)

		negative = int(n.Energy.GetPoints(st)) + int(n.Sugars.GetPoints(st)) + int(n.SaturatedFattyAcids.GetPoints(st)) + int(n.Sodium.GetPoints(st))
		positive = int(n.Protein.GetPoints(st)) + fruitPoints + fiberPoints
	}

	return NutritionalScore{
		Value:     value,
		Positive:  positive,
		Negative:  negative,
		ScoreType: st,
	}
}
