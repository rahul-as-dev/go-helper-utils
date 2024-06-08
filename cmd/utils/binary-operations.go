package utils

import "fmt"

func ConvertToBits() {
	num := 56
	binaryString := fmt.Sprintf("%b", num)
	fmt.Printf("Binary representation of %d is %s\n", num, binaryString)
}
