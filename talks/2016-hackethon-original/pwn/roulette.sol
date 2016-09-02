contract Roulette {
	uint seed = 1;

	function rand() private returns (uint) {
		seed = ((seed*3 + 1)/2 % 10**9);
		
		return (seed + block.number + block.difficulty + block.timestamp + block.gaslimit) % 37;
	}
	[...]
}
