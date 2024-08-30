package kucoin

import (
	"context"
	"encoding/json"
	"net/http"
)

type CreateEarnOrderReq struct {
	ProductId   string `json:"productId"`
	Amount      string `json:"amount"`
	AccountType string `json:"accountType"`
}

type CreateEarnOrderRes struct {
	OrderId   string `json:"orderId"`
	OrderTxId string `json:"orderTxId"`
}

// CreateEarnOrder subscribing to fixed income products. If the subscription fails, it returns the corresponding error code.
func (as *ApiService) CreateEarnOrder(ctx context.Context, createEarnOrderReq *CreateEarnOrderReq) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/earn/orders", createEarnOrderReq)
	return as.Call(ctx, req)
}

type DeleteEarnOrderRes struct {
	OrderTxId   string `json:"orderTxId"`
	DeliverTime int64  `json:"deliverTime"`
	Status      string `json:"status"`
	Amount      string `json:"amount"`
}

// DeleteEarnOrder initiating redemption by holding ID.
func (as *ApiService) DeleteEarnOrder(ctx context.Context, orderId, amount, fromAccountType, confirmPunishRedeem string) (*ApiResponse, error) {
	p := map[string]string{
		"orderId":             orderId,
		"fromAccountType":     fromAccountType,
		"confirmPunishRedeem": confirmPunishRedeem,
		"amount":              amount,
	}
	req := NewRequest(http.MethodDelete, "/api/v1/earn/orders", p)
	return as.Call(ctx, req)
}

type RedeemPreviewModel struct {
	Currency              string      `json:"currency"`
	RedeemAmount          json.Number `json:"redeemAmount"`
	PenaltyInterestAmount json.Number `json:"penaltyInterestAmount"`
	RedeemPeriod          int32       `json:"redeemPeriod"`
	DeliverTime           int64       `json:"deliverTime"`
	ManualRedeemable      bool        `json:"manualRedeemable"`
	RedeemAll             bool        `json:"redeemAll"`
}

// RedeemPreview retrieves redemption preview information by holding ID
func (as *ApiService) RedeemPreview(ctx context.Context, orderId, fromAccountType string) (*ApiResponse, error) {
	p := map[string]string{
		"orderId":         orderId,
		"fromAccountType": fromAccountType,
	}
	req := NewRequest(http.MethodGet, "/api/v1/earn/redeem-preview", p)
	return as.Call(ctx, req)
}

// QuerySavingProducts retrieves savings products.
func (as *ApiService) QuerySavingProducts(ctx context.Context, currency string) (*ApiResponse, error) {
	p := map[string]string{
		"currency": currency,
	}
	req := NewRequest(http.MethodGet, "/api/v1/earn/saving/products", p)
	return as.Call(ctx, req)
}

type HoldAssetsRes []*HoldAssetModel

type HoldAssetModel struct {
	OrderId              string      `json:"orderId"`
	ProductId            string      `json:"productId"`
	ProductCategory      string      `json:"productCategory"`
	ProductType          string      `json:"productType"`
	Currency             string      `json:"currency"`
	IncomeCurrency       string      `json:"incomeCurrency"`
	ReturnRate           json.Number `json:"returnRate"`
	HoldAmount           json.Number `json:"holdAmount"`
	RedeemedAmount       json.Number `json:"redeemedAmount"`
	RedeemingAmount      json.Number `json:"redeemingAmount"`
	LockStartTime        int64       `json:"lockStartTime"`
	LockEndTime          int64       `json:"lockEndTime"`
	PurchaseTime         int64       `json:"purchaseTime"`
	RedeemPeriod         int32       `json:"redeemPeriod"`
	Status               string      `json:"status"`
	EarlyRedeemSupported int32       `json:"earlyRedeemSupported"`
}

// QueryHoldAssets retrieves current holding assets of fixed income products
func (as *ApiService) QueryHoldAssets(ctx context.Context, productId, productCategory, currency string, pagination *PaginationParam) (*ApiResponse, error) {
	p := map[string]string{
		"productId":       productId,
		"productCategory": productCategory,
		"currency":        currency,
	}
	pagination.ReadParam(p)
	req := NewRequest(http.MethodGet, "/api/v1/earn/hold-assets", p)
	return as.Call(ctx, req)
}

type EarnProductsRes []*EarnProductModel

