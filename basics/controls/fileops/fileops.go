package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Get in this case must be uppercase to export to other packages
func GetFloat(fileName string) (float64, error) {
	// _ placeholder for a variable
	data, err := os.ReadFile(fileName)
	// nil means no error
	if err != nil {
		return 1000, errors.New("can't find file")
	}
	valueText := string(data)
	// converts string to float64
	balance, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 1000, errors.New("Can't parse float")
	}
	return balance, nil
}

// Write uppercase to export
func WriteFloat(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	// writes the balance amount to a text file, has a permission of 0644
	os.WriteFile(fileName, []byte(valueText), 0644)
}
