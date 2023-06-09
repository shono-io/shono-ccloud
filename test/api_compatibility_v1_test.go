/*
Confluent Cloud APIs

Testing CompatibilityV1ApiService

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

func Test_openapi_CompatibilityV1ApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CompatibilityV1ApiService TestCompatibilityBySubjectName", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var subject string
		var version string

		resp, httpRes, err := apiClient.CompatibilityV1Api.TestCompatibilityBySubjectName(context.Background(), subject, version).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CompatibilityV1ApiService TestCompatibilityForSubject", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var subject string

		resp, httpRes, err := apiClient.CompatibilityV1Api.TestCompatibilityForSubject(context.Background(), subject).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
