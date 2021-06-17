package ga

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

// DeleteEndpointGroup invokes the ga.DeleteEndpointGroup API synchronously
func (client *Client) DeleteEndpointGroup(request *DeleteEndpointGroupRequest) (response *DeleteEndpointGroupResponse, err error) {
	response = CreateDeleteEndpointGroupResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteEndpointGroupWithChan invokes the ga.DeleteEndpointGroup API asynchronously
func (client *Client) DeleteEndpointGroupWithChan(request *DeleteEndpointGroupRequest) (<-chan *DeleteEndpointGroupResponse, <-chan error) {
	responseChan := make(chan *DeleteEndpointGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteEndpointGroup(request)
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

// DeleteEndpointGroupWithCallback invokes the ga.DeleteEndpointGroup API asynchronously
func (client *Client) DeleteEndpointGroupWithCallback(request *DeleteEndpointGroupRequest, callback func(response *DeleteEndpointGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteEndpointGroupResponse
		var err error
		defer close(result)
		response, err = client.DeleteEndpointGroup(request)
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

// DeleteEndpointGroupRequest is the request struct for api DeleteEndpointGroup
type DeleteEndpointGroupRequest struct {
	*requests.RpcRequest
	ClientToken     string `position:"Query" name:"ClientToken"`
	AcceleratorId   string `position:"Query" name:"AcceleratorId"`
	EndpointGroupId string `position:"Query" name:"EndpointGroupId"`
}

// DeleteEndpointGroupResponse is the response struct for api DeleteEndpointGroup
type DeleteEndpointGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteEndpointGroupRequest creates a request to invoke DeleteEndpointGroup API
func CreateDeleteEndpointGroupRequest() (request *DeleteEndpointGroupRequest) {
	request = &DeleteEndpointGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ga", "2019-11-20", "DeleteEndpointGroup", "ga", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteEndpointGroupResponse creates a response to parse from DeleteEndpointGroup response
func CreateDeleteEndpointGroupResponse() (response *DeleteEndpointGroupResponse) {
	response = &DeleteEndpointGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
