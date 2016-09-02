contract Faucet {
	uint amount = 0.01 ether;
	uint freq   = 5760;

	mapping (address => uint) prev;

	function request() {
		if (address(this).balance < amount) {
			return;
		}
		if(block.number < prev[msg.sender] + freq) {
			return;
		}
		msg.sender.send(amount);
		prev[msg.sender] = block.number;
	}
}
