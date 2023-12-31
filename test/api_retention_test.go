/*
Facturapi

Testing RetentionAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package facturapi-go

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/diegoquiroz/facturapi-go"
)

func Test_facturapi-go_RetentionAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RetentionAPIService CancelRetention", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var retentionId string

		resp, httpRes, err := apiClient.RetentionAPI.CancelRetention(context.Background(), retentionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RetentionAPIService CreateRetention", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RetentionAPI.CreateRetention(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RetentionAPIService DownloadRetention", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var retentionId string
		var format string

		resp, httpRes, err := apiClient.RetentionAPI.DownloadRetention(context.Background(), retentionId, format).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RetentionAPIService GetRetention", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var retentionId string

		resp, httpRes, err := apiClient.RetentionAPI.GetRetention(context.Background(), retentionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RetentionAPIService ListRetentions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RetentionAPI.ListRetentions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RetentionAPIService SendRetentionByEmail", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var retentionId string

		resp, httpRes, err := apiClient.RetentionAPI.SendRetentionByEmail(context.Background(), retentionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
