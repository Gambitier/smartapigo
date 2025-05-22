package smartapigo

import "net/http"

type HistoricalDataRequest struct {
	Exchange    string `json:"exchange"`
	Symboltoken string `json:"symboltoken"`
	Interval    string `json:"interval"`
	Fromdate    string `json:"fromdate"`
	Todate      string `json:"todate"`
}

// GetHistoricalData gets historical candles data.
func (c *Client) GetHistoricalData(requestParams HistoricalDataRequest) ([][]interface{}, error) {
	var candles [][]interface{}

	params := structToMap(requestParams, "json")
	err := c.doEnvelope(http.MethodPost, URIHitorical, params, nil, &candles, true)

	if err != nil {
		return nil, err
	}
	return candles, err
}
