type Gopher struct {
	name    string
	picture []byte
	parents [2]*Gopher // Lightweight immutables // HL
}
