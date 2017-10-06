contract Crowdsale {
	// Adományozók listája
	struct Funder {
		address addr;   // Adományozó címe
		uint    amount; // Adomány összege
	}
	Funder[] public funders;

	// Finanszírozás részletei
	address public project;  // Projekt kontja aki kapja a pénzt ha összegyűl
	uint    public target;   // Pénz mennyisége ami össze kell gyűljön
	uint    public raised;   // Eddigi összegyűlt pénzmennyiség
	uint    public deadline; // Időpont amíg össze kell gyűljön

  // Constructor ami inicializálja a szerződésünket
	function Crowdsale(address owner, uint goal, uint duration) {
		project  = owner;
		target   = goal;
		deadline = now + duration * 1 days;
	}

	// Ha valaki adományoz, felírjuk a nagy könyvbe
	function donate() {
		funders[funders.length++] = Funder({addr: msg.sender, amount: msg.value});
		raised += msg.value;
	}

	// Ha lejárt az idő, vagy kiutaljuk a pénzt, vagy vissza mindenkinek
	function terminate() {
		if (now >= deadline) {
			if (raised >= target) {
				project.send(raised);
				return;
			}
			for (uint i = 0; i < funders.length; ++i) {
				funders[i].addr.send(funders[i].amount);
			}
		}
	}
}
