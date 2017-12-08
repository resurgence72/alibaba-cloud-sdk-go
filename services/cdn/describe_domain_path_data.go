
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

func (client *Client) DescribeDomainPathData(request *DescribeDomainPathDataRequest) (response *DescribeDomainPathDataResponse, err error) {
response = CreateDescribeDomainPathDataResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) DescribeDomainPathDataWithChan(request *DescribeDomainPathDataRequest) (<-chan *DescribeDomainPathDataResponse, <-chan error) {
responseChan := make(chan *DescribeDomainPathDataResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DescribeDomainPathData(request)
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

func (client *Client) DescribeDomainPathDataWithCallback(request *DescribeDomainPathDataRequest, callback func(response *DescribeDomainPathDataResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DescribeDomainPathDataResponse
var err error
defer close(result)
response, err = client.DescribeDomainPathData(request)
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

type DescribeDomainPathDataRequest struct {
*requests.RpcRequest
                EndTime  string `position:"Query" name:"EndTime"`
                PageSize  string `position:"Query" name:"PageSize"`
                Version  string `position:"Query" name:"Version"`
                DomainName  string `position:"Query" name:"DomainName"`
                Path  string `position:"Query" name:"Path"`
                StartTime  string `position:"Query" name:"StartTime"`
                PageNumber  string `position:"Query" name:"PageNumber"`
                OwnerId  string `position:"Query" name:"OwnerId"`
                SecurityToken  string `position:"Query" name:"SecurityToken"`
}


type DescribeDomainPathDataResponse struct {
*responses.BaseResponse
            DomainName     string `json:"DomainName" xml:"DomainName"`
            StartTime     string `json:"StartTime" xml:"StartTime"`
            EndTime     string `json:"EndTime" xml:"EndTime"`
            PageSize     int `json:"PageSize" xml:"PageSize"`
            PageNumber     int `json:"PageNumber" xml:"PageNumber"`
            DataInterval     string `json:"DataInterval" xml:"DataInterval"`
            TotalCount     int `json:"TotalCount" xml:"TotalCount"`
                PathDataPerInterval struct {
                    UsageData []struct {
            Traffic     int `json:"Traffic" xml:"Traffic"`
            Acc     int `json:"Acc" xml:"Acc"`
            Path     string `json:"Path" xml:"Path"`
            Time     string `json:"Time" xml:"Time"`
                    }   `json:"UsageData" xml:"UsageData"`
                } `json:"PathDataPerInterval" xml:"PathDataPerInterval"`
}

func CreateDescribeDomainPathDataRequest() (request *DescribeDomainPathDataRequest) {
request = &DescribeDomainPathDataRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainPathData", "", "")
return
}

func CreateDescribeDomainPathDataResponse() (response *DescribeDomainPathDataResponse) {
response = &DescribeDomainPathDataResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

