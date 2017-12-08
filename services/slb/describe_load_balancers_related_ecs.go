
package slb

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

func (client *Client) DescribeLoadBalancersRelatedEcs(request *DescribeLoadBalancersRelatedEcsRequest) (response *DescribeLoadBalancersRelatedEcsResponse, err error) {
response = CreateDescribeLoadBalancersRelatedEcsResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) DescribeLoadBalancersRelatedEcsWithChan(request *DescribeLoadBalancersRelatedEcsRequest) (<-chan *DescribeLoadBalancersRelatedEcsResponse, <-chan error) {
responseChan := make(chan *DescribeLoadBalancersRelatedEcsResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DescribeLoadBalancersRelatedEcs(request)
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

func (client *Client) DescribeLoadBalancersRelatedEcsWithCallback(request *DescribeLoadBalancersRelatedEcsRequest, callback func(response *DescribeLoadBalancersRelatedEcsResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DescribeLoadBalancersRelatedEcsResponse
var err error
defer close(result)
response, err = client.DescribeLoadBalancersRelatedEcs(request)
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

type DescribeLoadBalancersRelatedEcsRequest struct {
*requests.RpcRequest
                Tags  string `position:"Query" name:"Tags"`
                ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
                AccessKeyId  string `position:"Query" name:"access_key_id"`
                ResourceOwnerId  string `position:"Query" name:"ResourceOwnerId"`
                LoadBalancerId  string `position:"Query" name:"LoadBalancerId"`
                OwnerAccount  string `position:"Query" name:"OwnerAccount"`
                OwnerId  string `position:"Query" name:"OwnerId"`
}


type DescribeLoadBalancersRelatedEcsResponse struct {
*responses.BaseResponse
            Message     string `json:"Message" xml:"Message"`
            Success     bool `json:"Success" xml:"Success"`
            RequestId     string `json:"RequestId" xml:"RequestId"`
                LoadBalancers struct {
                    LoadBalancer []struct {
            LoadBalancerId     string `json:"LoadBalancerId" xml:"LoadBalancerId"`
            Count     int `json:"Count" xml:"Count"`
                MasterSlaveVServerGroups struct {
                    MasterSlaveVServerGroup []struct {
            GroupId     string `json:"GroupId" xml:"GroupId"`
            GroupName     string `json:"GroupName" xml:"GroupName"`
                BackendServers1 struct {
                    BackendServer []struct {
            VmName     string `json:"VmName" xml:"VmName"`
            Weight     int `json:"Weight" xml:"Weight"`
            Port     int `json:"Port" xml:"Port"`
            NetworkType     string `json:"NetworkType" xml:"NetworkType"`
                    }   `json:"BackendServer" xml:"BackendServer"`
                } `json:"BackendServers" xml:"BackendServers"`
                    }   `json:"MasterSlaveVServerGroup" xml:"MasterSlaveVServerGroup"`
                } `json:"MasterSlaveVServerGroups" xml:"MasterSlaveVServerGroups"`
                VServerGroups struct {
                    VServerGroup []struct {
            GroupId     string `json:"GroupId" xml:"GroupId"`
            GroupName     string `json:"GroupName" xml:"GroupName"`
                BackendServers2 struct {
                    BackendServer []struct {
            VmName     string `json:"VmName" xml:"VmName"`
            Weight     int `json:"Weight" xml:"Weight"`
            Port     int `json:"Port" xml:"Port"`
            NetworkType     string `json:"NetworkType" xml:"NetworkType"`
                    }   `json:"BackendServer" xml:"BackendServer"`
                } `json:"BackendServers" xml:"BackendServers"`
                    }   `json:"VServerGroup" xml:"VServerGroup"`
                } `json:"VServerGroups" xml:"VServerGroups"`
                    }   `json:"LoadBalancer" xml:"LoadBalancer"`
                } `json:"LoadBalancers" xml:"LoadBalancers"`
}

func CreateDescribeLoadBalancersRelatedEcsRequest() (request *DescribeLoadBalancersRelatedEcsRequest) {
request = &DescribeLoadBalancersRelatedEcsRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Slb", "2014-05-15", "DescribeLoadBalancersRelatedEcs", "", "")
return
}

func CreateDescribeLoadBalancersRelatedEcsResponse() (response *DescribeLoadBalancersRelatedEcsResponse) {
response = &DescribeLoadBalancersRelatedEcsResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

