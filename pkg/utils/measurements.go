package utils

import (
	"fmt"
	"strconv"
)

func MeasurementStatus(value string) string {
	number, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println("Error parsing integer:", err)
	}
	switch {
	case number <= 100 || number >= 400:
		return fmt.Sprintf("%v: Valor peligroso", number)
	case number <= 150 || number >= 300:
		return fmt.Sprintf("%v: Valor moderadamente peligroso", number)
	default:
		return fmt.Sprintf("%v: Valor correcto", number)
	}
}
