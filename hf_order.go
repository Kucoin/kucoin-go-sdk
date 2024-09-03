package kucoin

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

// HfPlaceOrder There are two types of orders:
// (limit) order: set price and quantity for the transaction.
// (market) order : set amount or quantity for the transaction.
func (as *ApiService) HfPlaceOrder(ctx context.Context, params map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/hf/orders", params)
	return as.Call(ctx, req)
}

// HfSyncPlaceOrder The difference between this interface
// and "Place hf order" is that this interface will synchronously
// return the order information after the order matching is completed.
// For higher latency requirements, please select the "Place hf order" interface.
// If there is a requirement for returning data integrity, please select this interface
func (as *ApiService) HfSyncPlaceOrder(ctx context.Context, params map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/hf/orders/sync", params)
	return as.Call(ctx, req)
}

// HfPlaceMultiOrders This endpoint supports sequential batch order placement from a single endpoint.
// A maximum of 5orders can be placed simultaneously.
// The order types must be limit orders of the same trading pair
// （this endpoint currently only supports spot trading and does not support margin trading）
func (as *ApiService) HfPlaceMultiOrders(ctx context.Context, orders []*HFCreateMultiOrderModel) (*ApiResponse, error) {
	p := map[string]interface{}{
		"orderList": orders,
	}
	req := NewRequest(http.MethodPost, "/api/v1/hf/orders/multi", p)
	return as.Call(ctx, req)
}

// HfSyncPlaceMultiOrders The request parameters of this interface
// are the same as those of the "Sync place multiple hf orders" interface
// The difference between this interface and "Sync place multiple hf orders" is that
// this interface will synchronously return the order information after the order matching is completed.
func (as *ApiService) HfSyncPlaceMultiOrders(ctx context.Context, orders []*HFCreateMultiOrderModel) (*ApiResponse, error) {
	p := map[string]interface{}{
		"orderList": orders,
	}
	req := NewRequest(http.MethodPost, "/api/v1/hf/orders/multi/sync", p)
	return as.Call(ctx, req)
}

type HfSyncPlaceMultiOrdersRes []*HfSyncPlaceOrderRes

// HfModifyOrder
// This interface can modify the price and quantity of the order according to orderId or clientOid.
func (as *ApiService) HfModifyOrder(ctx context.Context, params map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/hf/orders/alter", params)
	return as.Call(ctx, req)
}

// HfCancelOrder This endpoint can be used to cancel a high-frequency order by orderId.
func (as *ApiService) HfCancelOrder(ctx context.Context, orderId, symbol string) (*ApiResponse, error) {
	p := map[string]string{
		"symbol": symbol,
	}
	req := NewRequest(http.MethodDelete, "/api/v1/hf/orders/"+orderId, p)
	return as.Call(ctx, req)
}

// HfSyncCancelOrder The difference between this interface and "Cancel orders by orderId" is that
// this interface will synchronously return the order information after the order canceling is completed.
func (as *ApiService) HfSyncCancelOrder(ctx context.Context, orderId, symbol string) (*ApiResponse, error) {
	p := map[string]string{
		"symbol": symbol,
	}
	req := NewRequest(http.MethodDelete, "/api/v1/hf/orders/sync/"+orderId, p)
	return as.Call(ctx, req)
}

// HfCancelOrderByClientId This endpoint sends out a request to cancel a high-frequency order using clientOid.
func (as *ApiService) HfCancelOrderByClientId(ctx context.Context, clientOid, symbol string) (*ApiResponse, error) {
	p := map[string]string{
		"symbol": symbol,
	}
	req := NewRequest(http.MethodDelete, "/api/v1/hf/orders/client-order/"+clientOid, p)
	return as.Call(ctx, req)
}

// HfSyncCancelOrderByClientId The difference between this interface and "Cancellation of order by clientOid"
// is that this interface will synchronously return the order information after the order canceling is completed.
func (as *ApiService) HfSyncCancelOrderByClientId(ctx context.Context, clientOid, symbol string) (*ApiResponse, error) {
	p := map[string]string{
		"symbol": symbol,
	}
	req := NewRequest(http.MethodDelete, "/api/v1/hf/orders/sync/client-order/"+clientOid, p)
	return as.Call(ctx, req)
}

// HfSyncCancelOrderWithSize This interface can cancel the specified quantity of the order according to the orderId.
func (as *ApiService) HfSyncCancelOrderWithSize(ctx context.Context, orderId, symbol, cancelSize string) (*ApiResponse, error) {
	p := map[string]string{
		"symbol":     symbol,
		"cancelSize": cancelSize,
	}
	req := NewRequest(http.MethodDelete, "/api/v1/hf/orders/cancel/"+orderId, p)
	return as.Call(ctx, req)
}

