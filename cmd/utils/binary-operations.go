package utils

import (
	"fmt"
	"strconv"
)

func ConvertIntToBits() string {
	num := 56
	binaryString := fmt.Sprintf("%b", num)
	// binaryString := strconv.FormatInt(int64(num), 2)
	fmt.Printf("Binary representation of %d (int64) is %s\n", num, binaryString)
	return binaryString
}

func ConnvertBitsToInt() int64 {
	bitStr := "111000"
	num, err := strconv.ParseInt(bitStr, 2, 64)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("int64 representation of binary string %s is %d\n", bitStr, num)
	return num
}
