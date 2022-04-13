package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const layout = "2006-01-02T15:04:05.000Z"

var portfolio Portfolio
var url string
var debug bool
var stocks []Stock

func init() {
	if os.Getenv("POPULATE") == "true" {
		populate("2020-01-01T00:00:00.000Z", "2022-04-30T00:00:00.000Z")
	}

	if os.Getenv("DEBUG") == "true" {
		debug = true
	}
	// Is not working
	// "https://api.iextrading.com/1.0/stock/market/batch?symbols=aapl,tsla&types=quote&range=1m&last=1"
	if os.Getenv("STOCKS_URL") != "" {
		url = os.Getenv("STOCKS_URL")
	} else {
		// Uploaded to my personal RPI server
		url = "https://www.unnerv.xyz/stocks.json"
	}

	// Loading the stocks from the json file
	stocks = request()

	// Addinng stocks to the portfolio
	portfolio.Add("AAPL", 10) // 10 Apple's stock
	portfolio.Add("GOOG", 5)  // 5 Google's stock
	portfolio.Add("AMZN", 7)  // 7 Amazon's stock
}

func main() {
	t1 := time.Now()
	start, _ := time.Parse(layout, "2020-01-01T00:00:00.000Z")
	end, _ := time.Parse(layout, "2021-01-01T00:00:00.000Z")

	profit, ar := portfolio.Profit(start, end)

	if profit >= 0 {
		// Print green
		fmt.Printf("\x1b[32mProfit: %f\x1b[0m\n", profit)
		fmt.Printf("\x1b[32mAnnualized Return: %f\x1b[0m\n", ar)
	} else {
		// Print red
		fmt.Printf("\x1b[31mLoss: %f\x1b[0m\n", profit)
		fmt.Printf("\x1b[31mAnnualized Return: %f\x1b[0m\n", ar)
	}

	stock1 := portfolio.Stocks[0] // Apple (first one in my portfolio)
	price1 := stock1.GetPrice(start)

	stock2 := Stock{Name: "TSLA"} // Any other stock
	price2 := stock2.GetPrice(start)

	// Printing the prices of the stocks
	fmt.Printf("\x1b[32mPrice: %f\x1b[0m\n", price1)
	fmt.Printf("\x1b[32mPrice: %f\x1b[0m\n", price2)
	elapsed := time.Since(t1)
	fmt.Printf("\x1b[32mTime elapsed: %s\x1b[0m\n", elapsed)
}

func request() []Stock {
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var stocks []Stock
	err = json.NewDecoder(response.Body).Decode(&stocks)
	if err != nil {
		log.Fatal(err)
	}

	return stocks
}
