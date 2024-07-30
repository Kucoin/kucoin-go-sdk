package kucoin

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"net/http"
)

// HfPlaceOrder There are two types of orders:
// (limit) order: set price and quantity for the transaction.
// (market) order : set amount or quantity for the transaction.
func (as *ApiService) HfPlaceOrder(params map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/hf/orders", params)
	return as.Call(req)
}

// HfSyncPlaceOrder The difference between this interface
// and "Place hf order" is that this interface will synchronously
// return the order information after the order matching is completed.
// For higher latency requirements, please select the "Place hf order" interface.
// If there is a requirement for returning data integrity, please select this interface
func (as *ApiService) HfSyncPlaceOrder(params map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/hf/orders/sync", params)
	return as.Call(req)
}

// HfPlaceMultiOrders This endpoint supports sequential batch order placement from a single endpoint.
// A maximum of 5orders can be placed simultaneously.
// The order types must be limit orders of the same trading pair
// （this endpoint currently only supports spot trading and does not support margin trading）
func (as *ApiService) HfPlaceMultiOrders(orders []*HFCreateMultiOrderModel) (*ApiResponse, error) {
	p := map[string]interface{}{
		"orderList": orders,
	}
	req := NewRequest(http.MethodPost, "/api/v1/hf/orders/multi", p)
	return as.Call(req)
}

// HfSyncPlaceMultiOrders The request parameters of this interface
// are the same as those of the "Sync place multiple hf orders" interface
// The difference between this interface and "Sync place multiple hf orders" is that
// this interface will synchronously return the order information after the order matching is completed.
func (as *ApiService) HfSyncPlaceMultiOrders(orders []*HFCreateMultiOrderModel) (*ApiResponse, error) {
	p := map[string]interface{}{
		"orderList": orders,
	}
	req := NewRequest(http.MethodPost, "/api/v1/hf/orders/multi/sync", p)
	return as.Call(req)
}

type HfSyncPlaceMultiOrdersRes []*HfSyncPlaceOrderRes

// HfModifyOrder
// This interface can modify the price and quantity of the order according to orderId or clientOid.
func (as *ApiService) HfModifyOrder(params map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/hf/orders/alter", params)
	return as.Call(req)
}

// HfCancelOrder This endpoint can be used to cancel a high-frequency order by orderId.
func (as *ApiService) HfCancelOrder(orderId, symbol string) (*ApiResponse, error) {
	p := map[string]string{
		"symbol": symbol,
	}
	req := NewRequest(http.MethodDelete, "/api/v1/hf/orders/"+orderId, p)
	return as.Call(req)
}

// HfSyncCancelOrder The difference between this interface and "Cancel orders by orderId" is that
// this interface will synchronously return the order information after the order canceling is completed.
func (as *ApiService) HfSyncCancelOrder(orderId, symbol string) (*ApiResponse, error) {
	p := map[string]string{
		"symbol": symbol,
	}
	req := NewRequest(http.MethodDelete, "/api/v1/hf/orders/sync/"+orderId, p)
	return as.Call(req)
}

// HfCancelOrderByClientId This endpoint sends out a request to cancel a high-frequency order using clientOid.
func (as *ApiService) HfCancelOrderByClientId(clientOid, symbol string) (*ApiResponse, error) {
	p := map[string]string{
		"symbol": symbol,
	}
	req := NewRequest(http.MethodDelete, "/api/v1/hf/orders/client-order/"+clientOid, p)
	return as.Call(req)
}

// HfSyncCancelOrderByClientId The difference between this interface and "Cancellation of order by clientOid"
// is that this interface will synchronously return the order information after the order canceling is completed.
func (as *ApiService) HfSyncCancelOrderByClientId(clientOid, symbol string) (*ApiResponse, error) {
	p := map[string]string{
		"symbol": symbol,
	}
	req := NewRequest(http.MethodDelete, "/api/v1/hf/orders/sync/client-order/"+clientOid, p)
	return as.Call(req)
}

// HfSyncCancelOrderWithSize This interface can cancel the specified quantity of the order according to the orderId.
func (as *ApiService) HfSyncCancelOrderWithSize(orderId, symbol, cancelSize string) (*ApiResponse, error) {
	p := map[string]string{
		"symbol":     symbol,
		"cancelSize": cancelSize,
	}
	req := NewRequest(http.MethodDelete, "/api/v1/hf/orders/cancel/"+orderId, p)
	return as.Call(req)
}

