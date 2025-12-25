package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/rahul-as-dev/go-helper-utils/dsa"
)

func Solve() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int

	// It automatically skips spaces/newlines and converts types.
	fmt.Fscan(reader, &n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}
	fmt.Fprintln(writer, arr)
	fmt.Fprintln(writer, dsa.GreaterElementsToRight(arr))
}
