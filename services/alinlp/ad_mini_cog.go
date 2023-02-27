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

// ADMiniCog invokes the alinlp.ADMiniCog API synchronously
func (client *Client) ADMiniCog(request *ADMiniCogRequest) (response *ADMiniCogResponse, err error) {
	response = CreateADMiniCogResponse()
	err = client.DoAction(request, response)
	return
}

// ADMiniCogWithChan invokes the alinlp.ADMiniCog API asynchronously
func (client *Client) ADMiniCogWithChan(request *ADMiniCogRequest) (<-chan *ADMiniCogResponse, <-chan error) {
	responseChan := make(chan *ADMiniCogResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ADMiniCog(request)
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

// ADMiniCogWithCallback invokes the alinlp.ADMiniCog API asynchronously
func (client *Client) ADMiniCogWithCallback(request *ADMiniCogRequest, callback func(response *ADMiniCogResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ADMiniCogResponse
		var err error
		defer close(result)
		response, err = client.ADMiniCog(request)
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

// ADMiniCogRequest is the request struct for api ADMiniCog
type ADMiniCogRequest struct {
	*requests.RpcRequest
	Params      string `position:"Body" name:"Params"`
	ServiceCode string `position:"Body" name:"ServiceCode"`
}

// ADMiniCogResponse is the response struct for api ADMiniCog
type ADMiniCogResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateADMiniCogRequest creates a request to invoke ADMiniCog API
func CreateADMiniCogRequest() (request *ADMiniCogRequest) {
	request = &ADMiniCogRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alinlp", "2020-06-29", "ADMiniCog", "", "")
	request.Method = requests.POST
	return
}

// CreateADMiniCogResponse creates a response to parse from ADMiniCog response
func CreateADMiniCogResponse() (response *ADMiniCogResponse) {
	response = &ADMiniCogResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
