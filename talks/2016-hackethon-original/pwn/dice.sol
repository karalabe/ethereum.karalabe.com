contract Dice {
	function evaluate(bytes32 seed) {
		// [...] verify the seed
		for (uint i = 0; i < bets.length; i++) {
			// [...] evaluate bets and pay winners
		}
	}
	[...]
}
