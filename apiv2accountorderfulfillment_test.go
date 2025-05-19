// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dinariapisdk_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/dinaricrypto/dinari-api-sdk-go"
	"github.com/dinaricrypto/dinari-api-sdk-go/internal/testutil"
	"github.com/dinaricrypto/dinari-api-sdk-go/option"
)

func TestAPIV2AccountOrderFulfillmentQueryWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := dinariapisdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithSecret("My Secret"),
	)
	_, err := client.API.V2.Accounts.OrderFulfillments.Query(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		dinariapisdk.APIV2AccountOrderFulfillmentQueryParams{
			OrderIDs: []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
			Page:     dinariapisdk.Int(1),
			PageSize: dinariapisdk.Int(1),
		},
	)
	if err != nil {
		var apierr *dinariapisdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
