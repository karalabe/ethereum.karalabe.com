EthereumClient ec = Geth.newEthereumClient("ws://10.0.2.2:8546");

Block head = ec.getBlockByNumber(ctx, -1); // HL
textbox.append("Chain height: " + head.getNumber() + "\n\n");

BigInt balance = ec.getBalanceAt(ctx, tipjar, -1); // HL
textbox.append("TipJar funds: " + balance + "\n\n");

NewHeadHandler handler = new NewHeadHandler() { // HL
	@Override public void onNewHead(final Header header) {
		MainActivity.this.runOnUiThread(new Runnable() { public void run() { textbox.append("#" + header.getNumber() + ": " + header.getHash().getHex().substring(0, 10) + "…\n"); } });
	}};
ec.subscribeNewHead(ctx, handler,  16); // HL
