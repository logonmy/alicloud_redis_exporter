package mts

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

// SubmitVideoSplitJob invokes the mts.SubmitVideoSplitJob API synchronously
// api document: https://help.aliyun.com/api/mts/submitvideosplitjob.html
func (client *Client) SubmitVideoSplitJob(request *SubmitVideoSplitJobRequest) (response *SubmitVideoSplitJobResponse, err error) {
	response = CreateSubmitVideoSplitJobResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitVideoSplitJobWithChan invokes the mts.SubmitVideoSplitJob API asynchronously
// api document: https://help.aliyun.com/api/mts/submitvideosplitjob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SubmitVideoSplitJobWithChan(request *SubmitVideoSplitJobRequest) (<-chan *SubmitVideoSplitJobResponse, <-chan error) {
	responseChan := make(chan *SubmitVideoSplitJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitVideoSplitJob(request)
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

// SubmitVideoSplitJobWithCallback invokes the mts.SubmitVideoSplitJob API asynchronously
// api document: https://help.aliyun.com/api/mts/submitvideosplitjob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SubmitVideoSplitJobWithCallback(request *SubmitVideoSplitJobRequest, callback func(response *SubmitVideoSplitJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitVideoSplitJobResponse
		var err error
		defer close(result)
		response, err = client.SubmitVideoSplitJob(request)
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

// SubmitVideoSplitJobRequest is the request struct for api SubmitVideoSplitJob
type SubmitVideoSplitJobRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Input                string           `position:"Query" name:"Input"`
	PipelineId           string           `position:"Query" name:"PipelineId"`
	VideoSplitConfig     string           `position:"Query" name:"VideoSplitConfig"`
	UserData             string           `position:"Query" name:"UserData"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

// SubmitVideoSplitJobResponse is the response struct for api SubmitVideoSplitJob
type SubmitVideoSplitJobResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	JobId     string `json:"JobId" xml:"JobId"`
}

// CreateSubmitVideoSplitJobRequest creates a request to invoke SubmitVideoSplitJob API
func CreateSubmitVideoSplitJobRequest() (request *SubmitVideoSplitJobRequest) {
	request = &SubmitVideoSplitJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "SubmitVideoSplitJob", "mts", "openAPI")
	return
}

// CreateSubmitVideoSplitJobResponse creates a response to parse from SubmitVideoSplitJob response
func CreateSubmitVideoSplitJobResponse() (response *SubmitVideoSplitJobResponse) {
	response = &SubmitVideoSplitJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
