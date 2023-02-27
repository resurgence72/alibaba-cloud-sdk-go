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

// GetTableQAServiceInfoById invokes the alinlp.GetTableQAServiceInfoById API synchronously
func (client *Client) GetTableQAServiceInfoById(request *GetTableQAServiceInfoByIdRequest) (response *GetTableQAServiceInfoByIdResponse, err error) {
	response = CreateGetTableQAServiceInfoByIdResponse()
	err = client.DoAction(request, response)
	return
}

// GetTableQAServiceInfoByIdWithChan invokes the alinlp.GetTableQAServiceInfoById API asynchronously
func (client *Client) GetTableQAServiceInfoByIdWithChan(request *GetTableQAServiceInfoByIdRequest) (<-chan *GetTableQAServiceInfoByIdResponse, <-chan error) {
	responseChan := make(chan *GetTableQAServiceInfoByIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetTableQAServiceInfoById(request)
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

// GetTableQAServiceInfoByIdWithCallback invokes the alinlp.GetTableQAServiceInfoById API asynchronously
func (client *Client) GetTableQAServiceInfoByIdWithCallback(request *GetTableQAServiceInfoByIdRequest, callback func(response *GetTableQAServiceInfoByIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetTableQAServiceInfoByIdResponse
		var err error
		defer close(result)
		response, err = client.GetTableQAServiceInfoById(request)
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

// GetTableQAServiceInfoByIdRequest is the request struct for api GetTableQAServiceInfoById
type GetTableQAServiceInfoByIdRequest struct {
	*requests.RpcRequest
	ServiceCode string `position:"Body" name:"ServiceCode"`
	ServiceId   string `position:"Body" name:"ServiceId"`
}

// GetTableQAServiceInfoByIdResponse is the response struct for api GetTableQAServiceInfoById
type GetTableQAServiceInfoByIdResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateGetTableQAServiceInfoByIdRequest creates a request to invoke GetTableQAServiceInfoById API
func CreateGetTableQAServiceInfoByIdRequest() (request *GetTableQAServiceInfoByIdRequest) {
	request = &GetTableQAServiceInfoByIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alinlp", "2020-06-29", "GetTableQAServiceInfoById", "", "")
	request.Method = requests.POST
	return
}

// CreateGetTableQAServiceInfoByIdResponse creates a response to parse from GetTableQAServiceInfoById response
func CreateGetTableQAServiceInfoByIdResponse() (response *GetTableQAServiceInfoByIdResponse) {
	response = &GetTableQAServiceInfoByIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
