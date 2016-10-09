func NewGopher(name string, pic []byte, parents []*Gopher) *Gopher {
	// [...] Init basic fields as previously

	for i, parent := range parents { // Dereference any pointers // HL
		cpy := *parent      // HL
		g.parents[i] = &cpy // HL
	} // HL
	return g
}
