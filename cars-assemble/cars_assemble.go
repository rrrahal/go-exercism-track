package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100.0)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	totalCars := CalculateWorkingCarsPerHour(productionRate, successRate)
	return int(totalCars) / 60
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	numberOfIndividuals := uint(carsCount % 10)
	numberOfGroups := uint(carsCount / 10)

	return numberOfGroups*95000 + numberOfIndividuals*10000
}
