package exchangerates

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

func (c *Client) History(ctx context.Context, startAt string, endAt string, symbols ...string) (*HistoryResponse, *http.Response, error) {
	u := fmt.Sprintf("history?base=%s&start_at=%s&end_at=%s", c.baseCurrency, startAt, endAt)

	if len(symbols) != 0 {
		u = u + fmt.Sprintf("&symbols=%s", strings.Join(symbols, ","))
	}

	req, err := c.newRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	response := new(HistoryResponse)

	resp, err := c.do(ctx, req, response)
	if err != nil {
		return nil, resp, err
	}

	return response, resp, err
}
