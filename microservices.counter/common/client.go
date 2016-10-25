package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type (
	Client interface {
		GetItemCountByTenant(uint32) (*GetItemCountByTenantResult, error)
		AddNewItemWithTenant(uint32, uint32) (*AddNewItemWithTenantResult, error)
	}
	ClientImpl struct {
		itemsToken string
	}
)

func NewClient(itemsToken string) Client {
	return &ClientImpl{itemsToken}
}

func (c *ClientImpl) AddNewItemWithTenant(itemId uint32, tenantId uint32) (*AddNewItemWithTenantResult, error) {
	var (
		res AddNewItemWithTenantResponse
	)
	uri := fmt.Sprintf("/svc/items/v2/%d/%d.json?api-key=%s",
		itemId,
		tenantId,
		c.itemsToken)

	rawRes, err := c.do(uri)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(rawRes, &res)
	return res.Result, err
}

func (c *ClientImpl) GetItemCountByTenant(tenantId uint32) (*GetItemCountByTenantResult, error) {
	var (
		res GetItemCountByTenantResponse
	)
	uri := fmt.Sprintf("/svc/items/v2/%d.json?api-key=%s",
		tenantId,
		c.itemsToken)

	rawRes, err := c.do(uri)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(rawRes, &res)
	return res.Result, err
}

func (c *ClientImpl) do(uri string) (body []byte, err error) {
	hc := http.Client{
		Timeout: 5 * time.Second,
	}
	req, err := http.NewRequest("GET", "https://api.nytimes.com"+uri, nil)
	if err != nil {
		return nil, err
	}

	var res *http.Response
	res, err = hc.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		if cerr := res.Body.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()
	return ioutil.ReadAll(res.Body)
}
