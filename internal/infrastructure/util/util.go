package util

import (
	"fmt"
	"log"
	"math/big"
	"time"
)

func Elapsed(fun func()) float64 {
	start := time.Now()

	r := new(big.Int)
	fmt.Println(r.Binomial(1000, 10))

	fun()
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
	return elapsed.Seconds()
}
