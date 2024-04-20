package smartapigo

import (
	"net/http"
)

// LTPResponse represents LTP API Response.
type LTPResponse struct {
	Exchange      string  `json:"exchange"`
	TradingSymbol string  `json:"tradingsymbol"`
	SymbolToken   string  `json:"symboltoken"`
	Open          float64 `json:"open"`
	High          float64 `json:"high"`
	Low           float64 `json:"low"`
	Close         float64 `json:"close"`
	Ltp           float64 `json:"ltp"`
}

// LTPParams represents parameters for getting LTP.
type LTPParams struct {
	Exchange      string `json:"exchange"`
	TradingSymbol string `json:"tradingsymbol"`
	SymbolToken   string `json:"symboltoken"`
}

// GetLTP gets Last Traded Price.
func (c *Client) GetLTP(ltpParams LTPParams) (LTPResponse, error) {
	var ltp LTPResponse
	params := structToMap(ltpParams, "json")
	err := c.doEnvelope(http.MethodPost, URILTP, params, nil, &ltp, true)
	return ltp, err
}

// Market data modes
const (
	MARKET_DATA_FULL_MODE string = "FULL"
	MARKET_DATA_OHLC_MODE string = "OHLC"
	MARKET_DATA_LTP_MODE  string = "LTP"
)

// MarketDataRequest represents parameters for getting Market Data.
type MarketDataRequest struct {
	Mode           string              `json:"mode"`
	ExchangeTokens map[string][]string `json:"exchangeTokens"`
}

// MarketLTPData represents Market LTP Data.
type MarketLTPData struct {
	Exchange      string  `json:"exchange"`
	TradingSymbol string  `json:"tradingSymbol"`
	SymbolToken   string  `json:"symbolToken"`
	Ltp           float64 `json:"ltp"`
}

// MarketUnfetchedData represents Market Unfetched Data.
type MarketUnfetchedData struct {
	Exchange    string `json:"exchange"`
	SymbolToken string `json:"symbolToken"`
	Message     string `json:"message"`
	ErrorCode   string `json:"errorCode"`
}

// MarketDataResponse represents Market Data API Response Data.
type MarketDataResponse[ModeData any] struct {
	Fetched   []ModeData            `json:"fetched"`
	Unfetched []MarketUnfetchedData `json:"unfetched"`
}

// GetMarketData gets Market Data.
func (c *Client) GetMarketData(marketDataRequest MarketDataRequest) (interface{}, error) {
	var marketData MarketDataResponse[MarketLTPData]
	params := structToMap(marketDataRequest, "json")
	err := c.doEnvelope(http.MethodPost, URIMARKETDATA, params, nil, &marketData, true)
	return marketData, err
}
