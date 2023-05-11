/*
Confluent Cloud APIs

Testing ConfigV1ApiService

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

func Test_openapi_ConfigV1ApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ConfigV1ApiService DeleteSubjectConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subject string

		resp, httpRes, err := apiClient.ConfigV1Api.DeleteSubjectConfig(context.Background(), subject).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigV1ApiService DeleteTopLevelConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ConfigV1Api.DeleteTopLevelConfig(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigV1ApiService GetClusterConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ConfigV1Api.GetClusterConfig(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigV1ApiService GetSubjectLevelConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subject string

		resp, httpRes, err := apiClient.ConfigV1Api.GetSubjectLevelConfig(context.Background(), subject).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigV1ApiService GetTopLevelConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ConfigV1Api.GetTopLevelConfig(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigV1ApiService UpdateSubjectLevelConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subject string

		resp, httpRes, err := apiClient.ConfigV1Api.UpdateSubjectLevelConfig(context.Background(), subject).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigV1ApiService UpdateTopLevelConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ConfigV1Api.UpdateTopLevelConfig(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}