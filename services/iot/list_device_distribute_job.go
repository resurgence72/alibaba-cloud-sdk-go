package iot

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

// ListDeviceDistributeJob invokes the iot.ListDeviceDistributeJob API synchronously
func (client *Client) ListDeviceDistributeJob(request *ListDeviceDistributeJobRequest) (response *ListDeviceDistributeJobResponse, err error) {
	response = CreateListDeviceDistributeJobResponse()
	err = client.DoAction(request, response)
	return
}

// ListDeviceDistributeJobWithChan invokes the iot.ListDeviceDistributeJob API asynchronously
func (client *Client) ListDeviceDistributeJobWithChan(request *ListDeviceDistributeJobRequest) (<-chan *ListDeviceDistributeJobResponse, <-chan error) {
	responseChan := make(chan *ListDeviceDistributeJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDeviceDistributeJob(request)
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

// ListDeviceDistributeJobWithCallback invokes the iot.ListDeviceDistributeJob API asynchronously
func (client *Client) ListDeviceDistributeJobWithCallback(request *ListDeviceDistributeJobRequest, callback func(response *ListDeviceDistributeJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDeviceDistributeJobResponse
		var err error
		defer close(result)
		response, err = client.ListDeviceDistributeJob(request)
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

// ListDeviceDistributeJobRequest is the request struct for api ListDeviceDistributeJob
type ListDeviceDistributeJobRequest struct {
	*requests.RpcRequest
	RealTenantId      string           `position:"Body" name:"RealTenantId"`
	RealTripartiteKey string           `position:"Body" name:"RealTripartiteKey"`
	JobId             string           `position:"Body" name:"JobId"`
	NextToken         string           `position:"Query" name:"NextToken"`
	PageSize          requests.Integer `position:"Query" name:"PageSize"`
	CurrentPage       requests.Integer `position:"Query" name:"CurrentPage"`
	ProductKey        string           `position:"Query" name:"ProductKey"`
	JobType           string           `position:"Query" name:"JobType"`
	ApiProduct        string           `position:"Body" name:"ApiProduct"`
	ApiRevision       string           `position:"Body" name:"ApiRevision"`
	TargetUid         string           `position:"Query" name:"TargetUid"`
	SourceRegion      string           `position:"Query" name:"SourceRegion"`
	Status            requests.Integer `position:"Query" name:"Status"`
}

// ListDeviceDistributeJobResponse is the response struct for api ListDeviceDistributeJob
type ListDeviceDistributeJobResponse struct {
	*responses.BaseResponse
	RequestId    string                        `json:"RequestId" xml:"RequestId"`
	Success      bool                          `json:"Success" xml:"Success"`
	Code         string                        `json:"Code" xml:"Code"`
	ErrorMessage string                        `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         DataInListDeviceDistributeJob `json:"Data" xml:"Data"`
}

// CreateListDeviceDistributeJobRequest creates a request to invoke ListDeviceDistributeJob API
func CreateListDeviceDistributeJobRequest() (request *ListDeviceDistributeJobRequest) {
	request = &ListDeviceDistributeJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "ListDeviceDistributeJob", "", "")
	request.Method = requests.POST
	return
}

// CreateListDeviceDistributeJobResponse creates a response to parse from ListDeviceDistributeJob response
func CreateListDeviceDistributeJobResponse() (response *ListDeviceDistributeJobResponse) {
	response = &ListDeviceDistributeJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
