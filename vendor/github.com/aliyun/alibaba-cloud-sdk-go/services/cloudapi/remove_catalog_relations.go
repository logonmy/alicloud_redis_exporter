package cloudapi

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// RemoveCatalogRelations invokes the cloudapi.RemoveCatalogRelations API synchronously
// api document: https://help.aliyun.com/api/cloudapi/removecatalogrelations.html
func (client *Client) RemoveCatalogRelations(request *RemoveCatalogRelationsRequest) (response *RemoveCatalogRelationsResponse, err error) {
	response = CreateRemoveCatalogRelationsResponse()
	err = client.DoAction(request, response)
	return
}

// RemoveCatalogRelationsWithChan invokes the cloudapi.RemoveCatalogRelations API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/removecatalogrelations.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RemoveCatalogRelationsWithChan(request *RemoveCatalogRelationsRequest) (<-chan *RemoveCatalogRelationsResponse, <-chan error) {
	responseChan := make(chan *RemoveCatalogRelationsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveCatalogRelations(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// RemoveCatalogRelationsWithCallback invokes the cloudapi.RemoveCatalogRelations API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/removecatalogrelations.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RemoveCatalogRelationsWithCallback(request *RemoveCatalogRelationsRequest, callback func(response *RemoveCatalogRelationsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveCatalogRelationsResponse
		var err error
		defer close(result)
		response, err = client.RemoveCatalogRelations(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// RemoveCatalogRelationsRequest is the request struct for api RemoveCatalogRelations
type RemoveCatalogRelationsRequest struct {
	*requests.RpcRequest
	CatalogId string `position:"Query" name:"CatalogId"`
}

// RemoveCatalogRelationsResponse is the response struct for api RemoveCatalogRelations
type RemoveCatalogRelationsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRemoveCatalogRelationsRequest creates a request to invoke RemoveCatalogRelations API
func CreateRemoveCatalogRelationsRequest() (request *RemoveCatalogRelationsRequest) {
	request = &RemoveCatalogRelationsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "RemoveCatalogRelations", "apigateway", "openAPI")
	return
}

// CreateRemoveCatalogRelationsResponse creates a response to parse from RemoveCatalogRelations response
func CreateRemoveCatalogRelationsResponse() (response *RemoveCatalogRelationsResponse) {
	response = &RemoveCatalogRelationsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
