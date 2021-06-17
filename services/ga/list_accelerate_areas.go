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

// ListAccelerateAreas invokes the ga.ListAccelerateAreas API synchronously
func (client *Client) ListAccelerateAreas(request *ListAccelerateAreasRequest) (response *ListAccelerateAreasResponse, err error) {
	response = CreateListAccelerateAreasResponse()
	err = client.DoAction(request, response)
	return
}

// ListAccelerateAreasWithChan invokes the ga.ListAccelerateAreas API asynchronously
func (client *Client) ListAccelerateAreasWithChan(request *ListAccelerateAreasRequest) (<-chan *ListAccelerateAreasResponse, <-chan error) {
	responseChan := make(chan *ListAccelerateAreasResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAccelerateAreas(request)
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

// ListAccelerateAreasWithCallback invokes the ga.ListAccelerateAreas API asynchronously
func (client *Client) ListAccelerateAreasWithCallback(request *ListAccelerateAreasRequest, callback func(response *ListAccelerateAreasResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAccelerateAreasResponse
		var err error
		defer close(result)
		response, err = client.ListAccelerateAreas(request)
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

// ListAccelerateAreasRequest is the request struct for api ListAccelerateAreas
type ListAccelerateAreasRequest struct {
	*requests.RpcRequest
}

// ListAccelerateAreasResponse is the response struct for api ListAccelerateAreas
type ListAccelerateAreasResponse struct {
	*responses.BaseResponse
	RequestId string      `json:"RequestId" xml:"RequestId"`
	Areas     []AreasItem `json:"Areas" xml:"Areas"`
}

// CreateListAccelerateAreasRequest creates a request to invoke ListAccelerateAreas API
func CreateListAccelerateAreasRequest() (request *ListAccelerateAreasRequest) {
	request = &ListAccelerateAreasRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ga", "2019-11-20", "ListAccelerateAreas", "ga", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListAccelerateAreasResponse creates a response to parse from ListAccelerateAreas response
func CreateListAccelerateAreasResponse() (response *ListAccelerateAreasResponse) {
	response = &ListAccelerateAreasResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
