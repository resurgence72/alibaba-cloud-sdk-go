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

// GetNodeInfoForPod invokes the eflo.GetNodeInfoForPod API synchronously
func (client *Client) GetNodeInfoForPod(request *GetNodeInfoForPodRequest) (response *GetNodeInfoForPodResponse, err error) {
	response = CreateGetNodeInfoForPodResponse()
	err = client.DoAction(request, response)
	return
}

// GetNodeInfoForPodWithChan invokes the eflo.GetNodeInfoForPod API asynchronously
func (client *Client) GetNodeInfoForPodWithChan(request *GetNodeInfoForPodRequest) (<-chan *GetNodeInfoForPodResponse, <-chan error) {
	responseChan := make(chan *GetNodeInfoForPodResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetNodeInfoForPod(request)
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

// GetNodeInfoForPodWithCallback invokes the eflo.GetNodeInfoForPod API asynchronously
func (client *Client) GetNodeInfoForPodWithCallback(request *GetNodeInfoForPodRequest, callback func(response *GetNodeInfoForPodResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetNodeInfoForPodResponse
		var err error
		defer close(result)
		response, err = client.GetNodeInfoForPod(request)
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

// GetNodeInfoForPodRequest is the request struct for api GetNodeInfoForPod
type GetNodeInfoForPodRequest struct {
	*requests.RpcRequest
	NodeId string `position:"Body" name:"NodeId"`
}

// GetNodeInfoForPodResponse is the response struct for api GetNodeInfoForPod
type GetNodeInfoForPodResponse struct {
	*responses.BaseResponse
	Code      int     `json:"Code" xml:"Code"`
	Message   string  `json:"Message" xml:"Message"`
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Content   Content `json:"Content" xml:"Content"`
}

// CreateGetNodeInfoForPodRequest creates a request to invoke GetNodeInfoForPod API
func CreateGetNodeInfoForPodRequest() (request *GetNodeInfoForPodRequest) {
	request = &GetNodeInfoForPodRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("eflo", "2022-05-30", "GetNodeInfoForPod", "eflo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetNodeInfoForPodResponse creates a response to parse from GetNodeInfoForPod response
func CreateGetNodeInfoForPodResponse() (response *GetNodeInfoForPodResponse) {
	response = &GetNodeInfoForPodResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
