package bsdex

import (
	"encoding/json"
	"strconv"
)

const DEPOSITS_ENDPOINT = "v2/crypto/deposits"

type Deposit struct {
	Uuid             string  `json:"uuid"`
	Amount           string  `json:"amount"`
	AssetId          string  `json:"asset_id"`
	CreatedAt        string  `json:"created_at"`
	FinalizedAt      string  `json:"finalized_at"`
	SourceAddress    string  `json:"source_address"`
	SourceTag        *string `json:"source_tag"`
	TargetAddress    string  `json:"target_address"`
	TargetTag        *string `json:"target_tag"`
	TransactionState string  `json:"transaction_state"`
	TransactionType  string  `json:"transaction_type"`
}

type DepositsResponse struct {
	Data []Deposit `json:"data"`
}

func (a *APIClient) Deposits(assetId string, createdAfter *string, page *int) (*DepositsResponse, error) {
	query := map[string]string{
		"asset_id": assetId,
	}

	if createdAfter != nil {
		query["created_after"] = *createdAfter
	}

	if page != nil {
		tmp := strconv.FormatInt(int64(*page), 10)
		query["page"] = tmp
	}

	b, err := a.requestGET(DEPOSITS_ENDPOINT, query)
	if err != nil {
		return nil, err
	}

	resp := new(DepositsResponse)
	err = json.Unmarshal(b, &resp)
	return resp, err
}
