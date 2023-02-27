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

// GetNerChEcom invokes the alinlp.GetNerChEcom API synchronously
func (client *Client) GetNerChEcom(request *GetNerChEcomRequest) (response *GetNerChEcomResponse, err error) {
	response = CreateGetNerChEcomResponse()
	err = client.DoAction(request, response)
	return
}

// GetNerChEcomWithChan invokes the alinlp.GetNerChEcom API asynchronously
func (client *Client) GetNerChEcomWithChan(request *GetNerChEcomRequest) (<-chan *GetNerChEcomResponse, <-chan error) {
	responseChan := make(chan *GetNerChEcomResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetNerChEcom(request)
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

// GetNerChEcomWithCallback invokes the alinlp.GetNerChEcom API asynchronously
func (client *Client) GetNerChEcomWithCallback(request *GetNerChEcomRequest, callback func(response *GetNerChEcomResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetNerChEcomResponse
		var err error
		defer close(result)
		response, err = client.GetNerChEcom(request)
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

// GetNerChEcomRequest is the request struct for api GetNerChEcom
type GetNerChEcomRequest struct {
	*requests.RpcRequest
	LexerId     string `position:"Body" name:"LexerId"`
	ServiceCode string `position:"Body" name:"ServiceCode"`
	Text        string `position:"Body" name:"Text"`
}

// GetNerChEcomResponse is the response struct for api GetNerChEcom
type GetNerChEcomResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateGetNerChEcomRequest creates a request to invoke GetNerChEcom API
func CreateGetNerChEcomRequest() (request *GetNerChEcomRequest) {
	request = &GetNerChEcomRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alinlp", "2020-06-29", "GetNerChEcom", "", "")
	request.Method = requests.POST
	return
}

// CreateGetNerChEcomResponse creates a response to parse from GetNerChEcom response
func CreateGetNerChEcomResponse() (response *GetNerChEcomResponse) {
	response = &GetNerChEcomResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
