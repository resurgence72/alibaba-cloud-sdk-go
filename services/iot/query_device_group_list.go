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

// QueryDeviceGroupList invokes the iot.QueryDeviceGroupList API synchronously
func (client *Client) QueryDeviceGroupList(request *QueryDeviceGroupListRequest) (response *QueryDeviceGroupListResponse, err error) {
	response = CreateQueryDeviceGroupListResponse()
	err = client.DoAction(request, response)
	return
}

// QueryDeviceGroupListWithChan invokes the iot.QueryDeviceGroupList API asynchronously
func (client *Client) QueryDeviceGroupListWithChan(request *QueryDeviceGroupListRequest) (<-chan *QueryDeviceGroupListResponse, <-chan error) {
	responseChan := make(chan *QueryDeviceGroupListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryDeviceGroupList(request)
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

// QueryDeviceGroupListWithCallback invokes the iot.QueryDeviceGroupList API asynchronously
func (client *Client) QueryDeviceGroupListWithCallback(request *QueryDeviceGroupListRequest, callback func(response *QueryDeviceGroupListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryDeviceGroupListResponse
		var err error
		defer close(result)
		response, err = client.QueryDeviceGroupList(request)
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

// QueryDeviceGroupListRequest is the request struct for api QueryDeviceGroupList
type QueryDeviceGroupListRequest struct {
	*requests.RpcRequest
	GroupTypes    *[]string        `position:"Query" name:"GroupTypes"  type:"Repeated"`
	SuperGroupId  string           `position:"Query" name:"SuperGroupId"`
	IotInstanceId string           `position:"Query" name:"IotInstanceId"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	CurrentPage   requests.Integer `position:"Query" name:"CurrentPage"`
	GroupName     string           `position:"Query" name:"GroupName"`
	ApiProduct    string           `position:"Body" name:"ApiProduct"`
	ApiRevision   string           `position:"Body" name:"ApiRevision"`
}

// QueryDeviceGroupListResponse is the response struct for api QueryDeviceGroupList
type QueryDeviceGroupListResponse struct {
	*responses.BaseResponse
	RequestId    string                     `json:"RequestId" xml:"RequestId"`
	Success      bool                       `json:"Success" xml:"Success"`
	Code         string                     `json:"Code" xml:"Code"`
	ErrorMessage string                     `json:"ErrorMessage" xml:"ErrorMessage"`
	CurrentPage  int                        `json:"CurrentPage" xml:"CurrentPage"`
	PageCount    int                        `json:"PageCount" xml:"PageCount"`
	PageSize     int                        `json:"PageSize" xml:"PageSize"`
	Total        int                        `json:"Total" xml:"Total"`
	Data         DataInQueryDeviceGroupList `json:"Data" xml:"Data"`
}

// CreateQueryDeviceGroupListRequest creates a request to invoke QueryDeviceGroupList API
func CreateQueryDeviceGroupListRequest() (request *QueryDeviceGroupListRequest) {
	request = &QueryDeviceGroupListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QueryDeviceGroupList", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryDeviceGroupListResponse creates a response to parse from QueryDeviceGroupList response
func CreateQueryDeviceGroupListResponse() (response *QueryDeviceGroupListResponse) {
	response = &QueryDeviceGroupListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
