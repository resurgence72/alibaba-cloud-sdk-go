package ga

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

// AssociateAclsWithListener invokes the ga.AssociateAclsWithListener API synchronously
func (client *Client) AssociateAclsWithListener(request *AssociateAclsWithListenerRequest) (response *AssociateAclsWithListenerResponse, err error) {
	response = CreateAssociateAclsWithListenerResponse()
	err = client.DoAction(request, response)
	return
}

// AssociateAclsWithListenerWithChan invokes the ga.AssociateAclsWithListener API asynchronously
func (client *Client) AssociateAclsWithListenerWithChan(request *AssociateAclsWithListenerRequest) (<-chan *AssociateAclsWithListenerResponse, <-chan error) {
	responseChan := make(chan *AssociateAclsWithListenerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AssociateAclsWithListener(request)
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

// AssociateAclsWithListenerWithCallback invokes the ga.AssociateAclsWithListener API asynchronously
func (client *Client) AssociateAclsWithListenerWithCallback(request *AssociateAclsWithListenerRequest, callback func(response *AssociateAclsWithListenerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AssociateAclsWithListenerResponse
		var err error
		defer close(result)
		response, err = client.AssociateAclsWithListener(request)
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

// AssociateAclsWithListenerRequest is the request struct for api AssociateAclsWithListener
type AssociateAclsWithListenerRequest struct {
	*requests.RpcRequest
	DryRun      requests.Boolean `position:"Query" name:"DryRun"`
	ClientToken string           `position:"Query" name:"ClientToken"`
	AclIds      *[]string        `position:"Query" name:"AclIds"  type:"Repeated"`
	AclType     string           `position:"Query" name:"AclType"`
	ListenerId  string           `position:"Query" name:"ListenerId"`
}

// AssociateAclsWithListenerResponse is the response struct for api AssociateAclsWithListener
type AssociateAclsWithListenerResponse struct {
	*responses.BaseResponse
	RequestId  string   `json:"RequestId" xml:"RequestId"`
	ListenerId string   `json:"ListenerId" xml:"ListenerId"`
	AclIds     []string `json:"AclIds" xml:"AclIds"`
}

// CreateAssociateAclsWithListenerRequest creates a request to invoke AssociateAclsWithListener API
func CreateAssociateAclsWithListenerRequest() (request *AssociateAclsWithListenerRequest) {
	request = &AssociateAclsWithListenerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ga", "2019-11-20", "AssociateAclsWithListener", "ga", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAssociateAclsWithListenerResponse creates a response to parse from AssociateAclsWithListener response
func CreateAssociateAclsWithListenerResponse() (response *AssociateAclsWithListenerResponse) {
	response = &AssociateAclsWithListenerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
