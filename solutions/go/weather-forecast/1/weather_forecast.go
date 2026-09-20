// Package weather provides a tool for reporting
// the current weather condition in a specific location.
package weather

var (
    // CurrentCondition represents a current weather condition.
	CurrentCondition string
    // CurrentLocation represents a current location of the weather forecast.
	CurrentLocation  string
)

// Forecast returns a string value representing
// current weather condition in a specific city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
