package dcdn

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

// StartDcdnDomain invokes the dcdn.StartDcdnDomain API synchronously
// api document: https://help.aliyun.com/api/dcdn/startdcdndomain.html
func (client *Client) StartDcdnDomain(request *StartDcdnDomainRequest) (response *StartDcdnDomainResponse, err error) {
	response = CreateStartDcdnDomainResponse()
	err = client.DoAction(request, response)
	return
}

// StartDcdnDomainWithChan invokes the dcdn.StartDcdnDomain API asynchronously
// api document: https://help.aliyun.com/api/dcdn/startdcdndomain.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) StartDcdnDomainWithChan(request *StartDcdnDomainRequest) (<-chan *StartDcdnDomainResponse, <-chan error) {
	responseChan := make(chan *StartDcdnDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartDcdnDomain(request)
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

// StartDcdnDomainWithCallback invokes the dcdn.StartDcdnDomain API asynchronously
// api document: https://help.aliyun.com/api/dcdn/startdcdndomain.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) StartDcdnDomainWithCallback(request *StartDcdnDomainRequest, callback func(response *StartDcdnDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartDcdnDomainResponse
		var err error
		defer close(result)
		response, err = client.StartDcdnDomain(request)
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

// StartDcdnDomainRequest is the request struct for api StartDcdnDomain
type StartDcdnDomainRequest struct {
	*requests.RpcRequest
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	DomainName    string           `position:"Query" name:"DomainName"`
}

// StartDcdnDomainResponse is the response struct for api StartDcdnDomain
type StartDcdnDomainResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateStartDcdnDomainRequest creates a request to invoke StartDcdnDomain API
func CreateStartDcdnDomainRequest() (request *StartDcdnDomainRequest) {
	request = &StartDcdnDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "StartDcdnDomain", "dcdn", "openAPI")
	return
}

// CreateStartDcdnDomainResponse creates a response to parse from StartDcdnDomain response
func CreateStartDcdnDomainResponse() (response *StartDcdnDomainResponse) {
	response = &StartDcdnDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