// HfSyncCancelAllOrders his endpoint allows cancellation of all orders related to a specific trading pair
// with a status of open
// (including all orders pertaining to high-frequency trading accounts and non-high-frequency trading accounts)
func (as *ApiService) HfSyncCancelAllOrders(symbol string) (*ApiResponse, error) {
	p := map[string]string{
		"symbol": symbol,
	}
	req := NewRequest(http.MethodDelete, "/api/v1/hf/orders", p)
	return as.Call(req)
}

// HfObtainActiveOrders This endpoint obtains a list of all active HF orders.
// The return data is sorted in descending order based on the latest update times.
func (as *ApiService) HfObtainActiveOrders(symbol string) (*ApiResponse, error) {
	p := map[string]string{
		"symbol": symbol,
	}
	req := NewRequest(http.MethodGet, "/api/v1/hf/orders/active", p)
	return as.Call(req)
}

// HfObtainActiveSymbols This interface can query all trading pairs that the user has active orders
func (as *ApiService) HfObtainActiveSymbols() (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/hf/orders/active/symbols", nil)
	return as.Call(req)
}

type HfSymbolsModel struct {
	Symbols []string `json:"symbols"`
}

// HfObtainFilledOrders This endpoint obtains a list of filled HF orders and returns paginated data.
// The returned data is sorted in descending order based on the latest order update times.
func (as *ApiService) HfObtainFilledOrders(p map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/hf/orders/done", p)
	return as.Call(req)
}

type HfFilledOrdersModel struct {
	LastId json.Number     `json:"lastId"`
	Items  []*HfOrderModel `json:"items"`
}

// HfOrderDetail This endpoint can be used to obtain information for a single HF order using the order id.
func (as *ApiService) HfOrderDetail(orderId, symbol string) (*ApiResponse, error) {
	p := map[string]string{
		"symbol": symbol,
	}
	req := NewRequest(http.MethodGet, "/api/v1/hf/orders/"+orderId, p)
	return as.Call(req)
}

// HfOrderDetailByClientOid The endpoint can be used to obtain information about a single order using clientOid.
// If the order does not exist, then there will be a prompt saying that the order does not exist.
func (as *ApiService) HfOrderDetailByClientOid(clientOid, symbol string) (*ApiResponse, error) {
	p := map[string]string{
		"symbol": symbol,
	}
	req := NewRequest(http.MethodGet, "/api/v1/hf/orders/client-order/"+clientOid, p)
	return as.Call(req)
}

// HfAutoCancelSetting  automatically cancel all orders of the set trading pair after the specified time.
// If this interface is not called again for renewal or cancellation before the set time,
// the system will help the user to cancel the order of the corresponding trading pair.
// otherwise it will not.
func (as *ApiService) HfAutoCancelSetting(timeout int64, symbol string) (*ApiResponse, error) {
	p := map[string]interface{}{
		"symbol":  symbol,
		"timeout": timeout,
	}
	req := NewRequest(http.MethodPost, "/api/v1/hf/orders/dead-cancel-all", p)
	return as.Call(req)
}

// HfQueryAutoCancelSetting  Through this interface, you can query the settings of automatic order cancellation
func (as *ApiService) HfQueryAutoCancelSetting() (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/hf/orders/dead-cancel-all/query", nil)
	return as.Call(req)
}

// HfTransactionDetails This endpoint can be used to obtain a list of the latest HF transaction details.
// The returned results are paginated. The data is sorted in descending order according to time.
func (as *ApiService) HfTransactionDetails(p map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/hf/fills", p)
	return as.Call(req)
}

// HfCancelOrders This endpoint can be used to cancel all hf orders. return HfCancelOrdersResultModel
func (as *ApiService) HfCancelOrders() (*ApiResponse, error) {
	req := NewRequest(http.MethodDelete, "/api/v1/hf/orders/cancelAll", nil)
	return as.Call(req)
}

func (as *ApiService) HfPlaceOrderTest(p *HfPlaceOrderReq) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/hf/orders/test", p)
	return as.Call(req)
}

func (as *ApiService) HfMarginActiveSymbols(tradeType string) (*ApiResponse, error) {
	p := map[string]string{
		"tradeType": tradeType,
	}
	req := NewRequest(http.MethodGet, "/api/v3/hf/margin/order/active/symbols", p)
	return as.Call(req)
}

