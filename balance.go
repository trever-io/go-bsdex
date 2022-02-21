package bsdex

import "encoding/json"

const BALANCE_ENDPOINT = "v1/balance"

type AssetBalance struct {
	AssetId   string `json:"asset_id"`
	Available string `json:"available"`
	Locked    string `json:"locked"`
}

type BalanceResponse []AssetBalance

func (a *APIClient) Balance() (BalanceResponse, error) {
	b, err := a.requestGET(BALANCE_ENDPOINT, nil)
	if err != nil {
		return nil, err
	}

	resp := make(BalanceResponse, 0)
	err = json.Unmarshal(b, &resp)
	return resp, err
}
