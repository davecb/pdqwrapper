package iterator

type NumberCollection struct {
	numbers []int
}

type NumberIterator struct {
	index   int
	numbers []int
}

// Create new collection
func NewNumberCollection(numbers []int) *NumberCollection {
	return &NumberCollection{numbers: numbers}
}

// Create iterator for the collection
func (nc *NumberCollection) CreateIterator() *NumberIterator {
	return &NumberIterator{
		index:   0,
		numbers: nc.numbers,
	}
}

// Iterator methods
func (ni *NumberIterator) HasNext() bool {
	return ni.index < len(ni.numbers)
}

func (ni *NumberIterator) Next() int {
	if ni.HasNext() {
		number := ni.numbers[ni.index]
		ni.index++
		return number
	}
	return 0
}
