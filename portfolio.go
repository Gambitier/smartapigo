package smartapigo

import (
	"net/http"
)

// Holding is an individual holdings response.
type Holding struct {
	Tradingsymbol      string  `json:"tradingsymbol"`
	Exchange           string  `json:"exchange"`
	ISIN               string  `json:"isin"`
	T1Quantity         int32   `json:"t1quantity"`
	RealisedQuantity   int32   `json:"realisedquantity"`
	Quantity           int32   `json:"quantity"`
	AuthorisedQuantity int32   `json:"authorisedquantity"`
	ProfitAndLoss      float32 `json:"profitandloss"`
	Product            string  `json:"product"`
	CollateralQuantity int32   `json:"collateralquantity"`
	CollateralType     string  `json:"collateraltype"`
	Haircut            float32 `json:"haircut"`
	AveragePrice       float32 `json:"averageprice"`
	LTP                float32 `json:"ltp"`
	SymbolToken        string  `json:"symboltoken"`
	Close              float32 `json:"close"`
	PNLPercentage      float32 `json:"pnlpercentage"`
}

// Holdings is a list of holdings
type Holdings []Holding

// GetHoldings gets a list of holdings.
func (c *Client) GetHoldings() (Holdings, error) {
	var holdings Holdings
	err := c.doEnvelope(http.MethodGet, URIGetHoldings, nil, nil, &holdings, true)
	return holdings, err
}
