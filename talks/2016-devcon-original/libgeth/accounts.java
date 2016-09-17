AccountManager am = Geth.newAccountManager(this.getFilesDir(), Geth.LightScryptN, Geth.LightScryptP);

Account newAcc = am.newAccount("Creation password"); // HL
textbox.append("New: " + newAcc.getAddress().getHex() + "\n\n");

byte[] jsonAcc = am.exportKey(newAcc, "Creation password", "Export password"); // HL
textbox.append("Json: " + new String(jsonAcc) + "\n\n");

am.deleteAccount(newAcc, "Creation password"); // HL
textbox.append("Accs: " + am.getAccounts().size() + "\n\n");

Account impAcc = am.importKey(jsonAcc, "Export password", "Import password"); // HL
textbox.append("Imp: " + impAcc.getAddress().getHex() + "\n\n");
