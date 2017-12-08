
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

func (client *Client) DeleteVolume(request *DeleteVolumeRequest) (response *DeleteVolumeResponse, err error) {
response = CreateDeleteVolumeResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) DeleteVolumeWithChan(request *DeleteVolumeRequest) (<-chan *DeleteVolumeResponse, <-chan error) {
responseChan := make(chan *DeleteVolumeResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DeleteVolume(request)
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

func (client *Client) DeleteVolumeWithCallback(request *DeleteVolumeRequest, callback func(response *DeleteVolumeResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DeleteVolumeResponse
var err error
defer close(result)
response, err = client.DeleteVolume(request)
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

type DeleteVolumeRequest struct {
*requests.RpcRequest
                ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
                VolumeId  string `position:"Query" name:"VolumeId"`
                ResourceOwnerId  string `position:"Query" name:"ResourceOwnerId"`
                OwnerAccount  string `position:"Query" name:"OwnerAccount"`
                OwnerId  string `position:"Query" name:"OwnerId"`
}


type DeleteVolumeResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
}

func CreateDeleteVolumeRequest() (request *DeleteVolumeRequest) {
request = &DeleteVolumeRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Ecs", "2014-05-26", "DeleteVolume", "", "")
return
}

func CreateDeleteVolumeResponse() (response *DeleteVolumeResponse) {
response = &DeleteVolumeResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

