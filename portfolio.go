package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"time"
)

// Portfolio is a collection of stocks
type Portfolio struct {
	Stocks []Stock `json:"stocks"`
	Total  float64 `json:"total"`
}

// Stock is a single stock
type Stock struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Unit  int32   `json:"unit"`
	Date  string  `json:"date"`
}

// Add a new stock to the portfolio
func (p *Portfolio) Add(stock string, units int32) {
	var s Stock
	s.Name = stock
	s.Unit = units
	p.Stocks = append(p.Stocks, s)
}

// Profit : Return the profit or loss of the portfolio
func (p *Portfolio) Profit(start, end time.Time) (float64, float64) {
	p.load(start)
	totalA := p.Total
	p.load(end)
	totalB := p.Total

	profit := totalB - totalA
	days := float64(end.Sub(start).Hours() / 24)
	if debug {
		log.Println("Profit:", profit, "Cumulative Return: ", (profit / totalA), "Days: ", days)
	}
	ar := (math.Pow((1.0+(profit/totalA)), (365.0/days)) - 1.0) * 100

	return profit, ar
}

// GetPrice : Return the price of a stock
func (s *Stock) GetPrice(date time.Time) float64 {
	for _, stock := range stocks {
		if stock.Name == s.Name &&
			stock.Date == date.Format("2006-01-02") {
			return stock.Price
		}
	}

	return 0
}

func (p *Portfolio) load(date time.Time) {
	p.Total = 0

	for _, stock := range stocks {
		for i := range p.Stocks {
			if p.Stocks[i].Name == stock.Name &&
				stock.Date == date.Format("2006-01-02") {
				p.Stocks[i].Price = stock.Price
				p.Stocks[i].Date = stock.Date
				p.Total += (stock.Price * float64(p.Stocks[i].Unit))
				break
			}
		}
	}

	if debug {
		log.Printf("%+v", *p)
	}
}

// This is just to populate the portfolio with random stocks
func populate(s, e string) {
	rand.Seed(time.Now().UnixNano())
	var names []string
	names = append(names, "GOOG")
	names = append(names, "AAPL")
	names = append(names, "AMZN")
	names = append(names, "TSLA")
	names = append(names, "EBAY")
	names = append(names, "FB")

	start, _ := time.Parse(layout, s)
	end, _ := time.Parse(layout, e)

	var stocks []Stock
	inc := 20.0

	for start.Before(end) {
		for name := range names {
			var s Stock
			s.Name = names[name]
			s.Price = inc + (rand.Float64()*6 - 3)
			s.Date = start.Format("2006-01-02")
			stocks = append(stocks, s)
		}

		inc += 0.01

		start = start.AddDate(0, 0, 1)
	}

	jsonString, err := json.Marshal(stocks)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("stocks.json", jsonString, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
