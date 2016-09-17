Address oracleAddress = Geth.newAddressFromHex("0xfa7b9770ca4cb04296cac84f37736d4041251cdf");

EthereumClient ec = Geth.newEthereumClient("ws://10.0.2.2:8546");
ReleaseOracle ro = new ReleaseOracle(oracleAddress, ec); // HL

ReleaseOracle.CurrentVersionResults ver = ro.currentVersion(null); // HL
textbox.append("Latest Geth: v" + ver.Major + "." + ver.Minor + "." + ver.Patch);
