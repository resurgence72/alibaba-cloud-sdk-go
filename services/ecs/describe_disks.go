
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

func (client *Client) DescribeDisks(request *DescribeDisksRequest) (response *DescribeDisksResponse, err error) {
response = CreateDescribeDisksResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) DescribeDisksWithChan(request *DescribeDisksRequest) (<-chan *DescribeDisksResponse, <-chan error) {
responseChan := make(chan *DescribeDisksResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DescribeDisks(request)
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

func (client *Client) DescribeDisksWithCallback(request *DescribeDisksRequest, callback func(response *DescribeDisksResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DescribeDisksResponse
var err error
defer close(result)
response, err = client.DescribeDisks(request)
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

type DescribeDisksRequest struct {
*requests.RpcRequest
                EnableAutomatedSnapshotPolicy  string `position:"Query" name:"EnableAutomatedSnapshotPolicy"`
                PageSize  string `position:"Query" name:"PageSize"`
                Portable  string `position:"Query" name:"Portable"`
                ZoneId  string `position:"Query" name:"ZoneId"`
                Filter2Key  string `position:"Query" name:"Filter.2.Key"`
                DeleteWithInstance  string `position:"Query" name:"DeleteWithInstance"`
                DeleteAutoSnapshot  string `position:"Query" name:"DeleteAutoSnapshot"`
                AdditionalAttributes  *[]string `position:"Query" name:"AdditionalAttributes"  type:"Repeated"`
                Tag5Value  string `position:"Query" name:"Tag.5.Value"`
                Tag3Key  string `position:"Query" name:"Tag.3.Key"`
                Filter2Value  string `position:"Query" name:"Filter.2.Value"`
                ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
                SnapshotId  string `position:"Query" name:"SnapshotId"`
                Filter1Value  string `position:"Query" name:"Filter.1.Value"`
                Tag1Key  string `position:"Query" name:"Tag.1.Key"`
                Tag1Value  string `position:"Query" name:"Tag.1.Value"`
                DiskChargeType  string `position:"Query" name:"DiskChargeType"`
                ResourceOwnerId  string `position:"Query" name:"ResourceOwnerId"`
                EnableShared  string `position:"Query" name:"EnableShared"`
                Tag4Value  string `position:"Query" name:"Tag.4.Value"`
                OwnerAccount  string `position:"Query" name:"OwnerAccount"`
                DiskName  string `position:"Query" name:"DiskName"`
                ResourceGroupId  string `position:"Query" name:"ResourceGroupId"`
                DiskIds  string `position:"Query" name:"DiskIds"`
                Status  string `position:"Query" name:"Status"`
                PageNumber  string `position:"Query" name:"PageNumber"`
                OwnerId  string `position:"Query" name:"OwnerId"`
                Tag5Key  string `position:"Query" name:"Tag.5.Key"`
                LockReason  string `position:"Query" name:"LockReason"`
                Encrypted  string `position:"Query" name:"Encrypted"`
                Category  string `position:"Query" name:"Category"`
                DiskType  string `position:"Query" name:"DiskType"`
                EnableAutoSnapshot  string `position:"Query" name:"EnableAutoSnapshot"`
                Tag2Key  string `position:"Query" name:"Tag.2.Key"`
                Filter1Key  string `position:"Query" name:"Filter.1.Key"`
                AutoSnapshotPolicyId  string `position:"Query" name:"AutoSnapshotPolicyId"`
                Tag3Value  string `position:"Query" name:"Tag.3.Value"`
                InstanceId  string `position:"Query" name:"InstanceId"`
                Tag2Value  string `position:"Query" name:"Tag.2.Value"`
                Tag4Key  string `position:"Query" name:"Tag.4.Key"`
}


type DescribeDisksResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            TotalCount     int `json:"TotalCount" xml:"TotalCount"`
            PageNumber     int `json:"PageNumber" xml:"PageNumber"`
            PageSize     int `json:"PageSize" xml:"PageSize"`
                Disks struct {
                    Disk []struct {
            DiskId     string `json:"DiskId" xml:"DiskId"`
            RegionId     string `json:"RegionId" xml:"RegionId"`
            ZoneId     string `json:"ZoneId" xml:"ZoneId"`
            DiskName     string `json:"DiskName" xml:"DiskName"`
            Description     string `json:"Description" xml:"Description"`
            Type     string `json:"Type" xml:"Type"`
            Category     string `json:"Category" xml:"Category"`
            Size     int `json:"Size" xml:"Size"`
            ImageId     string `json:"ImageId" xml:"ImageId"`
            SourceSnapshotId     string `json:"SourceSnapshotId" xml:"SourceSnapshotId"`
            AutoSnapshotPolicyId     string `json:"AutoSnapshotPolicyId" xml:"AutoSnapshotPolicyId"`
            ProductCode     string `json:"ProductCode" xml:"ProductCode"`
            Portable     bool `json:"Portable" xml:"Portable"`
            Status     string `json:"Status" xml:"Status"`
            InstanceId     string `json:"InstanceId" xml:"InstanceId"`
            Device     string `json:"Device" xml:"Device"`
            DeleteWithInstance     bool `json:"DeleteWithInstance" xml:"DeleteWithInstance"`
            DeleteAutoSnapshot     bool `json:"DeleteAutoSnapshot" xml:"DeleteAutoSnapshot"`
            EnableAutoSnapshot     bool `json:"EnableAutoSnapshot" xml:"EnableAutoSnapshot"`
            EnableAutomatedSnapshotPolicy     bool `json:"EnableAutomatedSnapshotPolicy" xml:"EnableAutomatedSnapshotPolicy"`
            CreationTime     string `json:"CreationTime" xml:"CreationTime"`
            AttachedTime     string `json:"AttachedTime" xml:"AttachedTime"`
            DetachedTime     string `json:"DetachedTime" xml:"DetachedTime"`
            DiskChargeType     string `json:"DiskChargeType" xml:"DiskChargeType"`
            ExpiredTime     string `json:"ExpiredTime" xml:"ExpiredTime"`
            ResourceGroupId     string `json:"ResourceGroupId" xml:"ResourceGroupId"`
            Encrypted     bool `json:"Encrypted" xml:"Encrypted"`
            IOPS     int `json:"IOPS" xml:"IOPS"`
            IOPSRead     int `json:"IOPSRead" xml:"IOPSRead"`
            IOPSWrite     int `json:"IOPSWrite" xml:"IOPSWrite"`
                OperationLocks struct {
                    OperationLock []struct {
            LockReason     string `json:"LockReason" xml:"LockReason"`
                    }   `json:"OperationLock" xml:"OperationLock"`
                } `json:"OperationLocks" xml:"OperationLocks"`
                Tags struct {
                    Tag []struct {
            TagKey     string `json:"TagKey" xml:"TagKey"`
            TagValue     string `json:"TagValue" xml:"TagValue"`
                    }   `json:"Tag" xml:"Tag"`
                } `json:"Tags" xml:"Tags"`
                    }   `json:"Disk" xml:"Disk"`
                } `json:"Disks" xml:"Disks"`
}

func CreateDescribeDisksRequest() (request *DescribeDisksRequest) {
request = &DescribeDisksRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeDisks", "", "")
return
}

func CreateDescribeDisksResponse() (response *DescribeDisksResponse) {
response = &DescribeDisksResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

