package graphrbac

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 1.0.1.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// ObjectsClient is the the Graph RBAC Management Client
type ObjectsClient struct {
	ManagementClient
}

// NewObjectsClient creates an instance of the ObjectsClient client.
func NewObjectsClient(tenantID string) ObjectsClient {
	return NewObjectsClientWithBaseURI(DefaultBaseURI, tenantID)
}

// NewObjectsClientWithBaseURI creates an instance of the ObjectsClient client.
func NewObjectsClientWithBaseURI(baseURI string, tenantID string) ObjectsClient {
	return ObjectsClient{NewWithBaseURI(baseURI, tenantID)}
}

// GetCurrentUser gets the details for the currently logged-in user.
func (client ObjectsClient) GetCurrentUser() (result AADObject, err error) {
	req, err := client.GetCurrentUserPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "graphrbac.ObjectsClient", "GetCurrentUser", nil, "Failure preparing request")
	}

	resp, err := client.GetCurrentUserSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "graphrbac.ObjectsClient", "GetCurrentUser", resp, "Failure sending request")
	}

	result, err = client.GetCurrentUserResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ObjectsClient", "GetCurrentUser", resp, "Failure responding to request")
	}

	return
}

// GetCurrentUserPreparer prepares the GetCurrentUser request.
func (client ObjectsClient) GetCurrentUserPreparer() (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"tenantID": autorest.Encode("path", client.TenantID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{tenantID}/me", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetCurrentUserSender sends the GetCurrentUser request. The method will close the
// http.Response Body if it receives an error.
func (client ObjectsClient) GetCurrentUserSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetCurrentUserResponder handles the response to the GetCurrentUser request. The method always
// closes the http.Response Body.
func (client ObjectsClient) GetCurrentUserResponder(resp *http.Response) (result AADObject, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
