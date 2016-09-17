package main

import "fmt"

// START OMIT
import "github.com/ethereum/go-ethereum/accounts"

func main() {
	am := accounts.NewManager("keystore", accounts.StandardScryptN, accounts.StandardScryptP)

	newAcc, _ := am.NewAccount("Creation password") // HL
	fmt.Println("New:", newAcc.Address.Hex())

	jsonAcc, _ := am.Export(newAcc, "Creation password", "Export password") // HL
	fmt.Println("Json:", string(jsonAcc))

	_ = am.DeleteAccount(newAcc, "Creation password") // HL
	fmt.Println("Accs:", am.Accounts())

	impAcc, _ := am.Import(jsonAcc, "Export password", "Import password") // HL
	fmt.Println("Imp:", impAcc.Address.Hex())
}

// END OMIT
