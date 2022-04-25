package bsdex

import (
	"encoding/json"
)

const AUCTION_ENDPOINT = "/auction"
const SYMBOL_SEPARATOR = "-"

type AuctionResp struct {
	BuyPrice   string `json:"buy_price"`
	BuyVolume  string `json:"buy_volume"`
	Market     string `json:"market"`
	SellPrice  string `json:"sell_price"`
	SellVolume string `json:"sell_volume"`
}

func (a *APIClient) Auction(baseAsset string, quoteAsset string) (*AuctionResp, error) {
  b, err := a.requestGET(baseAsset + SYMBOL_SEPARATOR + quoteAsset + AUCTION_ENDPOINT, nil)
	if err != nil {
		return nil, err
	}

	resp := new(AuctionResp)
	err = json.Unmarshal(b, &resp)
	return resp, err
}
