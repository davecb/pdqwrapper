package main

import (
	"fmt"
	iterator "github.com/davecb/pdqwrapper/tests/testIterator/iterator"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	// Create collection
	numberCollection := iterator.NewNumberCollection(numbers)

	// Get iterator
	iter := numberCollection.CreateIterator()

	// Iterate through collection
	for iter.HasNext() {
		num := iter.Next()
		fmt.Printf("%d ", num)
	}
}