// HfSyncCancelAllOrders his endpoint allows cancellation of all orders related to a specific trading pair
// with a status of open
// (including all orders pertaining to high-frequency trading accounts and non-high-frequency trading accounts)
func (as *ApiService) HfSyncCancelAllOrders(ctx context.Context, symbol string) (*ApiResponse, error) {
	p := map[string]string{
		"symbol": symbol,
	}
	req := NewRequest(http.MethodDelete, "/api/v1/hf/orders", p)
	return as.Call(ctx, req)
}

// HfObtainActiveOrders This endpoint obtains a list of all active HF orders.
// The return data is sorted in descending order based on the latest update times.
func (as *ApiService) HfObtainActiveOrders(ctx context.Context, symbol string) (*ApiResponse, error) {
	p := map[string]string{
		"symbol": symbol,
	}
	req := NewRequest(http.MethodGet, "/api/v1/hf/orders/active", p)
	return as.Call(ctx, req)
}

// HfObtainActiveSymbols This interface can query all trading pairs that the user has active orders
func (as *ApiService) HfObtainActiveSymbols(ctx context.Context) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/hf/orders/active/symbols", nil)
	return as.Call(ctx, req)
}

type HfSymbolsModel struct {
	Symbols []string `json:"symbols"`
}

// HfObtainFilledOrders This endpoint obtains a list of filled HF orders and returns paginated data.
// The returned data is sorted in descending order based on the latest order update times.
func (as *ApiService) HfObtainFilledOrders(ctx context.Context, p map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/hf/orders/done", p)
	return as.Call(ctx, req)
}

type HfFilledOrdersModel struct {
	LastId json.Number     `json:"lastId"`
	Items  []*HfOrderModel `json:"items"`
}

// HfOrderDetail This endpoint can be used to obtain information for a single HF order using the order id.
func (as *ApiService) HfOrderDetail(ctx context.Context, orderId, symbol string) (*ApiResponse, error) {
	p := map[string]string{
		"symbol": symbol,
	}
	req := NewRequest(http.MethodGet, "/api/v1/hf/orders/"+orderId, p)
	return as.Call(ctx, req)
}

// HfOrderDetailByClientOid The endpoint can be used to obtain information about a single order using clientOid.
// If the order does not exist, then there will be a prompt saying that the order does not exist.
func (as *ApiService) HfOrderDetailByClientOid(ctx context.Context, clientOid, symbol string) (*ApiResponse, error) {
	p := map[string]string{
		"symbol": symbol,
	}
	req := NewRequest(http.MethodGet, "/api/v1/hf/orders/client-order/"+clientOid, p)
	return as.Call(ctx, req)
}

// HfAutoCancelSetting  automatically cancel all orders of the set trading pair after the specified time.
// If this interface is not called again for renewal or cancellation before the set time,
// the system will help the user to cancel the order of the corresponding trading pair.
// otherwise it will not.
func (as *ApiService) HfAutoCancelSetting(ctx context.Context, timeout int64, symbol string) (*ApiResponse, error) {
	p := map[string]interface{}{
		"symbol":  symbol,
		"timeout": timeout,
	}
	req := NewRequest(http.MethodPost, "/api/v1/hf/orders/dead-cancel-all", p)
	return as.Call(ctx, req)
}

// HfQueryAutoCancelSetting  Through this interface, you can query the settings of automatic order cancellation
func (as *ApiService) HfQueryAutoCancelSetting(ctx context.Context) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/hf/orders/dead-cancel-all/query", nil)
	return as.Call(ctx, req)
}

// HfTransactionDetails This endpoint can be used to obtain a list of the latest HF transaction details.
// The returned results are paginated. The data is sorted in descending order according to time.
func (as *ApiService) HfTransactionDetails(ctx context.Context, p map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/hf/fills", p)
	return as.Call(ctx, req)
}

// HfCancelOrders This endpoint can be used to cancel all hf orders. return HfCancelOrdersResultModel
func (as *ApiService) HfCancelOrders(ctx context.Context) (*ApiResponse, error) {
	req := NewRequest(http.MethodDelete, "/api/v1/hf/orders/cancelAll", nil)
	return as.Call(ctx, req)
}

func (as *ApiService) HfPlaceOrderTest(ctx context.Context, p *HfPlaceOrderReq) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/hf/orders/test", p)
	return as.Call(ctx, req)
}

func (as *ApiService) HfMarginActiveSymbols(ctx context.Context, tradeType string) (*ApiResponse, error) {
	p := map[string]string{
		"tradeType": tradeType,
	}
	req := NewRequest(http.MethodGet, "/api/v3/hf/margin/order/active/symbols", p)
	return as.Call(ctx, req)
}

