package cloudphoto

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

func (client *Client) TagPhoto(request *TagPhotoRequest) (response *TagPhotoResponse, err error) {
	response = CreateTagPhotoResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) TagPhotoWithChan(request *TagPhotoRequest) (<-chan *TagPhotoResponse, <-chan error) {
	responseChan := make(chan *TagPhotoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TagPhoto(request)
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

func (client *Client) TagPhotoWithCallback(request *TagPhotoRequest, callback func(response *TagPhotoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TagPhotoResponse
		var err error
		defer close(result)
		response, err = client.TagPhoto(request)
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

type TagPhotoRequest struct {
	*requests.RpcRequest
	LibraryId  string           `position:"Query" name:"LibraryId"`
	Confidence *[]string        `position:"Query" name:"Confidence"  type:"Repeated"`
	StoreName  string           `position:"Query" name:"StoreName"`
	PhotoId    requests.Integer `position:"Query" name:"PhotoId"`
	TagKey     *[]string        `position:"Query" name:"TagKey"  type:"Repeated"`
}

type TagPhotoResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Action    string `json:"Action" xml:"Action"`
}

func CreateTagPhotoRequest() (request *TagPhotoRequest) {
	request = &TagPhotoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudPhoto", "2017-07-11", "TagPhoto", "", "")
	return
}

func CreateTagPhotoResponse() (response *TagPhotoResponse) {
	response = &TagPhotoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
