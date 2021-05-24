package ccc

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

// ListPhoneNumbersOfSkillGroup invokes the ccc.ListPhoneNumbersOfSkillGroup API synchronously
func (client *Client) ListPhoneNumbersOfSkillGroup(request *ListPhoneNumbersOfSkillGroupRequest) (response *ListPhoneNumbersOfSkillGroupResponse, err error) {
	response = CreateListPhoneNumbersOfSkillGroupResponse()
	err = client.DoAction(request, response)
	return
}

// ListPhoneNumbersOfSkillGroupWithChan invokes the ccc.ListPhoneNumbersOfSkillGroup API asynchronously
func (client *Client) ListPhoneNumbersOfSkillGroupWithChan(request *ListPhoneNumbersOfSkillGroupRequest) (<-chan *ListPhoneNumbersOfSkillGroupResponse, <-chan error) {
	responseChan := make(chan *ListPhoneNumbersOfSkillGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListPhoneNumbersOfSkillGroup(request)
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

// ListPhoneNumbersOfSkillGroupWithCallback invokes the ccc.ListPhoneNumbersOfSkillGroup API asynchronously
func (client *Client) ListPhoneNumbersOfSkillGroupWithCallback(request *ListPhoneNumbersOfSkillGroupRequest, callback func(response *ListPhoneNumbersOfSkillGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListPhoneNumbersOfSkillGroupResponse
		var err error
		defer close(result)
		response, err = client.ListPhoneNumbersOfSkillGroup(request)
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

// ListPhoneNumbersOfSkillGroupRequest is the request struct for api ListPhoneNumbersOfSkillGroup
type ListPhoneNumbersOfSkillGroupRequest struct {
	*requests.RpcRequest
	Active        requests.Boolean `position:"Query" name:"Active"`
	PageNumber    requests.Integer `position:"Query" name:"PageNumber"`
	IsMember      requests.Boolean `position:"Query" name:"IsMember"`
	SearchPattern string           `position:"Query" name:"SearchPattern"`
	InstanceId    string           `position:"Query" name:"InstanceId"`
	SkillGroupId  string           `position:"Query" name:"SkillGroupId"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
}

// ListPhoneNumbersOfSkillGroupResponse is the response struct for api ListPhoneNumbersOfSkillGroup
type ListPhoneNumbersOfSkillGroupResponse struct {
	*responses.BaseResponse
	Code           string                             `json:"Code" xml:"Code"`
	HttpStatusCode int                                `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string                             `json:"Message" xml:"Message"`
	RequestId      string                             `json:"RequestId" xml:"RequestId"`
	Data           DataInListPhoneNumbersOfSkillGroup `json:"Data" xml:"Data"`
}

// CreateListPhoneNumbersOfSkillGroupRequest creates a request to invoke ListPhoneNumbersOfSkillGroup API
func CreateListPhoneNumbersOfSkillGroupRequest() (request *ListPhoneNumbersOfSkillGroupRequest) {
	request = &ListPhoneNumbersOfSkillGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "ListPhoneNumbersOfSkillGroup", "CCC", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListPhoneNumbersOfSkillGroupResponse creates a response to parse from ListPhoneNumbersOfSkillGroup response
func CreateListPhoneNumbersOfSkillGroupResponse() (response *ListPhoneNumbersOfSkillGroupResponse) {
	response = &ListPhoneNumbersOfSkillGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
