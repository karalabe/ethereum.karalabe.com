contract GovernMental {
	address[] creditors; uint[] credited;

	function invest() {
		if (block.timestamp - lastInvested < TWELVE_HOURS) { ... } else {
			creditors = new address[](0);
			credited  = new uint[](0);
		}
	}
	[...]
}
