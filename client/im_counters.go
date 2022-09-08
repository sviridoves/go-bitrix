package client

import "github.com/sviridoves/go-bitrix/types"

func (c *Client) ImCountersGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.counters.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
