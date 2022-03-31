package main

import "strings"

type Currency struct {
	Name   string
	Symbol string
}

// add euro
var Euro = Currency{"Euro", "€"}

// add dollar
var Dollar = Currency{"Dollar", "$"}

// add pound
var Pound = Currency{"Pound", "£"}

// add yen
var Yen = Currency{"Yen", "¥"}

// add rupee
var Rupee = Currency{"Rupee", "₹"}

// add won
var Won = Currency{"Won", "₩"}

// add ruble
var Ruble = Currency{"Ruble", "₽"}

// add zloty
var Zloty = Currency{"Zloty", "zł"}

// add swiss franc
var SwissFranc = Currency{"Swiss Franc", "CHF"}

// add sweedish krona
var SweedishKrona = Currency{"Sweedish Krona", "kr"}

// add finnish markka
var FinnishMarkka = Currency{"Finnish Markka", "mk"}

// add vietnamese dong
var VietnameseDong = Currency{"Vietnamese Dong", "₫"}

var AllCurrencySymbols = []string{Euro.Symbol, Dollar.Symbol, Pound.Symbol, Yen.Symbol, Rupee.Symbol, Won.Symbol, Ruble.Symbol, Zloty.Symbol, SwissFranc.Symbol, SweedishKrona.Symbol, FinnishMarkka.Symbol, VietnameseDong.Symbol}

var AllCurrencies = []Currency{Euro, Dollar, Pound, Yen, Rupee, Won, Ruble, Zloty, SwissFranc, SweedishKrona, FinnishMarkka, VietnameseDong}

func GetCurrencyBySymbol(Symbol string) Currency {
	for x := 0; x < len(AllCurrencies); x++ {
		if AllCurrencies[x].Symbol == Symbol {
			return AllCurrencies[x]
		}
	}
	return Dollar
}

func ParseCurrency(Price string) Currency {
	// iterate over all currency symbols
	text := Price
	for _, currency := range AllCurrencySymbols {
		if strings.Contains(text, currency) {
			return GetCurrencyBySymbol(currency)
		}
	}
	return Dollar
}
