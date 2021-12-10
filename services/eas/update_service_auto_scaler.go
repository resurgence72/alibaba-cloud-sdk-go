package eas

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

// UpdateServiceAutoScaler invokes the eas.UpdateServiceAutoScaler API synchronously
func (client *Client) UpdateServiceAutoScaler(request *UpdateServiceAutoScalerRequest) (response *UpdateServiceAutoScalerResponse, err error) {
	response = CreateUpdateServiceAutoScalerResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateServiceAutoScalerWithChan invokes the eas.UpdateServiceAutoScaler API asynchronously
func (client *Client) UpdateServiceAutoScalerWithChan(request *UpdateServiceAutoScalerRequest) (<-chan *UpdateServiceAutoScalerResponse, <-chan error) {
	responseChan := make(chan *UpdateServiceAutoScalerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateServiceAutoScaler(request)
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

// UpdateServiceAutoScalerWithCallback invokes the eas.UpdateServiceAutoScaler API asynchronously
func (client *Client) UpdateServiceAutoScalerWithCallback(request *UpdateServiceAutoScalerRequest, callback func(response *UpdateServiceAutoScalerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateServiceAutoScalerResponse
		var err error
		defer close(result)
		response, err = client.UpdateServiceAutoScaler(request)
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

// UpdateServiceAutoScalerRequest is the request struct for api UpdateServiceAutoScaler
type UpdateServiceAutoScalerRequest struct {
	*requests.RoaRequest
	ServiceName string `position:"Path" name:"ServiceName"`
	ClusterId   string `position:"Path" name:"ClusterId"`
	Body        string `position:"Body" name:"body"`
}

// UpdateServiceAutoScalerResponse is the response struct for api UpdateServiceAutoScaler
type UpdateServiceAutoScalerResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateUpdateServiceAutoScalerRequest creates a request to invoke UpdateServiceAutoScaler API
func CreateUpdateServiceAutoScalerRequest() (request *UpdateServiceAutoScalerRequest) {
	request = &UpdateServiceAutoScalerRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("eas", "2021-07-01", "UpdateServiceAutoScaler", "/api/v2/services/[ClusterId]/[ServiceName]/autoscaler", "eas", "openAPI")
	request.Method = requests.PUT
	return
}

// CreateUpdateServiceAutoScalerResponse creates a response to parse from UpdateServiceAutoScaler response
func CreateUpdateServiceAutoScalerResponse() (response *UpdateServiceAutoScalerResponse) {
	response = &UpdateServiceAutoScalerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
