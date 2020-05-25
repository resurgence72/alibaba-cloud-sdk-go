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

// GetDeviceVideoUrl invokes the vcs.GetDeviceVideoUrl API synchronously
// api document: https://help.aliyun.com/api/vcs/getdevicevideourl.html
func (client *Client) GetDeviceVideoUrl(request *GetDeviceVideoUrlRequest) (response *GetDeviceVideoUrlResponse, err error) {
	response = CreateGetDeviceVideoUrlResponse()
	err = client.DoAction(request, response)
	return
}

// GetDeviceVideoUrlWithChan invokes the vcs.GetDeviceVideoUrl API asynchronously
// api document: https://help.aliyun.com/api/vcs/getdevicevideourl.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetDeviceVideoUrlWithChan(request *GetDeviceVideoUrlRequest) (<-chan *GetDeviceVideoUrlResponse, <-chan error) {
	responseChan := make(chan *GetDeviceVideoUrlResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDeviceVideoUrl(request)
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

// GetDeviceVideoUrlWithCallback invokes the vcs.GetDeviceVideoUrl API asynchronously
// api document: https://help.aliyun.com/api/vcs/getdevicevideourl.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetDeviceVideoUrlWithCallback(request *GetDeviceVideoUrlRequest, callback func(response *GetDeviceVideoUrlResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDeviceVideoUrlResponse
		var err error
		defer close(result)
		response, err = client.GetDeviceVideoUrl(request)
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

// GetDeviceVideoUrlRequest is the request struct for api GetDeviceVideoUrl
type GetDeviceVideoUrlRequest struct {
	*requests.RpcRequest
	CorpId    string           `position:"Body" name:"CorpId"`
	GbId      string           `position:"Body" name:"GbId"`
	EndTime   requests.Integer `position:"Body" name:"EndTime"`
	StartTime requests.Integer `position:"Body" name:"StartTime"`
}

// GetDeviceVideoUrlResponse is the response struct for api GetDeviceVideoUrl
type GetDeviceVideoUrlResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Url       string `json:"Url" xml:"Url"`
}

// CreateGetDeviceVideoUrlRequest creates a request to invoke GetDeviceVideoUrl API
func CreateGetDeviceVideoUrlRequest() (request *GetDeviceVideoUrlRequest) {
	request = &GetDeviceVideoUrlRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vcs", "2020-05-15", "GetDeviceVideoUrl", "vcs", "openAPI")
	return
}

// CreateGetDeviceVideoUrlResponse creates a response to parse from GetDeviceVideoUrl response
func CreateGetDeviceVideoUrlResponse() (response *GetDeviceVideoUrlResponse) {
	response = &GetDeviceVideoUrlResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
