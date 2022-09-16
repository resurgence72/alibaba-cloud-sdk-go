package antiddos_public

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

// ModifyDdosStatus invokes the antiddos_public.ModifyDdosStatus API synchronously
func (client *Client) ModifyDdosStatus(request *ModifyDdosStatusRequest) (response *ModifyDdosStatusResponse, err error) {
	response = CreateModifyDdosStatusResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDdosStatusWithChan invokes the antiddos_public.ModifyDdosStatus API asynchronously
func (client *Client) ModifyDdosStatusWithChan(request *ModifyDdosStatusRequest) (<-chan *ModifyDdosStatusResponse, <-chan error) {
	responseChan := make(chan *ModifyDdosStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDdosStatus(request)
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

// ModifyDdosStatusWithCallback invokes the antiddos_public.ModifyDdosStatus API asynchronously
func (client *Client) ModifyDdosStatusWithCallback(request *ModifyDdosStatusRequest, callback func(response *ModifyDdosStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDdosStatusResponse
		var err error
		defer close(result)
		response, err = client.ModifyDdosStatus(request)
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

// ModifyDdosStatusRequest is the request struct for api ModifyDdosStatus
type ModifyDdosStatusRequest struct {
	*requests.RpcRequest
	InternetIp   string `position:"Query" name:"InternetIp"`
	SourceIp     string `position:"Query" name:"SourceIp"`
	DdosRegionId string `position:"Query" name:"DdosRegionId"`
	InstanceType string `position:"Query" name:"InstanceType"`
	Lang         string `position:"Query" name:"Lang"`
	InstanceId   string `position:"Query" name:"InstanceId"`
}

// ModifyDdosStatusResponse is the response struct for api ModifyDdosStatus
type ModifyDdosStatusResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDdosStatusRequest creates a request to invoke ModifyDdosStatus API
func CreateModifyDdosStatusRequest() (request *ModifyDdosStatusRequest) {
	request = &ModifyDdosStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("antiddos-public", "2017-05-18", "ModifyDdosStatus", "ddosbasic", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyDdosStatusResponse creates a response to parse from ModifyDdosStatus response
func CreateModifyDdosStatusResponse() (response *ModifyDdosStatusResponse) {
	response = &ModifyDdosStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
