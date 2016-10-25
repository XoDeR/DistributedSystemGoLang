package clientTest

import "microservices.counter/common"

type Client struct {
	TestGetMostPopular func(string, string, uint) ([]*common.MostPopularResult, error)
}

func (t *Client) GetMostPopular(resourceType, section string, timeframe uint) ([]*common.MostPopularResult, error) {
	return t.TestGetMostPopular(resourceType, section, timeframe)
}
