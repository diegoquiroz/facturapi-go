/*
Facturapi

Testing ReceiptAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_ReceiptAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ReceiptAPIService CancelReceipt", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var receiptId string

		resp, httpRes, err := apiClient.ReceiptAPI.CancelReceipt(context.Background(), receiptId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ReceiptAPIService CreateGlobalInvoice", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ReceiptAPI.CreateGlobalInvoice(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ReceiptAPIService CreateReceipt", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ReceiptAPI.CreateReceipt(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ReceiptAPIService DownloadReceiptPdf", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var receiptId string

		resp, httpRes, err := apiClient.ReceiptAPI.DownloadReceiptPdf(context.Background(), receiptId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ReceiptAPIService GetReceipt", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var receiptId string

		resp, httpRes, err := apiClient.ReceiptAPI.GetReceipt(context.Background(), receiptId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ReceiptAPIService InvoiceReceipt", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ReceiptAPI.InvoiceReceipt(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ReceiptAPIService ListReceipts", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ReceiptAPI.ListReceipts(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ReceiptAPIService SendReceiptByEmail", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var receiptId string

		resp, httpRes, err := apiClient.ReceiptAPI.SendReceiptByEmail(context.Background(), receiptId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
