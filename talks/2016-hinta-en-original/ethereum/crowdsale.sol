contract Crowdsale {
	// List of funders
	struct Funder {
		address addr;   // Funder's address
		uint    amount; // Funding amount
	}
	Funder[] public funders;

	// Details of the crowdsale
	address public project;  // Project account to fund if enough is collected
	uint    public target;   // Target amount to successfully fund the project
	uint    public raised;   // Amount of funds raised until this point in time
	uint    public deadline; // Deadline until donations may be sent in

  // Constructor to initialize the smar contract.
	function Crowdsale(address owner, uint goal, uint duration) {
		project  = owner;
		target   = goal;
		deadline = now + duration * 1 days;
	}

	// If someone donates, we add him to the big book.
	function donate() {
		funders[funders.length++] = Funder({addr: msg.sender, amount: msg.value});
		raised += msg.value;
	}

	// If wthe deadline is reached, either forward or return the funds.
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
