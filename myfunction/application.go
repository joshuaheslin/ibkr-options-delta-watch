package myfunction

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"
)

func sendMail(body string) {
	from := "joshua.heslin@gmail.com"
	pass := os.Getenv("GMAIL_PASSWORD")
	to := "jbmh@me.com"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Option Alert\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Println("Successfully sended to " + to)
}

type OptionSymbolQuote struct {
	Status          string    `json:"s"`
	OptionSymbol    []string  `json:"optionSymbol"`
	Underlying      []string  `json:"underlying"`
	UnderlyingPrice []float64 `json:"underlyingPrice"`
	Expiration      []int     `json:"expiration"`
	DTE             []int     `json:"dte"`
	Updated         []int     `json:"updated"`
	Mid             []float64 `json:"mid"`
	Delta           []float64 `json:"delta"`
}

func fetchSymbol(symbol string) (OptionSymbolQuote, error) {
	apiKey := os.Getenv("MARKET_DATA_API_KEY")
	URL := fmt.Sprintf("https://api.marketdata.app/v1/options/quotes/%s?format=json&token=%s", symbol, apiKey)

	// fmt.Println(URL)

	resp, err := http.Get(URL)
	if err != nil {
		fmt.Println(fmt.Errorf("ooopsss an error occurred, please try again, %s", err))
	}
	defer resp.Body.Close()

	var result OptionSymbolQuote

	//Decode the data
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Println(fmt.Errorf("ooopsss! an error occurred, please try again, %s", err))
	}

	return result, nil
}

func runSymbol(optionSymbol OptionSymbol) {
	result, err := fetchSymbol(optionSymbol.Symbol)
	if err != nil {
		fmt.Printf("could not fetch symbol %v", err)
	}

	if result.Status != "ok" {
		fmt.Println(optionSymbol.Symbol, result, err)
		return
	}

	underlyingPrice := result.UnderlyingPrice[0]
	dte := result.DTE[0]
	delta := result.Delta[0]
	underlying := result.Underlying[0]
	mid := result.Mid[0]

	message := ""
	sendAlert := false
	if delta < 0 {
		if delta < -0.35 {
			message = fmt.Sprintf("Delta is too high: %f", delta)
			sendAlert = true
		}
		if delta > -0.17 {
			message = fmt.Sprintf("Delta is winning: %f", delta)
			sendAlert = true
		}
	} else if delta > 0 {
		if delta > 0.35 {
			message = fmt.Sprintf("Delta is too high: %f", delta)
			sendAlert = true
		}
		if delta < 0.17 {
			message = fmt.Sprintf("Delta is winning: %f", delta)
			sendAlert = true
		}
	}

	optionDetail := fmt.Sprintf("%s %f dte:%d delta:%f mid:%f", underlying, underlyingPrice, dte, delta, mid)

	if !sendAlert {
		msg := fmt.Sprintf("%s Delta is ok %f. %s", optionSymbol.Symbol, delta, optionDetail)
		fmt.Println(msg)
		// sendMail(msg)
		return
	}

	emailMessage := fmt.Sprintf(`
Alert: 
%s

(Initial delta: %s)
(Initial Premium: %s)

Option: 
%s

Details:
%s
`, message, optionSymbol.InitialDelta, optionSymbol.InitialPremium, optionSymbol.Symbol, optionDetail)

	fmt.Println(emailMessage)
	sendMail(emailMessage)
}

type OptionSymbol struct {
	Symbol         string
	InitialDelta   string
	InitialPremium string
}

func Run() {
	fmt.Println("Fetching deltas...")

	// TODO: find a way to fetch from GSheets Spreadsheet tracker later.

	symbols := []OptionSymbol{{
		Symbol:         "SPY240126P00465000",
		InitialDelta:   "0.24",
		InitialPremium: "0.95",
	}, {
		Symbol:         "IWM240126P00195000",
		InitialDelta:   "0.25",
		InitialPremium: "1.26",
	}, {
		Symbol:         "SPY240216C00492000",
		InitialDelta:   "0.25",
		InitialPremium: "0.64",
	}, {
		Symbol:         "SPY240202C00490000",
		InitialDelta:   "0.23",
		InitialPremium: "0.96",
	}}

	for _, symbol := range symbols {
		runSymbol(symbol)
	}

	fmt.Println("Done!")
}
