package bsdex

import (
	"encoding/json"
)

const QUOTE_ENDPOINT = "/quote"
const SYMBOL_SEPARATOR = "-"

type QuoteResp struct {
	BuyPrice   string `json:"buy_price"`
	BuyVolume  string `json:"buy_volume"`
	Market     string `json:"market"`
	SellPrice  string `json:"sell_price"`
	SellVolume string `json:"sell_volume"`
}

func (a *APIClient) Quote(baseAsset string, quoteAsset string) (*QuoteResp, error) {
	b, err := a.requestGET("v1/"+baseAsset+SYMBOL_SEPARATOR+quoteAsset+QUOTE_ENDPOINT, nil)
	if err != nil {
		return nil, err
	}

	resp := new(QuoteResp)
	err = json.Unmarshal(b, &resp)
	return resp, err
}
