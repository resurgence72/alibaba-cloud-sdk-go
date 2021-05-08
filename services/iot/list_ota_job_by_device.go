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

// ListOTAJobByDevice invokes the iot.ListOTAJobByDevice API synchronously
func (client *Client) ListOTAJobByDevice(request *ListOTAJobByDeviceRequest) (response *ListOTAJobByDeviceResponse, err error) {
	response = CreateListOTAJobByDeviceResponse()
	err = client.DoAction(request, response)
	return
}

// ListOTAJobByDeviceWithChan invokes the iot.ListOTAJobByDevice API asynchronously
func (client *Client) ListOTAJobByDeviceWithChan(request *ListOTAJobByDeviceRequest) (<-chan *ListOTAJobByDeviceResponse, <-chan error) {
	responseChan := make(chan *ListOTAJobByDeviceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListOTAJobByDevice(request)
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

// ListOTAJobByDeviceWithCallback invokes the iot.ListOTAJobByDevice API asynchronously
func (client *Client) ListOTAJobByDeviceWithCallback(request *ListOTAJobByDeviceRequest, callback func(response *ListOTAJobByDeviceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListOTAJobByDeviceResponse
		var err error
		defer close(result)
		response, err = client.ListOTAJobByDevice(request)
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

// ListOTAJobByDeviceRequest is the request struct for api ListOTAJobByDevice
type ListOTAJobByDeviceRequest struct {
	*requests.RpcRequest
	IotInstanceId string           `position:"Query" name:"IotInstanceId"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	FirmwareId    string           `position:"Query" name:"FirmwareId"`
	CurrentPage   requests.Integer `position:"Query" name:"CurrentPage"`
	ProductKey    string           `position:"Query" name:"ProductKey"`
	ApiProduct    string           `position:"Body" name:"ApiProduct"`
	ApiRevision   string           `position:"Body" name:"ApiRevision"`
	DeviceName    string           `position:"Query" name:"DeviceName"`
}

// ListOTAJobByDeviceResponse is the response struct for api ListOTAJobByDevice
type ListOTAJobByDeviceResponse struct {
	*responses.BaseResponse
	RequestId    string                   `json:"RequestId" xml:"RequestId"`
	Success      bool                     `json:"Success" xml:"Success"`
	Code         string                   `json:"Code" xml:"Code"`
	ErrorMessage string                   `json:"ErrorMessage" xml:"ErrorMessage"`
	Total        int                      `json:"Total" xml:"Total"`
	PageSize     int                      `json:"PageSize" xml:"PageSize"`
	PageCount    int                      `json:"PageCount" xml:"PageCount"`
	CurrentPage  int                      `json:"CurrentPage" xml:"CurrentPage"`
	Data         DataInListOTAJobByDevice `json:"Data" xml:"Data"`
}

// CreateListOTAJobByDeviceRequest creates a request to invoke ListOTAJobByDevice API
func CreateListOTAJobByDeviceRequest() (request *ListOTAJobByDeviceRequest) {
	request = &ListOTAJobByDeviceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "ListOTAJobByDevice", "", "")
	request.Method = requests.POST
	return
}

// CreateListOTAJobByDeviceResponse creates a response to parse from ListOTAJobByDevice response
func CreateListOTAJobByDeviceResponse() (response *ListOTAJobByDeviceResponse) {
	response = &ListOTAJobByDeviceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
