func ExtendFamilyTree(gopher, ancestor *Gopher, ancestry []int) *Gopher {
	mutable := *gopher // Reconstruct all objects on the update path // HL

	if len(ancestry) == 0 {
		// Mutate the object at the deepest level // HL
		if mutable.parents[0] == nil {
			mutable.parents[0] = ancestor
		} else {
			mutable.parents[1] = ancestor
		}
	} else {
		// Unwind and recombine objects on the update path // HL
		index := ancestry[0]
		parent := mutable.parents[ancestry[0]]

		mutable.parents[index] = ExtendFamilyTree(parent, ancestry[1:], ancestor)
	}
	return &mutable
}
