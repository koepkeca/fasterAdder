package fasterAdder

//Summer is a function that provides a concurrent adder for a slice of 64-bit integers.
//It accepts a slice of integers and the number of partitions (go routines) to use to compute
//the sum of the integers in the given slice. The number of partitions is required to be between
//1 and eight, and numbers out of that bound are automatically selected. [zero becomes one, anything
//greater than 8 is eight. The function returns the sum of the integers in i.
func Summer(i []int64, partCt int64) (sum int64) {
	//sanity check to make sure the go routines do not get out of hand.
	if partCt == 0 {
		partCt = 1
	}
	if partCt > 8 {
		partCt = 8
	}
	//buffered channel to receive the sums of the partitions
	sumCh := make(chan int64, partCt)
	for z := int64(0); z < partCt; z++ {
		start := z * (int64(len(i)) / partCt)
		end := (z + 1) * (int64(len(i)) / partCt)
		//This solves an edge case surrounding the iteration counters and
		//a slice with an odd number of elements.
		if (z + 1) == partCt {
			end = int64(len(i))
		}
		//concurrent routine begins here.
		go func(s, f int64) {
			tmpSum := int64(0)
			for _, next := range i[s:f] {
				tmpSum += next
			}
			sumCh <- tmpSum
			return
		}(start, end)
	}
	//accept and sum the result for each go routine.
	for cr := int64(0); cr < partCt; cr++ {
		sum += <-sumCh
	}
	return
}
