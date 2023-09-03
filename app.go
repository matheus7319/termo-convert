package main

import "fmt"

var valueInKelvin float64 = 467

func main() {
    newValue := convertToCelsius(valueInKelvin)
    fmt.Println("O valor convertido de Kelvin para Celsius é: ", newValue)
}

func convertToCelsius(K float64) (C float64) {
    C = K - 273.0
    return C
}