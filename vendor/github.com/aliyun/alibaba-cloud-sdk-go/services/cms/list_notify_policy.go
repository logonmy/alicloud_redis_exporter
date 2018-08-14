package cms

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

// ListNotifyPolicy invokes the cms.ListNotifyPolicy API synchronously
// api document: https://help.aliyun.com/api/cms/listnotifypolicy.html
func (client *Client) ListNotifyPolicy(request *ListNotifyPolicyRequest) (response *ListNotifyPolicyResponse, err error) {
	response = CreateListNotifyPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// ListNotifyPolicyWithChan invokes the cms.ListNotifyPolicy API asynchronously
// api document: https://help.aliyun.com/api/cms/listnotifypolicy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListNotifyPolicyWithChan(request *ListNotifyPolicyRequest) (<-chan *ListNotifyPolicyResponse, <-chan error) {
	responseChan := make(chan *ListNotifyPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListNotifyPolicy(request)
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

// ListNotifyPolicyWithCallback invokes the cms.ListNotifyPolicy API asynchronously
// api document: https://help.aliyun.com/api/cms/listnotifypolicy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListNotifyPolicyWithCallback(request *ListNotifyPolicyRequest, callback func(response *ListNotifyPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListNotifyPolicyResponse
		var err error
		defer close(result)
		response, err = client.ListNotifyPolicy(request)
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

// ListNotifyPolicyRequest is the request struct for api ListNotifyPolicy
type ListNotifyPolicyRequest struct {
	*requests.RpcRequest
	PolicyType string           `position:"Query" name:"PolicyType"`
	AlertName  string           `position:"Query" name:"AlertName"`
	GroupId    string           `position:"Query" name:"GroupId"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	Id         string           `position:"Query" name:"Id"`
	Dimensions string           `position:"Query" name:"Dimensions"`
}

// ListNotifyPolicyResponse is the response struct for api ListNotifyPolicy
type ListNotifyPolicyResponse struct {
	*responses.BaseResponse
	Code             string           `json:"Code" xml:"Code"`
	Message          string           `json:"Message" xml:"Message"`
	Success          string           `json:"Success" xml:"Success"`
	RequestId        string           `json:"RequestId" xml:"RequestId"`
	Total            int              `json:"Total" xml:"Total"`
	NotifyPolicyList NotifyPolicyList `json:"NotifyPolicyList" xml:"NotifyPolicyList"`
}

// CreateListNotifyPolicyRequest creates a request to invoke ListNotifyPolicy API
func CreateListNotifyPolicyRequest() (request *ListNotifyPolicyRequest) {
	request = &ListNotifyPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2018-03-08", "ListNotifyPolicy", "cms", "openAPI")
	return
}

// CreateListNotifyPolicyResponse creates a response to parse from ListNotifyPolicy response
func CreateListNotifyPolicyResponse() (response *ListNotifyPolicyResponse) {
	response = &ListNotifyPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
