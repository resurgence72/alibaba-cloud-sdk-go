package videoenhan

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

// MergeVideoFace invokes the videoenhan.MergeVideoFace API synchronously
func (client *Client) MergeVideoFace(request *MergeVideoFaceRequest) (response *MergeVideoFaceResponse, err error) {
	response = CreateMergeVideoFaceResponse()
	err = client.DoAction(request, response)
	return
}

// MergeVideoFaceWithChan invokes the videoenhan.MergeVideoFace API asynchronously
func (client *Client) MergeVideoFaceWithChan(request *MergeVideoFaceRequest) (<-chan *MergeVideoFaceResponse, <-chan error) {
	responseChan := make(chan *MergeVideoFaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.MergeVideoFace(request)
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

// MergeVideoFaceWithCallback invokes the videoenhan.MergeVideoFace API asynchronously
func (client *Client) MergeVideoFaceWithCallback(request *MergeVideoFaceRequest, callback func(response *MergeVideoFaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *MergeVideoFaceResponse
		var err error
		defer close(result)
		response, err = client.MergeVideoFace(request)
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

// MergeVideoFaceRequest is the request struct for api MergeVideoFace
type MergeVideoFaceRequest struct {
	*requests.RpcRequest
	ReferenceURL string           `position:"Body" name:"ReferenceURL"`
	Enhance      requests.Boolean `position:"Body" name:"Enhance"`
	PostURL      string           `position:"Body" name:"PostURL"`
	Async        requests.Boolean `position:"Body" name:"Async"`
	VideoURL     string           `position:"Body" name:"VideoURL"`
	AddWatermark requests.Boolean `position:"Body" name:"AddWatermark"`
}

// MergeVideoFaceResponse is the response struct for api MergeVideoFace
type MergeVideoFaceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateMergeVideoFaceRequest creates a request to invoke MergeVideoFace API
func CreateMergeVideoFaceRequest() (request *MergeVideoFaceRequest) {
	request = &MergeVideoFaceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("videoenhan", "2020-03-20", "MergeVideoFace", "videoenhan", "openAPI")
	request.Method = requests.POST
	return
}

// CreateMergeVideoFaceResponse creates a response to parse from MergeVideoFace response
func CreateMergeVideoFaceResponse() (response *MergeVideoFaceResponse) {
	response = &MergeVideoFaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
