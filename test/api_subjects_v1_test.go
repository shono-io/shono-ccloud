/*
Confluent Cloud APIs

Testing SubjectsV1ApiService

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

func Test_openapi_SubjectsV1ApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SubjectsV1ApiService DeleteSchemaVersion", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subject string
		var version string

		resp, httpRes, err := apiClient.SubjectsV1Api.DeleteSchemaVersion(context.Background(), subject, version).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubjectsV1ApiService DeleteSubject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subject string

		resp, httpRes, err := apiClient.SubjectsV1Api.DeleteSubject(context.Background(), subject).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubjectsV1ApiService GetReferencedBy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subject string
		var version string

		resp, httpRes, err := apiClient.SubjectsV1Api.GetReferencedBy(context.Background(), subject, version).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubjectsV1ApiService GetSchemaByVersion", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subject string
		var version string

		resp, httpRes, err := apiClient.SubjectsV1Api.GetSchemaByVersion(context.Background(), subject, version).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubjectsV1ApiService GetSchemaOnly1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subject string
		var version string

		resp, httpRes, err := apiClient.SubjectsV1Api.GetSchemaOnly1(context.Background(), subject, version).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubjectsV1ApiService List", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SubjectsV1Api.List(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubjectsV1ApiService ListVersions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subject string

		resp, httpRes, err := apiClient.SubjectsV1Api.ListVersions(context.Background(), subject).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubjectsV1ApiService LookUpSchemaUnderSubject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subject string

		resp, httpRes, err := apiClient.SubjectsV1Api.LookUpSchemaUnderSubject(context.Background(), subject).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubjectsV1ApiService Register", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subject string

		resp, httpRes, err := apiClient.SubjectsV1Api.Register(context.Background(), subject).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
