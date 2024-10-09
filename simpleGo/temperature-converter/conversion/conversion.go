package conversion

const (
	Kelvin     = "k"
	Celsius    = "c"
	Fahrenheit = "f"
)

type Temperature struct {
	Value float64
	Unit  string
}

// var T temp.Temperature

func FahrenheitToCelcius(f float64) float64 {
	return (f - 32) * 5 / 9
}

func FahrenheitToKelvin(f float64) float64 {
	return (f + 459.67) * 5 / 9
}

func CelciusToFahrenheit(c float64) float64 {
	return c*9/5 + 32
}

func CelciusToKelvin(c float64) float64 {
	return c + 273.15
}

func KelvinToFahrenheit(k float64) float64 {
	return k*9/5 - 459.67
}

func KelvinToCelcius(k float64) float64 {
	return k - 273.15
}
