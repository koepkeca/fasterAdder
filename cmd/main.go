package main

import (
	"log"
	"math/rand"
	"time"
	"flag"

	"github.com/koepkeca/fasterAdder"
)

func main() {
	randMax := flag.Int64("randMax", 1024, "Maximum value for random Int64 values for randomly generated input data")
	sliceLen := flag.Int64("sliceLen", 1024000, "Number of random values to generate for the slice (random slice length")
	parts := flag.Int("parts", 2, "Number of partitions (co-routines) to divide the slice into for adding. Max value is 8. Values greater than 8 are set to 8.")
	flag.Parse()
	st := time.Now()
	log.Printf("Generating %d random numbers for the slice [0,%d]",*sliceLen, *randMax)
	rnList := []int64{}
	rand.Seed(time.Now().UTC().UnixNano())
	for i := int64(0); i < *sliceLen; i++ {
		next := rand.Int63n(*randMax)
		rnList = append(rnList, next)
	}
	et := time.Since(st)
	log.Printf("Random slice creation took %s", et)
	st = time.Now()
	log.Printf("Concurrent Addition Start\n")
	log.Printf("Sum: %d\n", fasterAdder.Summer(rnList, int64(*parts)))
	et = time.Since(st)
	log.Printf("Concurrent summation took %s", et)
	//To compare the times to a sequential addition, we loop and add the input.
	//This is a standard O(n) addition of a slice of integers.
	log.Printf("Linear Addition Start\n")
	st = time.Now()
	tmp := int64(0)
	for _, next := range rnList {
		tmp += next
	}
	log.Printf("Sum: %d\n", tmp)
	et = time.Since(st)
	log.Printf("Linear summation took %s", et)
	return
}
