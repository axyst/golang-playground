package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
	if time <= 0 {
		return 2 * len(layers)
	}
	return time * len(layers)
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	var reti int
	var retf float64
	for _, i := range layers {
		if i == "sauce" {
			retf += 0.2
		} else if i == "noodles" {
			reti += 50
		}
	}
	return reti, retf
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, amounts int) []float64 {
	var ret []float64
	for _, i := range quantities {
		ret = append(ret, i*float64(amounts)/2)
	}
	return ret
}
