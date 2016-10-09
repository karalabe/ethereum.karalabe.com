func (g *Gopher) Parent(idx int) *Gopher {
	cpy := *g.parents[idx] // Internal pointers never escape // HL
	return &cpy            // HL
}
