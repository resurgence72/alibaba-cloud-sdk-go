
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

func (client *Client) DescribeDomainSlowRatio(request *DescribeDomainSlowRatioRequest) (response *DescribeDomainSlowRatioResponse, err error) {
response = CreateDescribeDomainSlowRatioResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) DescribeDomainSlowRatioWithChan(request *DescribeDomainSlowRatioRequest) (<-chan *DescribeDomainSlowRatioResponse, <-chan error) {
responseChan := make(chan *DescribeDomainSlowRatioResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DescribeDomainSlowRatio(request)
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

func (client *Client) DescribeDomainSlowRatioWithCallback(request *DescribeDomainSlowRatioRequest, callback func(response *DescribeDomainSlowRatioResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DescribeDomainSlowRatioResponse
var err error
defer close(result)
response, err = client.DescribeDomainSlowRatio(request)
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

type DescribeDomainSlowRatioRequest struct {
*requests.RpcRequest
                EndTime  string `position:"Query" name:"EndTime"`
                PageSize  string `position:"Query" name:"PageSize"`
                Version  string `position:"Query" name:"Version"`
                DomainName  string `position:"Query" name:"DomainName"`
                StartTime  string `position:"Query" name:"StartTime"`
                PageNumber  string `position:"Query" name:"PageNumber"`
                OwnerId  string `position:"Query" name:"OwnerId"`
                SecurityToken  string `position:"Query" name:"SecurityToken"`
}


type DescribeDomainSlowRatioResponse struct {
*responses.BaseResponse
            EndTime     string `json:"EndTime" xml:"EndTime"`
            DataInterval     int `json:"DataInterval" xml:"DataInterval"`
            PageNumber     int `json:"PageNumber" xml:"PageNumber"`
            PageSize     int `json:"PageSize" xml:"PageSize"`
            TotalCount     int `json:"TotalCount" xml:"TotalCount"`
            StartTime     string `json:"StartTime" xml:"StartTime"`
                SlowRatioDataPerInterval struct {
                    SlowRatioData []struct {
            TotalUsers     int `json:"TotalUsers" xml:"TotalUsers"`
            SlowUsers     int `json:"SlowUsers" xml:"SlowUsers"`
            SlowRatio     float64 `json:"SlowRatio" xml:"SlowRatio"`
            RegionNameZh     string `json:"RegionNameZh" xml:"RegionNameZh"`
            RegionNameEn     string `json:"RegionNameEn" xml:"RegionNameEn"`
            IspNameZh     string `json:"IspNameZh" xml:"IspNameZh"`
            IspNameEn     string `json:"IspNameEn" xml:"IspNameEn"`
            Time     string `json:"Time" xml:"Time"`
                    }   `json:"SlowRatioData" xml:"SlowRatioData"`
                } `json:"SlowRatioDataPerInterval" xml:"SlowRatioDataPerInterval"`
}

func CreateDescribeDomainSlowRatioRequest() (request *DescribeDomainSlowRatioRequest) {
request = &DescribeDomainSlowRatioRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainSlowRatio", "", "")
return
}

func CreateDescribeDomainSlowRatioResponse() (response *DescribeDomainSlowRatioResponse) {
response = &DescribeDomainSlowRatioResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

