package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
	if time == 0 {
		return len(layers) * 2
	}

	return len(layers) * time
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	n_noodles := 0
	n_sauce := 0
	for i := 0; i < len(layers); i++ {
		if layers[i] == "sauce" {
			n_sauce++
		}
		if layers[i] == "noodles" {
			n_noodles++
		}
	}
	return n_noodles * 50, float64(n_sauce) * 0.2
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendReceipt, myReceipt []string) {
	myReceipt[len(myReceipt)-1] = friendReceipt[len(friendReceipt)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quant []float64, scale int) []float64 {
	newScale := []float64{}
	for i := 0; i < len(quant); i++ {
		newScale = append(newScale, quant[i]/2*float64(scale))
	}
	return newScale
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
