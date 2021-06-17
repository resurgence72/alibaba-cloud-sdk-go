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

// CreateEndpointGroup invokes the ga.CreateEndpointGroup API synchronously
func (client *Client) CreateEndpointGroup(request *CreateEndpointGroupRequest) (response *CreateEndpointGroupResponse, err error) {
	response = CreateCreateEndpointGroupResponse()
	err = client.DoAction(request, response)
	return
}

// CreateEndpointGroupWithChan invokes the ga.CreateEndpointGroup API asynchronously
func (client *Client) CreateEndpointGroupWithChan(request *CreateEndpointGroupRequest) (<-chan *CreateEndpointGroupResponse, <-chan error) {
	responseChan := make(chan *CreateEndpointGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateEndpointGroup(request)
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

// CreateEndpointGroupWithCallback invokes the ga.CreateEndpointGroup API asynchronously
func (client *Client) CreateEndpointGroupWithCallback(request *CreateEndpointGroupRequest, callback func(response *CreateEndpointGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateEndpointGroupResponse
		var err error
		defer close(result)
		response, err = client.CreateEndpointGroup(request)
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

// CreateEndpointGroupRequest is the request struct for api CreateEndpointGroup
type CreateEndpointGroupRequest struct {
	*requests.RpcRequest
	PortOverrides              *[]CreateEndpointGroupPortOverrides          `position:"Query" name:"PortOverrides"  type:"Repeated"`
	ClientToken                string                                       `position:"Query" name:"ClientToken"`
	HealthCheckIntervalSeconds requests.Integer                             `position:"Query" name:"HealthCheckIntervalSeconds"`
	Description                string                                       `position:"Query" name:"Description"`
	HealthCheckProtocol        string                                       `position:"Query" name:"HealthCheckProtocol"`
	EndpointRequestProtocol    string                                       `position:"Query" name:"EndpointRequestProtocol"`
	ListenerId                 string                                       `position:"Query" name:"ListenerId"`
	HealthCheckPath            string                                       `position:"Query" name:"HealthCheckPath"`
	EndpointConfigurations     *[]CreateEndpointGroupEndpointConfigurations `position:"Query" name:"EndpointConfigurations"  type:"Repeated"`
	EndpointGroupType          string                                       `position:"Query" name:"EndpointGroupType"`
	AcceleratorId              string                                       `position:"Query" name:"AcceleratorId"`
	TrafficPercentage          requests.Integer                             `position:"Query" name:"TrafficPercentage"`
	HealthCheckPort            requests.Integer                             `position:"Query" name:"HealthCheckPort"`
	ThresholdCount             requests.Integer                             `position:"Query" name:"ThresholdCount"`
	EndpointGroupRegion        string                                       `position:"Query" name:"EndpointGroupRegion"`
	Name                       string                                       `position:"Query" name:"Name"`
}

// CreateEndpointGroupPortOverrides is a repeated param struct in CreateEndpointGroupRequest
type CreateEndpointGroupPortOverrides struct {
	ListenerPort string `name:"ListenerPort"`
	EndpointPort string `name:"EndpointPort"`
}

// CreateEndpointGroupEndpointConfigurations is a repeated param struct in CreateEndpointGroupRequest
type CreateEndpointGroupEndpointConfigurations struct {
	Type                       string `name:"Type"`
	EnableClientIPPreservation string `name:"EnableClientIPPreservation"`
	Weight                     string `name:"Weight"`
	EnableProxyProtocol        string `name:"EnableProxyProtocol"`
	Endpoint                   string `name:"Endpoint"`
}

// CreateEndpointGroupResponse is the response struct for api CreateEndpointGroup
type CreateEndpointGroupResponse struct {
	*responses.BaseResponse
	EndpointGroupId string `json:"EndpointGroupId" xml:"EndpointGroupId"`
	RequestId       string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateEndpointGroupRequest creates a request to invoke CreateEndpointGroup API
func CreateCreateEndpointGroupRequest() (request *CreateEndpointGroupRequest) {
	request = &CreateEndpointGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ga", "2019-11-20", "CreateEndpointGroup", "ga", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateEndpointGroupResponse creates a response to parse from CreateEndpointGroup response
func CreateCreateEndpointGroupResponse() (response *CreateEndpointGroupResponse) {
	response = &CreateEndpointGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
