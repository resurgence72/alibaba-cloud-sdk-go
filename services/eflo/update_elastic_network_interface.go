package eflo

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

// UpdateElasticNetworkInterface invokes the eflo.UpdateElasticNetworkInterface API synchronously
func (client *Client) UpdateElasticNetworkInterface(request *UpdateElasticNetworkInterfaceRequest) (response *UpdateElasticNetworkInterfaceResponse, err error) {
	response = CreateUpdateElasticNetworkInterfaceResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateElasticNetworkInterfaceWithChan invokes the eflo.UpdateElasticNetworkInterface API asynchronously
func (client *Client) UpdateElasticNetworkInterfaceWithChan(request *UpdateElasticNetworkInterfaceRequest) (<-chan *UpdateElasticNetworkInterfaceResponse, <-chan error) {
	responseChan := make(chan *UpdateElasticNetworkInterfaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateElasticNetworkInterface(request)
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

// UpdateElasticNetworkInterfaceWithCallback invokes the eflo.UpdateElasticNetworkInterface API asynchronously
func (client *Client) UpdateElasticNetworkInterfaceWithCallback(request *UpdateElasticNetworkInterfaceRequest, callback func(response *UpdateElasticNetworkInterfaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateElasticNetworkInterfaceResponse
		var err error
		defer close(result)
		response, err = client.UpdateElasticNetworkInterface(request)
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

// UpdateElasticNetworkInterfaceRequest is the request struct for api UpdateElasticNetworkInterface
type UpdateElasticNetworkInterfaceRequest struct {
	*requests.RpcRequest
	Description               string `position:"Body" name:"Description"`
	ElasticNetworkInterfaceId string `position:"Body" name:"ElasticNetworkInterfaceId"`
}

// UpdateElasticNetworkInterfaceResponse is the response struct for api UpdateElasticNetworkInterface
type UpdateElasticNetworkInterfaceResponse struct {
	*responses.BaseResponse
	Code      int     `json:"Code" xml:"Code"`
	Message   string  `json:"Message" xml:"Message"`
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Content   Content `json:"Content" xml:"Content"`
}

// CreateUpdateElasticNetworkInterfaceRequest creates a request to invoke UpdateElasticNetworkInterface API
func CreateUpdateElasticNetworkInterfaceRequest() (request *UpdateElasticNetworkInterfaceRequest) {
	request = &UpdateElasticNetworkInterfaceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("eflo", "2022-05-30", "UpdateElasticNetworkInterface", "eflo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateElasticNetworkInterfaceResponse creates a response to parse from UpdateElasticNetworkInterface response
func CreateUpdateElasticNetworkInterfaceResponse() (response *UpdateElasticNetworkInterfaceResponse) {
	response = &UpdateElasticNetworkInterfaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
