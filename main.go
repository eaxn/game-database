package main

func main() {
	// kurz := true
	// fmt.Println(Steam_GameIDExists(big.NewInt(0)))
	// for i := 105600; i < 1056001212; i++ {
	// 	if kurz {
	// 		game := Steam_GetProductFromGameID(big.NewInt(int64(i)))
	// 		t := ""
	// 		if game.Name == "Game does not exist" {
	// 			t += "❌"
	// 		} else {
	// 			t += "✔️"
	// 		}
	// 		fmt.Println(game.ID, "-", t)
	// 	} else {
	// 		game := Steam_GetProductFromGameID(big.NewInt(int64(i)))
	// 		fmt.Println("----------------------------------")
	// 		fmt.Println("ID:", i)
	// 		fmt.Println("NAME:", game.Name)
	// 		fmt.Println("PRICE:", fmt.Sprint(game.Price)+game.Currency.Symbol)
	// 		fmt.Println("URL:", game.URL)
	// 		fmt.Println("---------------------------------")
	// 	}
	// }
	Steam_CreateDatabase("./steam")
}
