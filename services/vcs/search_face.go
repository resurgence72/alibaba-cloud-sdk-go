package vcs

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

// SearchFace invokes the vcs.SearchFace API synchronously
// api document: https://help.aliyun.com/api/vcs/searchface.html
func (client *Client) SearchFace(request *SearchFaceRequest) (response *SearchFaceResponse, err error) {
	response = CreateSearchFaceResponse()
	err = client.DoAction(request, response)
	return
}

// SearchFaceWithChan invokes the vcs.SearchFace API asynchronously
// api document: https://help.aliyun.com/api/vcs/searchface.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SearchFaceWithChan(request *SearchFaceRequest) (<-chan *SearchFaceResponse, <-chan error) {
	responseChan := make(chan *SearchFaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SearchFace(request)
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

// SearchFaceWithCallback invokes the vcs.SearchFace API asynchronously
// api document: https://help.aliyun.com/api/vcs/searchface.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SearchFaceWithCallback(request *SearchFaceRequest, callback func(response *SearchFaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SearchFaceResponse
		var err error
		defer close(result)
		response, err = client.SearchFace(request)
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

// SearchFaceRequest is the request struct for api SearchFace
type SearchFaceRequest struct {
	*requests.RpcRequest
	CorpId         string           `position:"Body" name:"CorpId"`
	GbId           string           `position:"Body" name:"GbId"`
	StartTimeStamp requests.Integer `position:"Body" name:"StartTimeStamp"`
	EndTimeStamp   requests.Integer `position:"Body" name:"EndTimeStamp"`
	PageNo         requests.Integer `position:"Body" name:"PageNo"`
	PageSize       requests.Integer `position:"Body" name:"PageSize"`
	OptionList     string           `position:"Body" name:"OptionList"`
}

// SearchFaceResponse is the response struct for api SearchFace
type SearchFaceResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateSearchFaceRequest creates a request to invoke SearchFace API
func CreateSearchFaceRequest() (request *SearchFaceRequest) {
	request = &SearchFaceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vcs", "2020-05-15", "SearchFace", "vcs", "openAPI")
	return
}

// CreateSearchFaceResponse creates a response to parse from SearchFace response
func CreateSearchFaceResponse() (response *SearchFaceResponse) {
	response = &SearchFaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
