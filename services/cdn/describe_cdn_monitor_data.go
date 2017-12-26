
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

func (client *Client) DescribeCdnMonitorData(request *DescribeCdnMonitorDataRequest) (response *DescribeCdnMonitorDataResponse, err error) {
response = CreateDescribeCdnMonitorDataResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) DescribeCdnMonitorDataWithChan(request *DescribeCdnMonitorDataRequest) (<-chan *DescribeCdnMonitorDataResponse, <-chan error) {
responseChan := make(chan *DescribeCdnMonitorDataResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DescribeCdnMonitorData(request)
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

func (client *Client) DescribeCdnMonitorDataWithCallback(request *DescribeCdnMonitorDataRequest, callback func(response *DescribeCdnMonitorDataResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DescribeCdnMonitorDataResponse
var err error
defer close(result)
response, err = client.DescribeCdnMonitorData(request)
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

type DescribeCdnMonitorDataRequest struct {
*requests.RpcRequest
                EndTime  string `position:"Query" name:"EndTime"`
                StartTime  string `position:"Query" name:"StartTime"`
                DomainName  string `position:"Query" name:"DomainName"`
                Interval  string `position:"Query" name:"Interval"`
                OwnerId  string `position:"Query" name:"OwnerId"`
                SecurityToken  string `position:"Query" name:"SecurityToken"`
}


type DescribeCdnMonitorDataResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            DomainName     string `json:"DomainName" xml:"DomainName"`
            MonitorInterval     requests.Integer `json:"MonitorInterval" xml:"MonitorInterval"`
            StartTime     string `json:"StartTime" xml:"StartTime"`
            EndTime     string `json:"EndTime" xml:"EndTime"`
                MonitorDatas struct {
                    CDNMonitorData []struct {
            TimeStamp     string `json:"TimeStamp" xml:"TimeStamp"`
            QueryPerSecond     string `json:"QueryPerSecond" xml:"QueryPerSecond"`
            BytesPerSecond     string `json:"BytesPerSecond" xml:"BytesPerSecond"`
            BytesHitRate     string `json:"BytesHitRate" xml:"BytesHitRate"`
            RequestHitRate     string `json:"RequestHitRate" xml:"RequestHitRate"`
            AverageObjectSize     string `json:"AverageObjectSize" xml:"AverageObjectSize"`
                    }   `json:"CDNMonitorData" xml:"CDNMonitorData"`
                } `json:"MonitorDatas" xml:"MonitorDatas"`
}

func CreateDescribeCdnMonitorDataRequest() (request *DescribeCdnMonitorDataRequest) {
request = &DescribeCdnMonitorDataRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeCdnMonitorData", "", "")
return
}

func CreateDescribeCdnMonitorDataResponse() (response *DescribeCdnMonitorDataResponse) {
response = &DescribeCdnMonitorDataResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

