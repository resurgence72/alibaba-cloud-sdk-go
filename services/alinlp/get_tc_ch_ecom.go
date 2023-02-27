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

// GetTcChEcom invokes the alinlp.GetTcChEcom API synchronously
func (client *Client) GetTcChEcom(request *GetTcChEcomRequest) (response *GetTcChEcomResponse, err error) {
	response = CreateGetTcChEcomResponse()
	err = client.DoAction(request, response)
	return
}

// GetTcChEcomWithChan invokes the alinlp.GetTcChEcom API asynchronously
func (client *Client) GetTcChEcomWithChan(request *GetTcChEcomRequest) (<-chan *GetTcChEcomResponse, <-chan error) {
	responseChan := make(chan *GetTcChEcomResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetTcChEcom(request)
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

// GetTcChEcomWithCallback invokes the alinlp.GetTcChEcom API asynchronously
func (client *Client) GetTcChEcomWithCallback(request *GetTcChEcomRequest, callback func(response *GetTcChEcomResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetTcChEcomResponse
		var err error
		defer close(result)
		response, err = client.GetTcChEcom(request)
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

// GetTcChEcomRequest is the request struct for api GetTcChEcom
type GetTcChEcomRequest struct {
	*requests.RpcRequest
	ServiceCode string `position:"Body" name:"ServiceCode"`
	Text        string `position:"Body" name:"Text"`
}

// GetTcChEcomResponse is the response struct for api GetTcChEcom
type GetTcChEcomResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateGetTcChEcomRequest creates a request to invoke GetTcChEcom API
func CreateGetTcChEcomRequest() (request *GetTcChEcomRequest) {
	request = &GetTcChEcomRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alinlp", "2020-06-29", "GetTcChEcom", "", "")
	request.Method = requests.POST
	return
}

// CreateGetTcChEcomResponse creates a response to parse from GetTcChEcom response
func CreateGetTcChEcomResponse() (response *GetTcChEcomResponse) {
	response = &GetTcChEcomResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
