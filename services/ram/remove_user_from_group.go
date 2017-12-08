
package ram

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

func (client *Client) RemoveUserFromGroup(request *RemoveUserFromGroupRequest) (response *RemoveUserFromGroupResponse, err error) {
response = CreateRemoveUserFromGroupResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) RemoveUserFromGroupWithChan(request *RemoveUserFromGroupRequest) (<-chan *RemoveUserFromGroupResponse, <-chan error) {
responseChan := make(chan *RemoveUserFromGroupResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.RemoveUserFromGroup(request)
responseChan <- response
errChan <- err
})
if err != nil {
errChan <- err
close(responseChan)
close(errChan)
}
return responseChan, errChan
}

func (client *Client) RemoveUserFromGroupWithCallback(request *RemoveUserFromGroupRequest, callback func(response *RemoveUserFromGroupResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *RemoveUserFromGroupResponse
var err error
defer close(result)
response, err = client.RemoveUserFromGroup(request)
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

type RemoveUserFromGroupRequest struct {
*requests.RpcRequest
                UserName  string `position:"Query" name:"UserName"`
                GroupName  string `position:"Query" name:"GroupName"`
}


type RemoveUserFromGroupResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
}

func CreateRemoveUserFromGroupRequest() (request *RemoveUserFromGroupRequest) {
request = &RemoveUserFromGroupRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Ram", "2015-05-01", "RemoveUserFromGroup", "", "")
return
}

func CreateRemoveUserFromGroupResponse() (response *RemoveUserFromGroupResponse) {
response = &RemoveUserFromGroupResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

