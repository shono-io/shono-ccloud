/*
Confluent Cloud APIs

Testing IdentityPoolsIamV2ApiService

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

func Test_openapi_IdentityPoolsIamV2ApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test IdentityPoolsIamV2ApiService CreateIamV2IdentityPool", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var providerId string

		resp, httpRes, err := apiClient.IdentityPoolsIamV2Api.CreateIamV2IdentityPool(context.Background(), providerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IdentityPoolsIamV2ApiService DeleteIamV2IdentityPool", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var providerId string
		var id string

		httpRes, err := apiClient.IdentityPoolsIamV2Api.DeleteIamV2IdentityPool(context.Background(), providerId, id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IdentityPoolsIamV2ApiService GetIamV2IdentityPool", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var providerId string
		var id string

		resp, httpRes, err := apiClient.IdentityPoolsIamV2Api.GetIamV2IdentityPool(context.Background(), providerId, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IdentityPoolsIamV2ApiService ListIamV2IdentityPools", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var providerId string

		resp, httpRes, err := apiClient.IdentityPoolsIamV2Api.ListIamV2IdentityPools(context.Background(), providerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IdentityPoolsIamV2ApiService UpdateIamV2IdentityPool", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var providerId string
		var id string

		resp, httpRes, err := apiClient.IdentityPoolsIamV2Api.UpdateIamV2IdentityPool(context.Background(), providerId, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
