func (g *Gopher) Name() string    { return g.name } // Shallow copy immutable // HL
func (g *Gopher) Picture() []byte {                 // Deep copy mutable // HL
	pic := make([]byte, len(g.picture))
	copy(pic, g.picture)
	return pic
}
