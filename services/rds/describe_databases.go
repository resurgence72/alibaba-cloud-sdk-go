
package rds

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

func (client *Client) DescribeDatabases(request *DescribeDatabasesRequest) (response *DescribeDatabasesResponse, err error) {
response = CreateDescribeDatabasesResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) DescribeDatabasesWithChan(request *DescribeDatabasesRequest) (<-chan *DescribeDatabasesResponse, <-chan error) {
responseChan := make(chan *DescribeDatabasesResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DescribeDatabases(request)
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

func (client *Client) DescribeDatabasesWithCallback(request *DescribeDatabasesRequest, callback func(response *DescribeDatabasesResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DescribeDatabasesResponse
var err error
defer close(result)
response, err = client.DescribeDatabases(request)
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

type DescribeDatabasesRequest struct {
*requests.RpcRequest
                DBInstanceId  string `position:"Query" name:"DBInstanceId"`
                ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
                DBName  string `position:"Query" name:"DBName"`
                ResourceOwnerId  string `position:"Query" name:"ResourceOwnerId"`
                OwnerAccount  string `position:"Query" name:"OwnerAccount"`
                OwnerId  string `position:"Query" name:"OwnerId"`
                DBStatus  string `position:"Query" name:"DBStatus"`
}


type DescribeDatabasesResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
                Databases struct {
                    Database []struct {
            DBName     string `json:"DBName" xml:"DBName"`
            DBInstanceId     string `json:"DBInstanceId" xml:"DBInstanceId"`
            Engine     string `json:"Engine" xml:"Engine"`
            DBStatus     string `json:"DBStatus" xml:"DBStatus"`
            CharacterSetName     string `json:"CharacterSetName" xml:"CharacterSetName"`
            DBDescription     string `json:"DBDescription" xml:"DBDescription"`
                Accounts struct {
                    AccountPrivilegeInfo []struct {
            Account     string `json:"Account" xml:"Account"`
            AccountPrivilege     string `json:"AccountPrivilege" xml:"AccountPrivilege"`
                    }   `json:"AccountPrivilegeInfo" xml:"AccountPrivilegeInfo"`
                } `json:"Accounts" xml:"Accounts"`
                    }   `json:"Database" xml:"Database"`
                } `json:"Databases" xml:"Databases"`
}

func CreateDescribeDatabasesRequest() (request *DescribeDatabasesRequest) {
request = &DescribeDatabasesRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2014-08-15", "DescribeDatabases", "", "")
return
}

func CreateDescribeDatabasesResponse() (response *DescribeDatabasesResponse) {
response = &DescribeDatabasesResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