// HfCreateMarinOrderV3 This interface is used to place cross-margin or isolated-margin high-frequency margin trading
func (as *ApiService) HfCreateMarinOrderV3(ctx context.Context, p *HfMarginOrderV3Req) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v3/hf/margin/order", p)
	return as.Call(ctx, req)
}

// HfCreateMarinOrderTestV3 Order test endpoint, the request parameters and return parameters of this endpoint are exactly the same as the order endpoint,
// and can be used to verify whether the signature is correct and other operations. After placing an order,
// the order will not enter the matching system, and the order cannot be queried.
func (as *ApiService) HfCreateMarinOrderTestV3(ctx context.Context, p *HfMarginOrderV3Req) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v3/hf/margin/order/test", p)
	return as.Call(ctx, req)
}

// HfCancelMarinOrderV3 Cancel a single order by orderId. If the order cannot be canceled (sold or canceled),
// an error message will be returned, and the reason can be obtained according to the returned msg.
func (as *ApiService) HfCancelMarinOrderV3(ctx context.Context, p *HfCancelMarinOrderV3Req) (*ApiResponse, error) {
	req := NewRequest(http.MethodDelete, fmt.Sprintf("/api/v3/hf/margin/orders/%s?symbol=%s", p.OrderId, p.Symbol), nil)
	return as.Call(ctx, req)
}

// HfCancelClientMarinOrderV3 Cancel a single order by clientOid.
func (as *ApiService) HfCancelClientMarinOrderV3(ctx context.Context, p *HfCancelClientMarinOrderV3Req) (*ApiResponse, error) {
	req := NewRequest(http.MethodDelete, fmt.Sprintf("/api/v3/hf/margin/orders/client-order/%s?symbol=%s", p.ClientOid, p.Symbol), nil)
	return as.Call(ctx, req)
}

// HfCancelAllMarginOrdersV3 This endpoint only sends cancellation requests.
// The results of the requests must be obtained by checking the order detail or subscribing to websocket.
func (as *ApiService) HfCancelAllMarginOrdersV3(ctx context.Context, p *HfCancelAllMarginOrdersV3Req) (*ApiResponse, error) {
	v, err := query.Values(p)
	if err != nil {
		return nil, err
	}
	req := NewRequest(http.MethodDelete, "/api/v3/hf/margin/orders", v)
	return as.Call(ctx, req)
}

// HfMarinActiveOrdersV3 This interface is to obtain all active hf margin order lists,
// and the return value of the active order interface is the paged data of all uncompleted order lists.
func (as *ApiService) HfMarinActiveOrdersV3(ctx context.Context, p *HfMarinActiveOrdersV3Req) (*ApiResponse, error) {
	v, err := query.Values(p)
	if err != nil {
		return nil, err
	}
	req := NewRequest(http.MethodGet, "/api/v3/hf/margin/orders/active", v)
	return as.Call(ctx, req)
}

// HfMarinDoneOrdersV3 This endpoint obtains a list of filled margin HF orders and returns paginated data.
// The returned data is sorted in descending order based on the latest order update times.
func (as *ApiService) HfMarinDoneOrdersV3(ctx context.Context, p *HfMarinDoneOrdersV3Req) (*ApiResponse, error) {
	v, err := query.Values(p)
	if err != nil {
		return nil, err
	}
	req := NewRequest(http.MethodGet, "/api/v3/hf/margin/orders/done", v)
	return as.Call(ctx, req)
}

// HfMarinOrderV3 This endpoint can be used to obtain information for a single margin HF order using the order id.
func (as *ApiService) HfMarinOrderV3(ctx context.Context, p *HfMarinOrderV3Req) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, fmt.Sprintf("/api/v3/hf/margin/orders/%s?symbol=%s", p.OrderId, p.Symbol), nil)
	return as.Call(ctx, req)
}

// HfMarinClientOrderV3 This endpoint can be used to obtain information for a single margin HF order using the clientOid.
func (as *ApiService) HfMarinClientOrderV3(ctx context.Context, p *HfMarinClientOrderV3Req) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, fmt.Sprintf("/api/v3/hf/margin/orders/client-order/%s?symbol=%s", p.ClientOid, p.Symbol), nil)
	return as.Call(ctx, req)
}

// HfMarinFillsV3 This endpoint can be used to obtain a list of the latest margin HF transaction details.
// The returned results are paginated.
// The data is sorted in descending order according to time.
func (as *ApiService) HfMarinFillsV3(ctx context.Context, p *HfMarinFillsV3Req) (*ApiResponse, error) {
	v, err := query.Values(p)
	if err != nil {
		return nil, err
	}
	req := NewRequest(http.MethodGet, "/api/v3/hf/margin/fills", v)
	return as.Call(ctx, req)
}
