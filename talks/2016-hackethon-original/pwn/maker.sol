contract MakerEthToken {
	function withdraw(uint amount) {
		if (balances[msg.sender] >= amount) {
			if (msg.sender.call.value(amount)()) {
				balances[msg.sender] -= amount;
			}
		}
	}
	[...]
}
