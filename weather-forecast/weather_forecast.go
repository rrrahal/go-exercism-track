// Package weather provides utilities to weather prediction in Goblinocus.
package weather

// CurrentCondition holds a string representing the current condition in the city.
var CurrentCondition string

// CurrentLocation holds a string representing the location of Goblinocus.
var CurrentLocation string

// Forecast is a function that given a city and a condition will return the
// forecast for that city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
