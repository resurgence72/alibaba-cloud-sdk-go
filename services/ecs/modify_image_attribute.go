
package ecs

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

func (client *Client) ModifyImageAttribute(request *ModifyImageAttributeRequest) (response *ModifyImageAttributeResponse, err error) {
response = CreateModifyImageAttributeResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) ModifyImageAttributeWithChan(request *ModifyImageAttributeRequest) (<-chan *ModifyImageAttributeResponse, <-chan error) {
responseChan := make(chan *ModifyImageAttributeResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.ModifyImageAttribute(request)
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

func (client *Client) ModifyImageAttributeWithCallback(request *ModifyImageAttributeRequest, callback func(response *ModifyImageAttributeResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *ModifyImageAttributeResponse
var err error
defer close(result)
response, err = client.ModifyImageAttribute(request)
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

type ModifyImageAttributeRequest struct {
*requests.RpcRequest
                ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
                Description  string `position:"Query" name:"Description"`
                ResourceOwnerId  string `position:"Query" name:"ResourceOwnerId"`
                OwnerAccount  string `position:"Query" name:"OwnerAccount"`
                OwnerId  string `position:"Query" name:"OwnerId"`
                ImageId  string `position:"Query" name:"ImageId"`
                ImageName  string `position:"Query" name:"ImageName"`
}


type ModifyImageAttributeResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
}

func CreateModifyImageAttributeRequest() (request *ModifyImageAttributeRequest) {
request = &ModifyImageAttributeRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Ecs", "2014-05-26", "ModifyImageAttribute", "", "")
return
}

func CreateModifyImageAttributeResponse() (response *ModifyImageAttributeResponse) {
response = &ModifyImageAttributeResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