type EarnProductModel struct {
	Id                   string      `json:"id"`
	Currency             string      `json:"currency"`
	Category             string      `json:"category"`
	Type                 string      `json:"type"`
	Precision            int32       `json:"precision"`
	ProductUpperLimit    string      `json:"productUpperLimit"`
	UserUpperLimit       string      `json:"userUpperLimit"`
	UserLowerLimit       string      `json:"userLowerLimit"`
	RedeemPeriod         int         `json:"redeemPeriod"`
	LockStartTime        int64       `json:"lockStartTime"`
	LockEndTime          int64       `json:"lockEndTime"`
	ApplyStartTime       int64       `json:"applyStartTime"`
	ApplyEndTime         int64       `json:"applyEndTime"`
	ReturnRate           json.Number `json:"returnRate"`
	IncomeCurrency       string      `json:"incomeCurrency"`
	EarlyRedeemSupported int32       `json:"earlyRedeemSupported"`
	ProductRemainAmount  json.Number `json:"productRemainAmount"`
	Status               string      `json:"status"`
	RedeemType           string      `json:"redeemType"`
	IncomeReleaseType    string      `json:"incomeReleaseType"`
	InterestDate         int64       `json:"interestDate"`
	Duration             int32       `json:"duration"`
	NewUserOnly          int32       `json:"newUserOnly"`
}

// QueryPromotionProducts retrieves limited-time promotion products
func (as *ApiService) QueryPromotionProducts(ctx context.Context, currency string) (*ApiResponse, error) {
	p := map[string]string{
		"currency": currency,
	}
	req := NewRequest(http.MethodGet, "/api/v1/earn/promotion/products", p)
	return as.Call(ctx, req)
}

// QueryKCSStakingProducts retrieves KCS Staking products
func (as *ApiService) QueryKCSStakingProducts(ctx context.Context, currency string) (*ApiResponse, error) {
	p := map[string]string{
		"currency": currency,
	}
	req := NewRequest(http.MethodGet, "/api/v1/earn/kcs-staking/products", p)
	return as.Call(ctx, req)
}

// QueryStakingProducts retrieves  Staking products
func (as *ApiService) QueryStakingProducts(ctx context.Context, currency string) (*ApiResponse, error) {
	p := map[string]string{
		"currency": currency,
	}
	req := NewRequest(http.MethodGet, "/api/v1/earn/staking/products", p)
	return as.Call(ctx, req)
}

// QueryETHProducts retrieves ETH  Staking products
func (as *ApiService) QueryETHProducts(ctx context.Context, currency string) (*ApiResponse, error) {
	p := map[string]string{
		"currency": currency,
	}
	req := NewRequest(http.MethodGet, "/api/v1/earn/eth-staking/products", p)
	return as.Call(ctx, req)
}

type OTCLoanModel struct {
	ParentUid string `json:"parentUid"`
	Orders    []struct {
		OrderId   string      `json:"orderId"`
		Currency  string      `json:"currency"`
		Principal json.Number `json:"principal"`
		Interest  json.Number `json:"interest"`
	} `json:"orders"`
	Ltv struct {
		TransferLtv           json.Number `json:"transferLtv"`
		OnlyClosePosLtv       json.Number `json:"onlyClosePosLtv"`
		DelayedLiquidationLtv json.Number `json:"delayedLiquidationLtv"`
		InstantLiquidationLtv json.Number `json:"instantLiquidationLtv"`
		CurrentLtv            json.Number `json:"currentLtv"`
	} `json:"ltv"`
	TotalMarginAmount    json.Number `json:"totalMarginAmount"`
	TransferMarginAmount json.Number `json:"transferMarginAmount"`
	Margins              []struct {
		MarginCcy    string      `json:"marginCcy"`
		MarginQty    json.Number `json:"marginQty"`
		MarginFactor json.Number `json:"marginFactor"`
	} `json:"margins"`
}

// QueryOTCLoanInfo querying accounts that are currently involved in loans.
func (as *ApiService) QueryOTCLoanInfo(ctx context.Context) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/otc-loan/loan", nil)
	return as.Call(ctx, req)
}

type OTCAccountsModel []*OTCAccountModel

type OTCAccountModel struct {
	Uid          string `json:"uid"`
	MarginCcy    string `json:"marginCcy"`
	MarginQty    string `json:"marginQty"`
	MarginFactor string `json:"marginFactor"`
	AccountType  string `json:"accountType"`
	IsParent     bool   `json:"isParent"`
}

// QueryOTCLoanAccountsInfo querying accounts that are currently involved in off-exchange funding and loans.
func (as *ApiService) QueryOTCLoanAccountsInfo(ctx context.Context) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/otc-loan/accounts", nil)
	return as.Call(ctx, req)
}
