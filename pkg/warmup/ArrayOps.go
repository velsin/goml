package warmup

// Probably want to use a Generic? We don't really care about array element type.

func GetIntersection[T comparable](a []T, b []T) []T {
	// Returns intersection of the arrays a and b
	// Uses a hashmap for linear time complexity, each array only needs to be
	// iterated over once, and we use the constant lookup time of the hashmap
	// to check elements
	var res []T
	hashmap := make(map[T]bool)

	// Populate the hashmap
	for _, v := range a {
		hashmap[v] = true
	}
	for _, v := range b {
		if hashmap[v] == true {
			res = append(res, v)
		}
	}
	return res
}

func GetUnion[T comparable](a []T, b []T) []T {
	// Returns the union of the arrays a and b
	// Uses hashmap to store a frequency independent membership list of each element
	var res []T
	m := make(map[T]bool)
	for _, v := range a {
		m[v] = true
	}
	for _, v := range b {
		m[v] = true
	}
	for k := range m {
		res = append(res, k)
	}

	return res
}
