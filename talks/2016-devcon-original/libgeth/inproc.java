NodeConfig conf = Geth.newNodeConfig();
Node node = Geth.newNode(getFilesDir() + "/.ethereum", conf); // HL
node.start(); // HL

NodeInfo info = node.getNodeInfo(); // HL
textbox.append("My name: " + info.getName() + "\n");
textbox.append("My address: " + info.getListenerAddress() + "\n");
textbox.append("My protocols: " + info.getProtocols() + "\n\n");

EthereumClient ec = node.getEthereumClient(); // HL
textbox.append("Latest block: " + ec.getBlockByNumber(ctx, -1).getNumber());
