package exchangerates

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	defaultBaseURL      = "https://api.exchangeratesapi.io"
	defaultBaseCurrency = "USD"
)

type service struct {
	client *Client
}

type Client struct {
	httpClient   *http.Client
	baseURL      *url.URL
	baseCurrency string
}

func (c *Client) SetBaseCurrency(currency string) {
	c.baseCurrency = currency
}

type ClientOptions struct {
	APIURL       string
	BaseCurrency string
}

func NewClient(httpClient *http.Client, clientOpts ...*ClientOptions) (*Client, error) {
	var opt *ClientOptions
	switch len(clientOpts) {
	case 1:
	case 0:
		opt = &ClientOptions{
			APIURL:       defaultBaseURL,
			BaseCurrency: defaultBaseCurrency,
		}
	default:
		return nil, errors.New("exchangerates: invalid client options")
	}

	apiURL := defaultBaseURL
	if opt != nil && opt.APIURL != "" {
		apiURL = opt.APIURL
	}

	baseURL, err := url.Parse(apiURL)
	if err != nil {
		return nil, err
	}

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseCurrency := defaultBaseCurrency
	if opt != nil && opt.BaseCurrency != "" {
		baseCurrency = opt.BaseCurrency
	}

	c := &Client{
		baseURL:      baseURL,
		httpClient:   httpClient,
		baseCurrency: baseCurrency,
	}
	return c, nil
}

type ErrResponse struct {
	ErrorMsg  string `json:"error"`
	Exception string `json:"exception"`
}

func (e *ErrResponse) Error() string {
	if e.Exception != "" {
		return fmt.Sprintf("exchangerates: %s, %s", e.ErrorMsg, e.Exception)
	}
	return fmt.Sprintf("exchangerates: %s", e.ErrorMsg)
}

func checkResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	errResp := &ErrResponse{}
	err = json.Unmarshal(body, errResp)

	// If server response with sturctured error
	if err == nil && errResp.ErrorMsg != "" {
		return errResp
	}

	return errors.New(string(body))
}

func (c *Client) newRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	u, err := c.baseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	return req, nil
}

func (c *Client) do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	if ctx == nil {
		ctx = context.TODO()
	}

	req = req.WithContext(ctx)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		return nil, err
	}
	defer resp.Body.Close()

	err = checkResponse(resp)
	if err != nil {
		return resp, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			decErr := json.NewDecoder(resp.Body).Decode(v)
			if decErr == io.EOF {
				decErr = nil
			}
			if decErr != nil {
				err = decErr
			}
		}
	}

	return resp, err
}
