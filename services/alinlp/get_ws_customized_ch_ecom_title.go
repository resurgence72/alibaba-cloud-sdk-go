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

// GetWsCustomizedChEcomTitle invokes the alinlp.GetWsCustomizedChEcomTitle API synchronously
func (client *Client) GetWsCustomizedChEcomTitle(request *GetWsCustomizedChEcomTitleRequest) (response *GetWsCustomizedChEcomTitleResponse, err error) {
	response = CreateGetWsCustomizedChEcomTitleResponse()
	err = client.DoAction(request, response)
	return
}

// GetWsCustomizedChEcomTitleWithChan invokes the alinlp.GetWsCustomizedChEcomTitle API asynchronously
func (client *Client) GetWsCustomizedChEcomTitleWithChan(request *GetWsCustomizedChEcomTitleRequest) (<-chan *GetWsCustomizedChEcomTitleResponse, <-chan error) {
	responseChan := make(chan *GetWsCustomizedChEcomTitleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetWsCustomizedChEcomTitle(request)
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

// GetWsCustomizedChEcomTitleWithCallback invokes the alinlp.GetWsCustomizedChEcomTitle API asynchronously
func (client *Client) GetWsCustomizedChEcomTitleWithCallback(request *GetWsCustomizedChEcomTitleRequest, callback func(response *GetWsCustomizedChEcomTitleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetWsCustomizedChEcomTitleResponse
		var err error
		defer close(result)
		response, err = client.GetWsCustomizedChEcomTitle(request)
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

// GetWsCustomizedChEcomTitleRequest is the request struct for api GetWsCustomizedChEcomTitle
type GetWsCustomizedChEcomTitleRequest struct {
	*requests.RpcRequest
	ServiceCode string `position:"Body" name:"ServiceCode"`
	TokenizerId string `position:"Body" name:"TokenizerId"`
	Text        string `position:"Body" name:"Text"`
	OutType     string `position:"Body" name:"OutType"`
}

// GetWsCustomizedChEcomTitleResponse is the response struct for api GetWsCustomizedChEcomTitle
type GetWsCustomizedChEcomTitleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateGetWsCustomizedChEcomTitleRequest creates a request to invoke GetWsCustomizedChEcomTitle API
func CreateGetWsCustomizedChEcomTitleRequest() (request *GetWsCustomizedChEcomTitleRequest) {
	request = &GetWsCustomizedChEcomTitleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alinlp", "2020-06-29", "GetWsCustomizedChEcomTitle", "", "")
	request.Method = requests.POST
	return
}

// CreateGetWsCustomizedChEcomTitleResponse creates a response to parse from GetWsCustomizedChEcomTitle response
func CreateGetWsCustomizedChEcomTitleResponse() (response *GetWsCustomizedChEcomTitleResponse) {
	response = &GetWsCustomizedChEcomTitleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
