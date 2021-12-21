package config

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

// DetachConfigRuleToCompliancePack invokes the config.DetachConfigRuleToCompliancePack API synchronously
func (client *Client) DetachConfigRuleToCompliancePack(request *DetachConfigRuleToCompliancePackRequest) (response *DetachConfigRuleToCompliancePackResponse, err error) {
	response = CreateDetachConfigRuleToCompliancePackResponse()
	err = client.DoAction(request, response)
	return
}

// DetachConfigRuleToCompliancePackWithChan invokes the config.DetachConfigRuleToCompliancePack API asynchronously
func (client *Client) DetachConfigRuleToCompliancePackWithChan(request *DetachConfigRuleToCompliancePackRequest) (<-chan *DetachConfigRuleToCompliancePackResponse, <-chan error) {
	responseChan := make(chan *DetachConfigRuleToCompliancePackResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DetachConfigRuleToCompliancePack(request)
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

// DetachConfigRuleToCompliancePackWithCallback invokes the config.DetachConfigRuleToCompliancePack API asynchronously
func (client *Client) DetachConfigRuleToCompliancePackWithCallback(request *DetachConfigRuleToCompliancePackRequest, callback func(response *DetachConfigRuleToCompliancePackResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DetachConfigRuleToCompliancePackResponse
		var err error
		defer close(result)
		response, err = client.DetachConfigRuleToCompliancePack(request)
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

// DetachConfigRuleToCompliancePackRequest is the request struct for api DetachConfigRuleToCompliancePack
type DetachConfigRuleToCompliancePackRequest struct {
	*requests.RpcRequest
	ConfigRuleIds    string `position:"Query" name:"ConfigRuleIds"`
	CompliancePackId string `position:"Query" name:"CompliancePackId"`
}

// DetachConfigRuleToCompliancePackResponse is the response struct for api DetachConfigRuleToCompliancePack
type DetachConfigRuleToCompliancePackResponse struct {
	*responses.BaseResponse
	RequestId         string            `json:"RequestId" xml:"RequestId"`
	OperateRuleResult OperateRuleResult `json:"OperateRuleResult" xml:"OperateRuleResult"`
}

// CreateDetachConfigRuleToCompliancePackRequest creates a request to invoke DetachConfigRuleToCompliancePack API
func CreateDetachConfigRuleToCompliancePackRequest() (request *DetachConfigRuleToCompliancePackRequest) {
	request = &DetachConfigRuleToCompliancePackRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2020-09-07", "DetachConfigRuleToCompliancePack", "", "")
	request.Method = requests.POST
	return
}

// CreateDetachConfigRuleToCompliancePackResponse creates a response to parse from DetachConfigRuleToCompliancePack response
func CreateDetachConfigRuleToCompliancePackResponse() (response *DetachConfigRuleToCompliancePackResponse) {
	response = &DetachConfigRuleToCompliancePackResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
