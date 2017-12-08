
package vpc

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

func (client *Client) ModifyBgpGroupAttribute(request *ModifyBgpGroupAttributeRequest) (response *ModifyBgpGroupAttributeResponse, err error) {
response = CreateModifyBgpGroupAttributeResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) ModifyBgpGroupAttributeWithChan(request *ModifyBgpGroupAttributeRequest) (<-chan *ModifyBgpGroupAttributeResponse, <-chan error) {
responseChan := make(chan *ModifyBgpGroupAttributeResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.ModifyBgpGroupAttribute(request)
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

func (client *Client) ModifyBgpGroupAttributeWithCallback(request *ModifyBgpGroupAttributeRequest, callback func(response *ModifyBgpGroupAttributeResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *ModifyBgpGroupAttributeResponse
var err error
defer close(result)
response, err = client.ModifyBgpGroupAttribute(request)
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

type ModifyBgpGroupAttributeRequest struct {
*requests.RpcRequest
                ClientToken  string `position:"Query" name:"ClientToken"`
                BgpGroupId  string `position:"Query" name:"BgpGroupId"`
                OwnerId  string `position:"Query" name:"OwnerId"`
                AuthKey  string `position:"Query" name:"AuthKey"`
                ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
                Description  string `position:"Query" name:"Description"`
                Name  string `position:"Query" name:"Name"`
                IsFakeAsn  string `position:"Query" name:"IsFakeAsn"`
                ResourceOwnerId  string `position:"Query" name:"ResourceOwnerId"`
                OwnerAccount  string `position:"Query" name:"OwnerAccount"`
                PeerAsn  string `position:"Query" name:"PeerAsn"`
}


type ModifyBgpGroupAttributeResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
}

func CreateModifyBgpGroupAttributeRequest() (request *ModifyBgpGroupAttributeRequest) {
request = &ModifyBgpGroupAttributeRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Vpc", "2016-04-28", "ModifyBgpGroupAttribute", "", "")
return
}

func CreateModifyBgpGroupAttributeResponse() (response *ModifyBgpGroupAttributeResponse) {
response = &ModifyBgpGroupAttributeResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

