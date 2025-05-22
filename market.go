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

// MarketOHLCData represents Market OHLC Data.
type MarketOHLCData struct {
	Exchange      string  `json:"exchange"`
	TradingSymbol string  `json:"tradingSymbol"`
	SymbolToken   string  `json:"symbolToken"`
	Ltp           float64 `json:"ltp"`
	Open          float64 `json:"open"`
	High          float64 `json:"high"`
	Low           float64 `json:"low"`
	Close         float64 `json:"close"`
}

// DepthData represents Market Depth Data.
type DepthData struct {
	Price    float64 `json:"price"`
	Quantity int64   `json:"quantity"`
	Orders   int64   `json:"orders"`
}

// Depth represents Market Depth Data.
type Depth struct {
	Buy  []DepthData `json:"buy"`
	Sell []DepthData `json:"sell"`
}

// MarketFullData represents Market Full Data.
type MarketFullData struct {
	Exchange         string  `json:"exchange"`
	TradingSymbol    string  `json:"tradingSymbol"`
	SymbolToken      string  `json:"symbolToken"`
	Ltp              float64 `json:"ltp"`
	Open             float64 `json:"open"`
	High             float64 `json:"high"`
	Low              float64 `json:"low"`
	Close            float64 `json:"close"`
	LastTradeQty     int64   `json:"lastTradeQty"`
	ExchFeedTime     string  `json:"exchFeedTime"`
	ExchTradeTime    string  `json:"exchTradeTime"`
	NetChange        float64 `json:"netChange"`
	PercentChange    float64 `json:"percentChange"`
	AvgPrice         float64 `json:"avgPrice"`
	TradeVolume      int64   `json:"tradeVolume"`
	OpnInterest      int64   `json:"opnInterest"`
	LowerCircuit     float64 `json:"lowerCircuit"`
	UpperCircuit     float64 `json:"upperCircuit"`
	TotalBuyQty      int64   `json:"totBuyQuan"`
	TotalSellQty     int64   `json:"totSellQuan"`
	FiftyTwoWeekLow  float64 `json:"52WeekLow"`
	FiftyTwoWeekHigh float64 `json:"52WeekHigh"`
	Depth            Depth   `json:"depth"`
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

// GetMarketLTPData gets Market LTP Data.
func (c *Client) GetMarketLTPData(exchangeTokens map[string][]string) (MarketDataResponse[MarketLTPData], error) {
	marketDataRequest := MarketDataRequest{
		Mode:           "LTP",
		ExchangeTokens: exchangeTokens,
	}

	var marketData MarketDataResponse[MarketLTPData]

	params := structToMap(marketDataRequest, "json")
	err := c.doEnvelope(http.MethodPost, URIMARKETDATA, params, nil, &marketData, true)

	return marketData, err
}

// GetMarketOHLCData gets Market OHLC Data.
func (c *Client) GetMarketOHLCData(exchangeTokens map[string][]string) (MarketDataResponse[MarketOHLCData], error) {
	marketDataRequest := MarketDataRequest{
		Mode:           "OHLC",
		ExchangeTokens: exchangeTokens,
	}

	var marketData MarketDataResponse[MarketOHLCData]

	params := structToMap(marketDataRequest, "json")
	err := c.doEnvelope(http.MethodPost, URIMARKETDATA, params, nil, &marketData, true)

	return marketData, err
}

// GetMarketFullData gets Market Full Data.
func (c *Client) GetMarketFullData(exchangeTokens map[string][]string) (MarketDataResponse[MarketFullData], error) {
	marketDataRequest := MarketDataRequest{
		Mode:           "FULL",
		ExchangeTokens: exchangeTokens,
	}

	var marketData MarketDataResponse[MarketFullData]

	params := structToMap(marketDataRequest, "json")
	err := c.doEnvelope(http.MethodPost, URIMARKETDATA, params, nil, &marketData, true)

	return marketData, err
}
