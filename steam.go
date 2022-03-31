package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"
)

type SteamGame struct {
	ID       *big.Int
	Name     string
	URL      string
	RawPrice string
	Price    float32 // can be converted to string (fmt.Sprint)
	Currency Currency
}

func Steam_GetProductFromGameID(ID *big.Int) SteamGame {
	Name := Steam_GameNameFromID(ID)
	RawPrice := Steam_GamePriceFromID(ID)
	Price, err := strconv.ParseFloat(ParsePrice(RawPrice), 32)
	if err != nil {
		Price = -1
	}
	if RawPrice == "invalid" {
		Price = -2
	}
	// -1 = no game
	// -2 = no price
	return SteamGame{ID, Name, Steam_GetProductURLFromGameID(ID), RawPrice, float32(Price), ParseCurrency(RawPrice)}
}

func Steam_GetProductURLFromGameID(ID *big.Int) string {
	URL := "https://store.steampowered.com/app/" + ID.String() + "/"
	return URL
}

func Steam_GetStartPage() []byte {
	url := Steam_GetProductURLFromGameID(big.NewInt(0))
	response, err := httpGet(url)
	if err != nil {
		log.Println(err.Error())
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err.Error())
	}
	return body
}

func ParsePrice(Price string) string {
	if !strings.Contains(Price, ">") {
		// if only contains digits and dots, then it is a price
		Result := ""
		Price = strings.Replace(Price, ",", ".", 1)
		for x := 0; x < len(Price); x++ {
			// if price[x] is not a digit or a dot, then it is not a price
			if strings.Contains("0123456789,.", string(Price[x])) {
				// check for , and replace with .
				if Price[x] == ',' {
					Price = strings.Replace(Price, ",", ".", 1)
				}
				Result += string(Price[x])
			}
		}
		return Result
	}
	Price = strings.Split(Price, ">")[1]
	CurrencySymbol := ParseCurrency(Price).Symbol
	Price = strings.Split(Price, CurrencySymbol)[0] // get content before currency symbol
	Price = strings.Replace(Price, ",", ".", -1)    // replace comma with dot
	write := false
	Result := ""
	for x := 0; x < len(Price); x++ {
		// check if character is a number
		if Price[x] >= '0' && Price[x] <= '9' {
			write = true
		}
		if write {
			Result += string(Price[x])
		}
	}
	return Result
}

func Steam_GameNameFromID(ID *big.Int) string {
	url := Steam_GetProductURLFromGameID(ID)
	response, err := httpGet(url)
	if err != nil {
		log.Println(err.Error())
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err.Error())
	}
	defer response.Body.Close()
	// extract <title> from body
	text := string(body)
	if !strings.Contains(text, "<title>") || !strings.Contains(text, "</title>") {
		return "Game does not exist"
	}

	text = strings.Split(text, "<title>")[1]
	text = strings.Split(text, "</title>")[0]
	if text == "Welcome to Steam" {
		return "Game does not exist"
	}
	// remove " on Steam" as suffix from text
	text = strings.Split(text, " on Steam")[0]
	return text
}

func Steam_GamePriceFromID(ID *big.Int) string {
	url := Steam_GetProductURLFromGameID(ID)
	response, err := httpGet(url)
	if err != nil {
		log.Println(err.Error())
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err.Error())
	}
	defer response.Body.Close()
	// extract <title> from body
	title := string(body)
	if !strings.Contains(title, "<title>") || !strings.Contains(title, "</title>") {
		return "Game does not exist"
	}
	title = strings.Split(title, "<title>")[1]
	title = strings.Split(title, "</title>")[0]
	if title == "Welcome to Steam" {
		return "Game does not exist"
	}
	text := string(body)

	if strings.Contains(text, "<meta itemprop=\"price\" content=\"") {
		if strings.Contains(text, "\">") {
			text = strings.Split(text, "<meta itemprop=\"price\" content=\"")[1]
			text = strings.Split(text, "\">")[0]
			return ParsePrice(text)
		}
	}

	// check if <div class="game_purchase_price price" data-price-final= exists in text
	if strings.Contains(text, "<div class=\"game_purchase_price price\" data-price-final=") {
		// check if </div> exists in text
		if strings.Contains(text, "</div>") {
			text = strings.Split(text, "<div class=\"game_purchase_price price\" data-price-final=")[1]
			text = strings.Split(text, "</div>")[0]
			return ParsePrice(text)
		}
	}

	// check if <div class="game_purchase_price price"> exists in text
	if !strings.Contains(text, "<div class=\"game_purchase_price price\">") {
		return "invalid"
	}
	if strings.Contains(text, "Free") {
		return "0"
	} // check if </div> exists in text
	text = strings.Split(text, "<div class=\"game_purchase_price price\" data-price-final=\"")[1]

	if !strings.Contains(text, "</div>") {
		return "invalid"
	}
	text = strings.Split(text, "</div>")[0]
	text = ParsePrice(text)
	return text
}

