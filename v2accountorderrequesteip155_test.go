// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dinariapisdkgo_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/dinaricrypto/dinari-api-sdk-go"
	"github.com/dinaricrypto/dinari-api-sdk-go/internal/testutil"
	"github.com/dinaricrypto/dinari-api-sdk-go/option"
)

func TestV2AccountOrderRequestEip155NewPermitWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := dinariapisdkgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKeyID("My API Key ID"),
		option.WithAPISecretKey("My API Secret Key"),
	)
	_, err := client.V2.Accounts.OrderRequests.Eip155.NewPermit(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		dinariapisdkgo.V2AccountOrderRequestEip155NewPermitParams{
			ChainID:              dinariapisdkgo.ChainEip155_1,
			OrderSide:            dinariapisdkgo.OrderSideBuy,
			OrderTif:             dinariapisdkgo.OrderTifDay,
			OrderType:            dinariapisdkgo.OrderTypeMarket,
			PaymentToken:         "payment_token",
			AssetTokenQuantity:   dinariapisdkgo.Float(0),
			ClientOrderID:        dinariapisdkgo.String("client_order_id"),
			LimitPrice:           dinariapisdkgo.Float(0),
			PaymentTokenQuantity: dinariapisdkgo.Float(0),
			StockID:              dinariapisdkgo.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			TokenID:              dinariapisdkgo.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		},
	)
	if err != nil {
		var apierr *dinariapisdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV2AccountOrderRequestEip155NewPermitTransaction(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := dinariapisdkgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKeyID("My API Key ID"),
		option.WithAPISecretKey("My API Secret Key"),
	)
	_, err := client.V2.Accounts.OrderRequests.Eip155.NewPermitTransaction(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		dinariapisdkgo.V2AccountOrderRequestEip155NewPermitTransactionParams{
			Eip155OrderRequestPermitTransaction: dinariapisdkgo.Eip155OrderRequestPermitTransactionParam{
				OrderRequestID:  "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
				PermitSignature: "0xeaF12bD1DfFd",
			},
		},
	)
	if err != nil {
		var apierr *dinariapisdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV2AccountOrderRequestEip155Submit(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := dinariapisdkgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKeyID("My API Key ID"),
		option.WithAPISecretKey("My API Secret Key"),
	)
	_, err := client.V2.Accounts.OrderRequests.Eip155.Submit(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		dinariapisdkgo.V2AccountOrderRequestEip155SubmitParams{
			Eip155OrderRequestPermitTransaction: dinariapisdkgo.Eip155OrderRequestPermitTransactionParam{
				OrderRequestID:  "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
				PermitSignature: "0xeaF12bD1DfFd",
			},
		},
	)
	if err != nil {
		var apierr *dinariapisdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
