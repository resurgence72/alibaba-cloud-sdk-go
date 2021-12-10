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

// DescribeServiceMirror invokes the eas.DescribeServiceMirror API synchronously
func (client *Client) DescribeServiceMirror(request *DescribeServiceMirrorRequest) (response *DescribeServiceMirrorResponse, err error) {
	response = CreateDescribeServiceMirrorResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeServiceMirrorWithChan invokes the eas.DescribeServiceMirror API asynchronously
func (client *Client) DescribeServiceMirrorWithChan(request *DescribeServiceMirrorRequest) (<-chan *DescribeServiceMirrorResponse, <-chan error) {
	responseChan := make(chan *DescribeServiceMirrorResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeServiceMirror(request)
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

// DescribeServiceMirrorWithCallback invokes the eas.DescribeServiceMirror API asynchronously
func (client *Client) DescribeServiceMirrorWithCallback(request *DescribeServiceMirrorRequest, callback func(response *DescribeServiceMirrorResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeServiceMirrorResponse
		var err error
		defer close(result)
		response, err = client.DescribeServiceMirror(request)
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

// DescribeServiceMirrorRequest is the request struct for api DescribeServiceMirror
type DescribeServiceMirrorRequest struct {
	*requests.RoaRequest
	ServiceName string `position:"Path" name:"ServiceName"`
	ClusterId   string `position:"Path" name:"ClusterId"`
}

// DescribeServiceMirrorResponse is the response struct for api DescribeServiceMirror
type DescribeServiceMirrorResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	ServiceName string `json:"ServiceName" xml:"ServiceName"`
	Target      string `json:"Target" xml:"Target"`
	Ratio       string `json:"Ratio" xml:"Ratio"`
}

// CreateDescribeServiceMirrorRequest creates a request to invoke DescribeServiceMirror API
func CreateDescribeServiceMirrorRequest() (request *DescribeServiceMirrorRequest) {
	request = &DescribeServiceMirrorRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("eas", "2021-07-01", "DescribeServiceMirror", "/api/v2/services/[ClusterId]/[ServiceName]/mirror", "eas", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeServiceMirrorResponse creates a response to parse from DescribeServiceMirror response
func CreateDescribeServiceMirrorResponse() (response *DescribeServiceMirrorResponse) {
	response = &DescribeServiceMirrorResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