// HfCreateMarinOrderV3 This interface is used to place cross-margin or isolated-margin high-frequency margin trading
func (as *ApiService) HfCreateMarinOrderV3(p *HfMarginOrderV3Req) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v3/hf/margin/order", p)
	return as.Call(req)
}

// HfCreateMarinOrderTestV3 Order test endpoint, the request parameters and return parameters of this endpoint are exactly the same as the order endpoint,
// and can be used to verify whether the signature is correct and other operations. After placing an order,
// the order will not enter the matching system, and the order cannot be queried.
func (as *ApiService) HfCreateMarinOrderTestV3(p *HfMarginOrderV3Req) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v3/hf/margin/order/test", p)
	return as.Call(req)
}

// HfCancelMarinOrderV3 Cancel a single order by orderId. If the order cannot be canceled (sold or canceled),
// an error message will be returned, and the reason can be obtained according to the returned msg.
func (as *ApiService) HfCancelMarinOrderV3(p *HfCancelMarinOrderV3Req) (*ApiResponse, error) {
	req := NewRequest(http.MethodDelete, fmt.Sprintf("/api/v3/hf/margin/orders/%s?symbol=%s", p.OrderId, p.Symbol), nil)
	return as.Call(req)
}

// HfCancelClientMarinOrderV3 Cancel a single order by clientOid.
func (as *ApiService) HfCancelClientMarinOrderV3(p *HfCancelClientMarinOrderV3Req) (*ApiResponse, error) {
	req := NewRequest(http.MethodDelete, fmt.Sprintf("/api/v3/hf/margin/orders/client-order/%s?symbol=%s", p.ClientOid, p.Symbol), nil)
	return as.Call(req)
}

// HfCancelAllMarginOrdersV3 This endpoint only sends cancellation requests.
// The results of the requests must be obtained by checking the order detail or subscribing to websocket.
func (as *ApiService) HfCancelAllMarginOrdersV3(p *HfCancelAllMarginOrdersV3Req) (*ApiResponse, error) {
	v, err := query.Values(p)
	if err != nil {
		return nil, err
	}
	req := NewRequest(http.MethodDelete, "/api/v3/hf/margin/orders", v)
	return as.Call(req)
}

// HfMarinActiveOrdersV3 This interface is to obtain all active hf margin order lists,
// and the return value of the active order interface is the paged data of all uncompleted order lists.
func (as *ApiService) HfMarinActiveOrdersV3(p *HfMarinActiveOrdersV3Req) (*ApiResponse, error) {
	v, err := query.Values(p)
	if err != nil {
		return nil, err
	}
	req := NewRequest(http.MethodGet, "/api/v3/hf/margin/orders/active", v)
	return as.Call(req)
}

// HfMarinDoneOrdersV3 This endpoint obtains a list of filled margin HF orders and returns paginated data.
// The returned data is sorted in descending order based on the latest order update times.
func (as *ApiService) HfMarinDoneOrdersV3(p *HfMarinDoneOrdersV3Req) (*ApiResponse, error) {
	v, err := query.Values(p)
	if err != nil {
		return nil, err
	}
	req := NewRequest(http.MethodGet, "/api/v3/hf/margin/orders/done", v)
	return as.Call(req)
}

// HfMarinOrderV3 This endpoint can be used to obtain information for a single margin HF order using the order id.
func (as *ApiService) HfMarinOrderV3(p *HfMarinOrderV3Req) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, fmt.Sprintf("/api/v3/hf/margin/orders/%s?symbol=%s", p.OrderId, p.Symbol), nil)
	return as.Call(req)
}

// HfMarinClientOrderV3 This endpoint can be used to obtain information for a single margin HF order using the clientOid.
func (as *ApiService) HfMarinClientOrderV3(p *HfMarinClientOrderV3Req) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, fmt.Sprintf("/api/v3/hf/margin/orders/client-order/%s?symbol=%s", p.ClientOid, p.Symbol), nil)
	return as.Call(req)
}

// HfMarinFillsV3 This endpoint can be used to obtain a list of the latest margin HF transaction details.
// The returned results are paginated.
// The data is sorted in descending order according to time.
func (as *ApiService) HfMarinFillsV3(p *HfMarinFillsV3Req) (*ApiResponse, error) {
	v, err := query.Values(p)
	if err != nil {
		return nil, err
	}
	req := NewRequest(http.MethodGet, "/api/v3/hf/margin/fills", v)
	return as.Call(req)
}
