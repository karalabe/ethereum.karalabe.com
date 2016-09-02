contract KingOfTheEtherThrone {
	address monarch;

	function claim() {
		// [...] calculate the ruler's compensation

		monarch.send(compensation);
		monarch = msg.sender;
	}
	[...]
}
