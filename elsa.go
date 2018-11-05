package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const LEN int = 512

func main() {
	var chain []MarkovNode
	var inputTokens []string
	rand.Seed(int64(time.Now().Nanosecond()))

	in := bufio.NewScanner(bufio.NewReader(os.Stdin))
	in.Split(bufio.ScanWords)

	for in.Scan() {
		inputTokens = append(inputTokens, in.Text())
	}
	chain = createChain(inputTokens)
	fmt.Println(markov(chain, LEN))
}
