package exchangerates_test

import (
	"context"
	"os"
	"testing"

	"github.com/investing-kr/go-exchangerates"
)

var c *exchangerates.Client

func TestMain(m *testing.M) {
	var err error
	c, err = exchangerates.NewClient(nil)
	if err != nil {
		os.Exit(1)
	}
	c.SetBaseCurrency("USD")
	os.Exit(m.Run())
}

func TestLatest(t *testing.T) {
	ctx := context.Background()
	rates, _, err := c.Latest(ctx, "USD", "JPY")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", *rates)
}

func TestHistory(t *testing.T) {
	ctx := context.Background()
	rates, _, err := c.History(ctx, "2020-01-01", "2020-01-03", "USD", "JPY")
	if err != nil {
		t.Fatal(err)
	}

	if (*rates).Rates["2020-01-02"].USD != float64(1) {
		t.Fatalf("expected 1, but got %v ", (*rates).Rates["2020-01-02"].USD)
	}
}
