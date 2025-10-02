package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(filename string) (float64, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return 1000, errors.New(fmt.Sprint("Error reading file: ", err))
	}
	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 1000, errors.New(fmt.Sprint("Error parsing value: ", err))
	}
	return value, nil
}

func WriteFloatToFile(value float64, filename string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(filename, []byte(valueText), 0644)
}
