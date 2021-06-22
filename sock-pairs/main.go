package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// the function takes an array of length n and returns pairs
// ie [red red red red blue yellow blue]
// and returns the number of pairs of socks in the array
// ie 3

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	socksToMatch := strings.Split(input, " ")

	fmt.Println(getPairs(socksToMatch))

}

func getPairs(socksToMatch []string) int32 {
	socks := make(map[string]int32)
	var count int32
	arrLen := int32(len(socksToMatch))
	for i := int32(0); i < arrLen-1; i++ {
		socks[string(socksToMatch[i])]++
	}
	for key := range socks {
		count += socks[key] / 2
	}
	return count
}
