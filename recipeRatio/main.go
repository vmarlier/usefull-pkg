package main

import (
	"fmt"
	"log"
)

// createMapMacros will create a map which contain protain (p), carbohydrate (g), sugar (s), lipids (l), saturated fatty acid (as), fibers (f) and kcal for 100g
func createMapMacros(p float32, g float32, s float32, l float32, as float32, f float32, kcal float32) map[string]float32 {
	return map[string]float32{
		"p":    p,
		"g":    g,
		"s":    s,
		"l":    l,
		"as":   as,
		"f":    f,
		"kcal": kcal,
	}
}

// muesliMacros call the createMapMacros with the desired ingredients and return a map of all ingredients with their associate macros
// you have to fill the macros by hand and to change the value in the return field
func muesliMacros() map[string]map[string]float32 {

	kamut := createMapMacros(9.9, 74, 29, 0.9, 0, 6.2, 356)
	billetsEpeautre := createMapMacros(14.5, 60.9, 1, 3.1, 0.4, 6.2, 351)
	flakesEpeautre := createMapMacros(14, 68, 4.9, 1.5, 0.2, 10, 362)
	avoine := createMapMacros(13, 59.1, 1.1, 7, 1.3, 13, 372)
	amandes := createMapMacros(22.1, 5.4, 0.1, 54.1, 4.1, 13.5, 583)
	pommes := createMapMacros(1.3, 92.2, 77.4, 0.4, 0.11, 13.4, 378)
	myrtilles := createMapMacros(2.93, 77.3, 51.4, 2.43, 0, 8, 359)

	return map[string]map[string]float32{"kamut": kamut, "billetsEpeautre": billetsEpeautre, "flakesEpeautre": flakesEpeautre, "avoine": avoine, "amandes": amandes, "pommes": pommes, "myrtilles": myrtilles}
}

// macrosPercentage take an ingredient and calculate the macros with the given percentage
func macrosPercentage(mp map[string]float32, p float32) map[string]float32 {

	// define the possible percentage between 0.05 and 0.99
	if p < 0 && p > 0.99 {
		log.Fatal("the percentage need to be set between 0.05 - 0.99")
	}

	for i, m := range mp {
		mp[i] = m * p
	}

	return mp
}

// here we are going to calculate the macros for the total with differents given percentage per ingredient
func main() {
	macros := muesliMacros()

	// create the final macros
	var p, g, s, l, as, f, kcal float32

	// define the proportion you want
	// the percentage total must be equal to 1
	kamut := macrosPercentage(macros["kamut"], 0)
	be := macrosPercentage(macros["billetsEpeautre"], 0.2)
	fe := macrosPercentage(macros["flakesEpeautre"], 0.2)
	avoine := macrosPercentage(macros["avoine"], 0.4)
	amandes := macrosPercentage(macros["amandes"], 0.1)
	pommes := macrosPercentage(macros["pommes"], 0.1)
	myrt := macrosPercentage(macros["myrtilles"], 0)

	p = kamut["p"] + be["p"] + fe["p"] + avoine["p"] + amandes["p"] + pommes["p"] + myrt["p"]
	g = kamut["g"] + be["g"] + fe["g"] + avoine["g"] + amandes["g"] + pommes["g"] + myrt["g"]
	s = kamut["s"] + be["s"] + fe["s"] + avoine["s"] + amandes["s"] + pommes["s"] + myrt["s"]
	l = kamut["l"] + be["l"] + fe["l"] + avoine["l"] + amandes["l"] + pommes["l"] + myrt["l"]
	as = kamut["as"] + be["as"] + fe["as"] + avoine["as"] + amandes["as"] + pommes["as"] + myrt["as"]
	f = kamut["f"] + be["f"] + fe["f"] + avoine["f"] + amandes["f"] + pommes["f"] + myrt["f"]
	kcal = kamut["kcal"] + be["kcal"] + fe["kcal"] + avoine["kcal"] + amandes["kcal"] + pommes["kcal"] + myrt["kcal"]

	fmt.Printf("P: %f\n", p)
	fmt.Printf("G: %f\n", g)
	fmt.Printf("S: %f\n", s)
	fmt.Printf("L: %f\n", l)
	fmt.Printf("AS: %f\n", as)
	fmt.Printf("F: %f\n", f)
	fmt.Printf("Kcal: %f\n", kcal)
}
