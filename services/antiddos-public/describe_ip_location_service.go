package antiddos_public

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

// DescribeIpLocationService invokes the antiddos_public.DescribeIpLocationService API synchronously
func (client *Client) DescribeIpLocationService(request *DescribeIpLocationServiceRequest) (response *DescribeIpLocationServiceResponse, err error) {
	response = CreateDescribeIpLocationServiceResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeIpLocationServiceWithChan invokes the antiddos_public.DescribeIpLocationService API asynchronously
func (client *Client) DescribeIpLocationServiceWithChan(request *DescribeIpLocationServiceRequest) (<-chan *DescribeIpLocationServiceResponse, <-chan error) {
	responseChan := make(chan *DescribeIpLocationServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeIpLocationService(request)
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

// DescribeIpLocationServiceWithCallback invokes the antiddos_public.DescribeIpLocationService API asynchronously
func (client *Client) DescribeIpLocationServiceWithCallback(request *DescribeIpLocationServiceRequest, callback func(response *DescribeIpLocationServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeIpLocationServiceResponse
		var err error
		defer close(result)
		response, err = client.DescribeIpLocationService(request)
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

// DescribeIpLocationServiceRequest is the request struct for api DescribeIpLocationService
type DescribeIpLocationServiceRequest struct {
	*requests.RpcRequest
	InternetIp       string `position:"Query" name:"InternetIp"`
	EagleEyeRpcId    string `position:"Body" name:"EagleEyeRpcId"`
	SourceIp         string `position:"Query" name:"SourceIp"`
	EagleEyeTraceId  string `position:"Body" name:"EagleEyeTraceId"`
	Lang             string `position:"Query" name:"Lang"`
	EagleEyeUserData string `position:"Body" name:"EagleEyeUserData"`
}

// DescribeIpLocationServiceResponse is the response struct for api DescribeIpLocationService
type DescribeIpLocationServiceResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Instance  Instance `json:"Instance" xml:"Instance"`
}

// CreateDescribeIpLocationServiceRequest creates a request to invoke DescribeIpLocationService API
func CreateDescribeIpLocationServiceRequest() (request *DescribeIpLocationServiceRequest) {
	request = &DescribeIpLocationServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("antiddos-public", "2017-05-18", "DescribeIpLocationService", "ddosbasic", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeIpLocationServiceResponse creates a response to parse from DescribeIpLocationService response
func CreateDescribeIpLocationServiceResponse() (response *DescribeIpLocationServiceResponse) {
	response = &DescribeIpLocationServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
