
package vpc

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

func (client *Client) DescribeEipMonitorData(request *DescribeEipMonitorDataRequest) (response *DescribeEipMonitorDataResponse, err error) {
response = CreateDescribeEipMonitorDataResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) DescribeEipMonitorDataWithChan(request *DescribeEipMonitorDataRequest) (<-chan *DescribeEipMonitorDataResponse, <-chan error) {
responseChan := make(chan *DescribeEipMonitorDataResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DescribeEipMonitorData(request)
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

func (client *Client) DescribeEipMonitorDataWithCallback(request *DescribeEipMonitorDataRequest, callback func(response *DescribeEipMonitorDataResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DescribeEipMonitorDataResponse
var err error
defer close(result)
response, err = client.DescribeEipMonitorData(request)
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

type DescribeEipMonitorDataRequest struct {
*requests.RpcRequest
                EndTime  string `position:"Query" name:"EndTime"`
                StartTime  string `position:"Query" name:"StartTime"`
                ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
                ResourceOwnerId  string `position:"Query" name:"ResourceOwnerId"`
                OwnerAccount  string `position:"Query" name:"OwnerAccount"`
                Period  string `position:"Query" name:"Period"`
                OwnerId  string `position:"Query" name:"OwnerId"`
                AllocationId  string `position:"Query" name:"AllocationId"`
}


type DescribeEipMonitorDataResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
                EipMonitorDatas struct {
                    EipMonitorData []struct {
            EipRX     requests.Integer `json:"EipRX" xml:"EipRX"`
            EipTX     requests.Integer `json:"EipTX" xml:"EipTX"`
            EipFlow     requests.Integer `json:"EipFlow" xml:"EipFlow"`
            EipBandwidth     requests.Integer `json:"EipBandwidth" xml:"EipBandwidth"`
            EipPackets     requests.Integer `json:"EipPackets" xml:"EipPackets"`
            TimeStamp     string `json:"TimeStamp" xml:"TimeStamp"`
                    }   `json:"EipMonitorData" xml:"EipMonitorData"`
                } `json:"EipMonitorDatas" xml:"EipMonitorDatas"`
}

func CreateDescribeEipMonitorDataRequest() (request *DescribeEipMonitorDataRequest) {
request = &DescribeEipMonitorDataRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Vpc", "2016-04-28", "DescribeEipMonitorData", "", "")
return
}

func CreateDescribeEipMonitorDataResponse() (response *DescribeEipMonitorDataResponse) {
response = &DescribeEipMonitorDataResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

