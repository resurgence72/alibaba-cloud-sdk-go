
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

func (client *Client) DescribeDBInstanceAttribute(request *DescribeDBInstanceAttributeRequest) (response *DescribeDBInstanceAttributeResponse, err error) {
response = CreateDescribeDBInstanceAttributeResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) DescribeDBInstanceAttributeWithChan(request *DescribeDBInstanceAttributeRequest) (<-chan *DescribeDBInstanceAttributeResponse, <-chan error) {
responseChan := make(chan *DescribeDBInstanceAttributeResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DescribeDBInstanceAttribute(request)
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

func (client *Client) DescribeDBInstanceAttributeWithCallback(request *DescribeDBInstanceAttributeRequest, callback func(response *DescribeDBInstanceAttributeResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DescribeDBInstanceAttributeResponse
var err error
defer close(result)
response, err = client.DescribeDBInstanceAttribute(request)
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

type DescribeDBInstanceAttributeRequest struct {
*requests.RpcRequest
                DBInstanceId  string `position:"Query" name:"DBInstanceId"`
                ResourceGroupId  string `position:"Query" name:"ResourceGroupId"`
                ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
                ResourceOwnerId  string `position:"Query" name:"ResourceOwnerId"`
                OwnerAccount  string `position:"Query" name:"OwnerAccount"`
                OwnerId  string `position:"Query" name:"OwnerId"`
}


type DescribeDBInstanceAttributeResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
                Items struct {
                    DBInstanceAttribute []struct {
            DBInstanceDiskUsed     string `json:"DBInstanceDiskUsed" xml:"DBInstanceDiskUsed"`
            GuardDBInstanceName     string `json:"GuardDBInstanceName" xml:"GuardDBInstanceName"`
            CanTempUpgrade     bool `json:"CanTempUpgrade" xml:"CanTempUpgrade"`
            TempUpgradeTimeStart     string `json:"TempUpgradeTimeStart" xml:"TempUpgradeTimeStart"`
            TempUpgradeTimeEnd     string `json:"TempUpgradeTimeEnd" xml:"TempUpgradeTimeEnd"`
            TempUpgradeRecoveryTime     string `json:"TempUpgradeRecoveryTime" xml:"TempUpgradeRecoveryTime"`
            TempUpgradeRecoveryClass     string `json:"TempUpgradeRecoveryClass" xml:"TempUpgradeRecoveryClass"`
            TempUpgradeRecoveryCpu     int `json:"TempUpgradeRecoveryCpu" xml:"TempUpgradeRecoveryCpu"`
            TempUpgradeRecoveryMemory     int `json:"TempUpgradeRecoveryMemory" xml:"TempUpgradeRecoveryMemory"`
            TempUpgradeRecoveryMaxIOPS     string `json:"TempUpgradeRecoveryMaxIOPS" xml:"TempUpgradeRecoveryMaxIOPS"`
            TempUpgradeRecoveryMaxConnections     string `json:"TempUpgradeRecoveryMaxConnections" xml:"TempUpgradeRecoveryMaxConnections"`
            InsId     int `json:"InsId" xml:"InsId"`
            DBInstanceId     string `json:"DBInstanceId" xml:"DBInstanceId"`
            PayType     string `json:"PayType" xml:"PayType"`
            DBInstanceClassType     string `json:"DBInstanceClassType" xml:"DBInstanceClassType"`
            DBInstanceType     string `json:"DBInstanceType" xml:"DBInstanceType"`
            RegionId     string `json:"RegionId" xml:"RegionId"`
            ConnectionString     string `json:"ConnectionString" xml:"ConnectionString"`
            Port     string `json:"Port" xml:"Port"`
            Engine     string `json:"Engine" xml:"Engine"`
            EngineVersion     string `json:"EngineVersion" xml:"EngineVersion"`
            DBInstanceClass     string `json:"DBInstanceClass" xml:"DBInstanceClass"`
            DBInstanceMemory     int64 `json:"DBInstanceMemory" xml:"DBInstanceMemory"`
            DBInstanceStorage     int `json:"DBInstanceStorage" xml:"DBInstanceStorage"`
            DBInstanceNetType     string `json:"DBInstanceNetType" xml:"DBInstanceNetType"`
            DBInstanceStatus     string `json:"DBInstanceStatus" xml:"DBInstanceStatus"`
            DBInstanceDescription     string `json:"DBInstanceDescription" xml:"DBInstanceDescription"`
            LockMode     string `json:"LockMode" xml:"LockMode"`
            LockReason     string `json:"LockReason" xml:"LockReason"`
            ReadDelayTime     string `json:"ReadDelayTime" xml:"ReadDelayTime"`
            DBMaxQuantity     int `json:"DBMaxQuantity" xml:"DBMaxQuantity"`
            AccountMaxQuantity     int `json:"AccountMaxQuantity" xml:"AccountMaxQuantity"`
            CreationTime     string `json:"CreationTime" xml:"CreationTime"`
            ExpireTime     string `json:"ExpireTime" xml:"ExpireTime"`
            MaintainTime     string `json:"MaintainTime" xml:"MaintainTime"`
            AvailabilityValue     string `json:"AvailabilityValue" xml:"AvailabilityValue"`
            MaxIOPS     int `json:"MaxIOPS" xml:"MaxIOPS"`
            MaxConnections     int `json:"MaxConnections" xml:"MaxConnections"`
            MasterInstanceId     string `json:"MasterInstanceId" xml:"MasterInstanceId"`
            DBInstanceCPU     string `json:"DBInstanceCPU" xml:"DBInstanceCPU"`
            IncrementSourceDBInstanceId     string `json:"IncrementSourceDBInstanceId" xml:"IncrementSourceDBInstanceId"`
            GuardDBInstanceId     string `json:"GuardDBInstanceId" xml:"GuardDBInstanceId"`
            TempDBInstanceId     string `json:"TempDBInstanceId" xml:"TempDBInstanceId"`
            SecurityIPList     string `json:"SecurityIPList" xml:"SecurityIPList"`
            ZoneId     string `json:"ZoneId" xml:"ZoneId"`
            InstanceNetworkType     string `json:"InstanceNetworkType" xml:"InstanceNetworkType"`
            Category     string `json:"Category" xml:"Category"`
            AccountType     string `json:"AccountType" xml:"AccountType"`
            SupportUpgradeAccountType     string `json:"SupportUpgradeAccountType" xml:"SupportUpgradeAccountType"`
            VpcId     string `json:"VpcId" xml:"VpcId"`
            VSwitchId     string `json:"VSwitchId" xml:"VSwitchId"`
            ConnectionMode     string `json:"ConnectionMode" xml:"ConnectionMode"`
            ResourceGroupId     string `json:"ResourceGroupId" xml:"ResourceGroupId"`
                ReadOnlyDBInstanceIds struct {
                    ReadOnlyDBInstanceId []struct {
            DBInstanceId     string `json:"DBInstanceId" xml:"DBInstanceId"`
                    }   `json:"ReadOnlyDBInstanceId" xml:"ReadOnlyDBInstanceId"`
                } `json:"ReadOnlyDBInstanceIds" xml:"ReadOnlyDBInstanceIds"`
                    }   `json:"DBInstanceAttribute" xml:"DBInstanceAttribute"`
                } `json:"Items" xml:"Items"`
}

func CreateDescribeDBInstanceAttributeRequest() (request *DescribeDBInstanceAttributeRequest) {
request = &DescribeDBInstanceAttributeRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2014-08-15", "DescribeDBInstanceAttribute", "", "")
return
}

func CreateDescribeDBInstanceAttributeResponse() (response *DescribeDBInstanceAttributeResponse) {
response = &DescribeDBInstanceAttributeResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

