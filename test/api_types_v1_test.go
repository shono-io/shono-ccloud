/*
Confluent Cloud APIs

Testing TypesV1ApiService

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

func Test_openapi_TypesV1ApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TypesV1ApiService CreateTagDefs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TypesV1Api.CreateTagDefs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TypesV1ApiService DeleteTagDef", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tagName string

		httpRes, err := apiClient.TypesV1Api.DeleteTagDef(context.Background(), tagName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TypesV1ApiService GetAllTagDefs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TypesV1Api.GetAllTagDefs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TypesV1ApiService GetTagDefByName", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tagName string

		resp, httpRes, err := apiClient.TypesV1Api.GetTagDefByName(context.Background(), tagName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TypesV1ApiService UpdateTagDefs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TypesV1Api.UpdateTagDefs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
