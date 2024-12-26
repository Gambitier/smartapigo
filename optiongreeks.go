package smartapigo

import (
	"net/http"
)

type OptionGreeksRequest struct {
	Name       string `json:"name"`
	Expirydate string `json:"expirydate"`
}

type OptionStrikeGreeks struct {
	Name              string `json:"name"`
	Expirydate        string `json:"expiry"`
	StrikePrice       string `json:"strikePrice"`
	OptionType        string `json:"optionType"`
	Delta             string `json:"delta"`
	Gamma             string `json:"gamma"`
	Theta             string `json:"theta"`
	Vega              string `json:"vega"`
	ImpliedVolatility string `json:"impliedVolatility"`
	TradeVolume       string `json:"tradeVolume"`
}

// GetOptionGreeks gets Option Greeks.
func (c *Client) GetOptionGreeks(optionGreeksRequest OptionGreeksRequest) ([]OptionStrikeGreeks, error) {
	var optionStrikesGreeks []OptionStrikeGreeks
	params := structToMap(optionGreeksRequest, "json")
	err := c.doEnvelope(http.MethodPost, URIOPTIONGREEKS, params, nil, &optionStrikesGreeks, true)
	return optionStrikesGreeks, err
}
