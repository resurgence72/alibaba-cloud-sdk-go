package devops_rdc

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

// DeleteDevopsProjectSprint invokes the devops_rdc.DeleteDevopsProjectSprint API synchronously
func (client *Client) DeleteDevopsProjectSprint(request *DeleteDevopsProjectSprintRequest) (response *DeleteDevopsProjectSprintResponse, err error) {
	response = CreateDeleteDevopsProjectSprintResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteDevopsProjectSprintWithChan invokes the devops_rdc.DeleteDevopsProjectSprint API asynchronously
func (client *Client) DeleteDevopsProjectSprintWithChan(request *DeleteDevopsProjectSprintRequest) (<-chan *DeleteDevopsProjectSprintResponse, <-chan error) {
	responseChan := make(chan *DeleteDevopsProjectSprintResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteDevopsProjectSprint(request)
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

// DeleteDevopsProjectSprintWithCallback invokes the devops_rdc.DeleteDevopsProjectSprint API asynchronously
func (client *Client) DeleteDevopsProjectSprintWithCallback(request *DeleteDevopsProjectSprintRequest, callback func(response *DeleteDevopsProjectSprintResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteDevopsProjectSprintResponse
		var err error
		defer close(result)
		response, err = client.DeleteDevopsProjectSprint(request)
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

// DeleteDevopsProjectSprintRequest is the request struct for api DeleteDevopsProjectSprint
type DeleteDevopsProjectSprintRequest struct {
	*requests.RpcRequest
	SprintId string `position:"Body" name:"SprintId"`
	OrgId    string `position:"Body" name:"OrgId"`
}

// DeleteDevopsProjectSprintResponse is the response struct for api DeleteDevopsProjectSprint
type DeleteDevopsProjectSprintResponse struct {
	*responses.BaseResponse
	Successful bool   `json:"Successful" xml:"Successful"`
	ErrorCode  string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg   string `json:"ErrorMsg" xml:"ErrorMsg"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
	Object     bool   `json:"Object" xml:"Object"`
}

// CreateDeleteDevopsProjectSprintRequest creates a request to invoke DeleteDevopsProjectSprint API
func CreateDeleteDevopsProjectSprintRequest() (request *DeleteDevopsProjectSprintRequest) {
	request = &DeleteDevopsProjectSprintRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("devops-rdc", "2020-03-03", "DeleteDevopsProjectSprint", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteDevopsProjectSprintResponse creates a response to parse from DeleteDevopsProjectSprint response
func CreateDeleteDevopsProjectSprintResponse() (response *DeleteDevopsProjectSprintResponse) {
	response = &DeleteDevopsProjectSprintResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
