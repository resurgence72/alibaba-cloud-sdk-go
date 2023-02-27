package alinlp

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

// RequestTableQA invokes the alinlp.RequestTableQA API synchronously
func (client *Client) RequestTableQA(request *RequestTableQARequest) (response *RequestTableQAResponse, err error) {
	response = CreateRequestTableQAResponse()
	err = client.DoAction(request, response)
	return
}

// RequestTableQAWithChan invokes the alinlp.RequestTableQA API asynchronously
func (client *Client) RequestTableQAWithChan(request *RequestTableQARequest) (<-chan *RequestTableQAResponse, <-chan error) {
	responseChan := make(chan *RequestTableQAResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RequestTableQA(request)
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

// RequestTableQAWithCallback invokes the alinlp.RequestTableQA API asynchronously
func (client *Client) RequestTableQAWithCallback(request *RequestTableQARequest, callback func(response *RequestTableQAResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RequestTableQAResponse
		var err error
		defer close(result)
		response, err = client.RequestTableQA(request)
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

// RequestTableQARequest is the request struct for api RequestTableQA
type RequestTableQARequest struct {
	*requests.RpcRequest
	Params      string `position:"Body" name:"Params"`
	ServiceCode string `position:"Body" name:"ServiceCode"`
}

// RequestTableQAResponse is the response struct for api RequestTableQA
type RequestTableQAResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateRequestTableQARequest creates a request to invoke RequestTableQA API
func CreateRequestTableQARequest() (request *RequestTableQARequest) {
	request = &RequestTableQARequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alinlp", "2020-06-29", "RequestTableQA", "", "")
	request.Method = requests.POST
	return
}

// CreateRequestTableQAResponse creates a response to parse from RequestTableQA response
func CreateRequestTableQAResponse() (response *RequestTableQAResponse) {
	response = &RequestTableQAResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
