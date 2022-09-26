package bsdex

import (
	"encoding/json"
	"fmt"
	"strconv"
)

const WITHDRAWALS_ENDPOINT = "v2/crypto/withdrawals"
const FIAT_WITHDRAWALS_ENDPOINT = "v2/fiat/withdrawals"

type Withdrawal struct {
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

type CreateWithdrawalRequest struct {
	Amount         string  `json:"amount"`
	AssetId        string  `json:"asset_id"`
	DestinationTag *string `json:"destination_tag"`
	TargetAddress  string  `json:"target_address"`
}

type CreateFiatWithdrawalRequest struct {
	Amount        string `json:"amount"`
	AssetId       string `json:"asset_id"`
	TargetAddress string `json:"target_address"`
}

type WithdrawalsResponse struct {
	Data []Withdrawal `json:"data"`
}

type WithdrawalResponse struct {
	Data Withdrawal `json:"data"`
}

type FiatWithdrawalsResponse struct {
	Data []Withdrawal `json:"data"`
}

type FiatWithdrawalResponse struct {
	Data Withdrawal `json:"data"`
}

func (a *APIClient) Withdrawals(assetId string, createdAfter *string, page *int) (*WithdrawalsResponse, error) {
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

	b, err := a.requestGET(WITHDRAWALS_ENDPOINT, query)
	if err != nil {
		return nil, err
	}

	resp := new(WithdrawalsResponse)
	err = json.Unmarshal(b, &resp)
	return resp, err
}

func (a *APIClient) FiatWithdrawals(assetId string, createdAfter *string, page *int) (*FiatWithdrawalsResponse, error) {
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

	b, err := a.requestGET(FIAT_WITHDRAWALS_ENDPOINT, query)
	if err != nil {
		return nil, err
	}

	resp := new(FiatWithdrawalsResponse)
	err = json.Unmarshal(b, &resp)
	return resp, err
}

func (a *APIClient) CreateWithdrawal(assetId string, amount string, targetAddress string, destinationTag *string) (*WithdrawalResponse, error) {
	req := &CreateWithdrawalRequest{
		AssetId:        assetId,
		Amount:         amount,
		TargetAddress:  targetAddress,
		DestinationTag: destinationTag,
	}

	b, err := a.requestPOST(WITHDRAWALS_ENDPOINT, req)
	if err != nil {
		return nil, err
	}

	resp := new(WithdrawalResponse)
	err = json.Unmarshal(b, &resp)
	return resp, err
}

func (a *APIClient) CreateFiatWithdrawal(assetId string, amount string, targetAddress string) (*FiatWithdrawalResponse, error) {
	req := &CreateFiatWithdrawalRequest{
		AssetId:       assetId,
		Amount:        amount,
		TargetAddress: targetAddress,
	}

	b, err := a.requestPOST(FIAT_WITHDRAWALS_ENDPOINT, req)
	if err != nil {
		return nil, err
	}

	resp := new(FiatWithdrawalResponse)
	err = json.Unmarshal(b, &resp)
	return resp, err
}

func (a *APIClient) ExecuteWithdrawal(uuid string) (*WithdrawalResponse, error) {
	endpoint := fmt.Sprintf("%v/%v/execute", WITHDRAWALS_ENDPOINT, uuid)
	b, err := a.requestPOST(endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(WithdrawalResponse)
	err = json.Unmarshal(b, &resp)
	return resp, err
}

func (a *APIClient) ExecuteFiatWithdrawal(uuid string) (*FiatWithdrawalResponse, error) {
	endpoint := fmt.Sprintf("%v/%v/execute", FIAT_WITHDRAWALS_ENDPOINT, uuid)
	b, err := a.requestPOST(endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(FiatWithdrawalResponse)
	err = json.Unmarshal(b, &resp)
	return resp, err
}
