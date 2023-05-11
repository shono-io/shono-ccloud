/*
Confluent Cloud APIs

Testing EntityV1ApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	openapiclient "github.com/shono-io/shono-ccloud"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_openapi_EntityV1ApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test EntityV1ApiService CreateTags", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.EntityV1Api.CreateTags(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EntityV1ApiService DeleteTag", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var typeName string
		var qualifiedName string
		var tagName string

		httpRes, err := apiClient.EntityV1Api.DeleteTag(context.Background(), typeName, qualifiedName, tagName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EntityV1ApiService GetByUniqueAttributes", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var typeName string
		var qualifiedName string

		resp, httpRes, err := apiClient.EntityV1Api.GetByUniqueAttributes(context.Background(), typeName, qualifiedName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EntityV1ApiService GetTags", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var typeName string
		var qualifiedName string

		resp, httpRes, err := apiClient.EntityV1Api.GetTags(context.Background(), typeName, qualifiedName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EntityV1ApiService UpdateTags", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.EntityV1Api.UpdateTags(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
