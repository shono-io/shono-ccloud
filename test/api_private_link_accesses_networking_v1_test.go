/*
Confluent Cloud APIs

Testing PrivateLinkAccessesNetworkingV1ApiService

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

func Test_openapi_PrivateLinkAccessesNetworkingV1ApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PrivateLinkAccessesNetworkingV1ApiService CreateNetworkingV1PrivateLinkAccess", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PrivateLinkAccessesNetworkingV1Api.CreateNetworkingV1PrivateLinkAccess(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PrivateLinkAccessesNetworkingV1ApiService DeleteNetworkingV1PrivateLinkAccess", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.PrivateLinkAccessesNetworkingV1Api.DeleteNetworkingV1PrivateLinkAccess(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PrivateLinkAccessesNetworkingV1ApiService GetNetworkingV1PrivateLinkAccess", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.PrivateLinkAccessesNetworkingV1Api.GetNetworkingV1PrivateLinkAccess(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PrivateLinkAccessesNetworkingV1ApiService ListNetworkingV1PrivateLinkAccesses", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PrivateLinkAccessesNetworkingV1Api.ListNetworkingV1PrivateLinkAccesses(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PrivateLinkAccessesNetworkingV1ApiService UpdateNetworkingV1PrivateLinkAccess", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.PrivateLinkAccessesNetworkingV1Api.UpdateNetworkingV1PrivateLinkAccess(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