func Steam_GameIDExists(ID *big.Int) bool {
	url := Steam_GetProductURLFromGameID(ID)
	response, err := httpGet(url)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	defer response.Body.Close()
	if !strings.Contains(string(body), "<title>") || !strings.Contains(string(body), "</title>") {
		return false
	}
	return !strings.Contains(string(body), "Welcome to Steam")
}

var currentID = big.NewInt(139404) // starts at 139404
var biggestNoMoreIDs = 0
var noMoreIDs = 0

func GetNextID() *big.Int {
	x := big.NewInt(0)
	x.Set(currentID)
	currentID.Add(currentID, big.NewInt(1))
	return x
}

func GetCurrentID() *big.Int {
	x := big.NewInt(0)
	x.Set(currentID)
	return x
}

func x() {
	for {
		// write currentid to currentid.bak
		ioutil.WriteFile("current.bak", []byte(currentID.String()+"\n"+fmt.Sprint(biggestNoMoreIDs)), 0644)
		time.Sleep(time.Second * 100)
	}
}

func Steam_CreateDatabase(locationPath string) {
	go x()
	workers := 10
	// try to open locationPath/1.data
	end := big.NewInt(0)
	start := big.NewInt(0)
	startMilli := time.Now().UnixMilli()

	w := func() {
		for {
			if biggestNoMoreIDs < noMoreIDs {
				biggestNoMoreIDs = noMoreIDs
			}
			x := big.NewInt(0)
			x.Set(GetNextID())
			id := x
			if noMoreIDs > 200000 { // more than 200000
				fmt.Println(id.String(), "🔴", noMoreIDs)
				break
			}
			if !Steam_GameIDExists(id) {
				fmt.Println(id.String(), "🔴", noMoreIDs)
				noMoreIDs++
				continue
			}
			if start.Cmp(big.NewInt(-1)) == 0 {
				start = id
			}
			end = id
			// a = gameID
			game := Steam_GetProductFromGameID(id)
			if len(game.Name) == 0 {
				fmt.Println(id.String(), "🔴", noMoreIDs)
				noMoreIDs++
				continue
			} else {
			}
			noMoreIDs = 0
			folderName := fmt.Sprint(int32(game.Name[0]))
			// checks if folder exists
			if _, err := os.Stat(locationPath + "/" + folderName); os.IsNotExist(err) {
				// if not, create it
				os.Mkdir(locationPath+"/"+folderName, 0777)
			}

			targetFilePath := locationPath + "/" + folderName + "/" + id.String() + ".game"
			// try to open the file path
			_, err := os.Open(targetFilePath)
			// no warning, because we want to check if the file is already there
			// and if the file is not existing, we don't want to warn the user
			if err != nil {
				fmt.Println(id.String(), "🔵🔵", game.Name)
				// file does not exist, so create it
				file, err := os.Create(locationPath + "/" + folderName + "/" + id.String() + ".game")
				if err != nil {
					log.Println(err.Error())
				}
				// write name of the game to file
				_, err = file.WriteString(game.Name)
				if err != nil {
					log.Println(err.Error())
				}
				// write price of the game to file
				_, err = file.WriteString("\n" + fmt.Sprint(game.Price))
				if err != nil {
					log.Println(err.Error())
				}
				// close the file
				defer file.Close()
			} else {
				fmt.Println(id.String(), "🟢🔵", game.Name)
				// fmt.Println("❎")
				// fmt.Println("❎")
				// fmt.Print("✅(")
				// fmt.Print(game.Name + ") [")
				// fmt.Print(fmt.Sprint(game.Price) + "] ")
				// fmt.Println("=> create 🟢 {", id.String(), "}")
				// fmt.Println("=> update 🟡 {", id.String(), "}")

				// file does not exist, so create it
				file, err := os.OpenFile(locationPath+"/"+folderName+"/"+id.String()+".game", os.O_RDWR, 0777)
				if err != nil {
					log.Println(err.Error())
				}
				// write name of the game to file
				_, err = file.WriteString(game.Name)
				if err != nil {
					log.Println(err.Error())
				}
				// write price of the game to file
				_, err = file.WriteString("\n" + fmt.Sprint(game.Price))
				if err != nil {
					log.Println(err.Error())
				}
				// close the file
				defer file.Close()
			}
		}
		// start => ?
		// end => ?
	}
	for x := 0; x < workers-1; x++ {
		go w()
	}
	w()

	fmt.Println("starts at", start)
	fmt.Println("ends at", end)
	endmilli := time.Now().UnixMilli()
	fmt.Println("time elapsed:", endmilli-startMilli, "ms")
}
