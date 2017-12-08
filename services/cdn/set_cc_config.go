
package cdn

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

func (client *Client) SetCcConfig(request *SetCcConfigRequest) (response *SetCcConfigResponse, err error) {
response = CreateSetCcConfigResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) SetCcConfigWithChan(request *SetCcConfigRequest) (<-chan *SetCcConfigResponse, <-chan error) {
responseChan := make(chan *SetCcConfigResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.SetCcConfig(request)
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

func (client *Client) SetCcConfigWithCallback(request *SetCcConfigRequest, callback func(response *SetCcConfigResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *SetCcConfigResponse
var err error
defer close(result)
response, err = client.SetCcConfig(request)
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

type SetCcConfigRequest struct {
*requests.RpcRequest
                DomainName  string `position:"Query" name:"DomainName"`
                AllowIps  string `position:"Query" name:"AllowIps"`
                BlockIps  string `position:"Query" name:"BlockIps"`
                OwnerId  string `position:"Query" name:"OwnerId"`
                SecurityToken  string `position:"Query" name:"SecurityToken"`
}


type SetCcConfigResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
}

func CreateSetCcConfigRequest() (request *SetCcConfigRequest) {
request = &SetCcConfigRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Cdn", "2014-11-11", "SetCcConfig", "", "")
return
}

func CreateSetCcConfigResponse() (response *SetCcConfigResponse) {
response = &SetCcConfigResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

