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

// QueryDeviceEventData invokes the iot.QueryDeviceEventData API synchronously
func (client *Client) QueryDeviceEventData(request *QueryDeviceEventDataRequest) (response *QueryDeviceEventDataResponse, err error) {
	response = CreateQueryDeviceEventDataResponse()
	err = client.DoAction(request, response)
	return
}

// QueryDeviceEventDataWithChan invokes the iot.QueryDeviceEventData API asynchronously
func (client *Client) QueryDeviceEventDataWithChan(request *QueryDeviceEventDataRequest) (<-chan *QueryDeviceEventDataResponse, <-chan error) {
	responseChan := make(chan *QueryDeviceEventDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryDeviceEventData(request)
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

// QueryDeviceEventDataWithCallback invokes the iot.QueryDeviceEventData API asynchronously
func (client *Client) QueryDeviceEventDataWithCallback(request *QueryDeviceEventDataRequest, callback func(response *QueryDeviceEventDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryDeviceEventDataResponse
		var err error
		defer close(result)
		response, err = client.QueryDeviceEventData(request)
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

// QueryDeviceEventDataRequest is the request struct for api QueryDeviceEventData
type QueryDeviceEventDataRequest struct {
	*requests.RpcRequest
	RealTenantId      string           `position:"Query" name:"RealTenantId"`
	StartTime         requests.Integer `position:"Query" name:"StartTime"`
	RealTripartiteKey string           `position:"Query" name:"RealTripartiteKey"`
	IotId             string           `position:"Query" name:"IotId"`
	IotInstanceId     string           `position:"Query" name:"IotInstanceId"`
	PageSize          requests.Integer `position:"Query" name:"PageSize"`
	Identifier        string           `position:"Query" name:"Identifier"`
	EndTime           requests.Integer `position:"Query" name:"EndTime"`
	ProductKey        string           `position:"Query" name:"ProductKey"`
	Asc               requests.Integer `position:"Query" name:"Asc"`
	ApiProduct        string           `position:"Body" name:"ApiProduct"`
	ApiRevision       string           `position:"Body" name:"ApiRevision"`
	DeviceName        string           `position:"Query" name:"DeviceName"`
	EventType         string           `position:"Query" name:"EventType"`
}

// QueryDeviceEventDataResponse is the response struct for api QueryDeviceEventData
type QueryDeviceEventDataResponse struct {
	*responses.BaseResponse
	RequestId    string                     `json:"RequestId" xml:"RequestId"`
	Success      bool                       `json:"Success" xml:"Success"`
	Code         string                     `json:"Code" xml:"Code"`
	ErrorMessage string                     `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         DataInQueryDeviceEventData `json:"Data" xml:"Data"`
}

// CreateQueryDeviceEventDataRequest creates a request to invoke QueryDeviceEventData API
func CreateQueryDeviceEventDataRequest() (request *QueryDeviceEventDataRequest) {
	request = &QueryDeviceEventDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QueryDeviceEventData", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryDeviceEventDataResponse creates a response to parse from QueryDeviceEventData response
func CreateQueryDeviceEventDataResponse() (response *QueryDeviceEventDataResponse) {
	response = &QueryDeviceEventDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
