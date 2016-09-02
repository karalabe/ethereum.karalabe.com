contract MultiOwned {
	mapping(uint => uint) owners;

	modifier onlyowner {
		if (owners[tx.origin] > 0) {
			_
		}
	}
	[...]
}
