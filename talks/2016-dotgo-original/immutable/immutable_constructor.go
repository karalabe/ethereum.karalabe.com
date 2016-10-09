func NewGopher(name string, pic []byte) *Gopher {
	g := &Gopher{
		name:    name,                   // Shallow copy immutable // HL
		picture: make([]byte, len(pic)), // Deep copy mutable // HL
	}
	copy(g.picture, pic)
	return g
}
