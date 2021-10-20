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

// SubscribeTopic invokes the iot.SubscribeTopic API synchronously
func (client *Client) SubscribeTopic(request *SubscribeTopicRequest) (response *SubscribeTopicResponse, err error) {
	response = CreateSubscribeTopicResponse()
	err = client.DoAction(request, response)
	return
}

// SubscribeTopicWithChan invokes the iot.SubscribeTopic API asynchronously
func (client *Client) SubscribeTopicWithChan(request *SubscribeTopicRequest) (<-chan *SubscribeTopicResponse, <-chan error) {
	responseChan := make(chan *SubscribeTopicResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubscribeTopic(request)
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

// SubscribeTopicWithCallback invokes the iot.SubscribeTopic API asynchronously
func (client *Client) SubscribeTopicWithCallback(request *SubscribeTopicRequest, callback func(response *SubscribeTopicResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubscribeTopicResponse
		var err error
		defer close(result)
		response, err = client.SubscribeTopic(request)
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

// SubscribeTopicRequest is the request struct for api SubscribeTopic
type SubscribeTopicRequest struct {
	*requests.RpcRequest
	IotInstanceId string    `position:"Query" name:"IotInstanceId"`
	ProductKey    string    `position:"Query" name:"ProductKey"`
	ApiProduct    string    `position:"Body" name:"ApiProduct"`
	Topic         *[]string `position:"Query" name:"Topic"  type:"Repeated"`
	ApiRevision   string    `position:"Body" name:"ApiRevision"`
	DeviceName    string    `position:"Query" name:"DeviceName"`
}

// SubscribeTopicResponse is the response struct for api SubscribeTopic
type SubscribeTopicResponse struct {
	*responses.BaseResponse
	RequestId     string   `json:"RequestId" xml:"RequestId"`
	Success       bool     `json:"Success" xml:"Success"`
	Code          string   `json:"Code" xml:"Code"`
	ErrorMessage  string   `json:"ErrorMessage" xml:"ErrorMessage"`
	FailureTopics []string `json:"FailureTopics" xml:"FailureTopics"`
}

// CreateSubscribeTopicRequest creates a request to invoke SubscribeTopic API
func CreateSubscribeTopicRequest() (request *SubscribeTopicRequest) {
	request = &SubscribeTopicRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "SubscribeTopic", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSubscribeTopicResponse creates a response to parse from SubscribeTopic response
func CreateSubscribeTopicResponse() (response *SubscribeTopicResponse) {
	response = &SubscribeTopicResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
