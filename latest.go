package exchangerates

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

func (c *Client) Latest(ctx context.Context, symbols ...string) (*LatestResponse, *http.Response, error) {
	u := fmt.Sprintf("latest?base=%s", c.baseCurrency)

	if len(symbols) != 0 {
		u = u + fmt.Sprintf("&symbols=%s", strings.Join(symbols, ","))
	}

	req, err := c.newRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	response := new(LatestResponse)

	resp, err := c.do(ctx, req, response)
	if err != nil {
		return nil, resp, err
	}

	return response, resp, err
}
