
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

func (client *Client) DescribeRecommendInstanceType(request *DescribeRecommendInstanceTypeRequest) (response *DescribeRecommendInstanceTypeResponse, err error) {
response = CreateDescribeRecommendInstanceTypeResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) DescribeRecommendInstanceTypeWithChan(request *DescribeRecommendInstanceTypeRequest) (<-chan *DescribeRecommendInstanceTypeResponse, <-chan error) {
responseChan := make(chan *DescribeRecommendInstanceTypeResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DescribeRecommendInstanceType(request)
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

func (client *Client) DescribeRecommendInstanceTypeWithCallback(request *DescribeRecommendInstanceTypeRequest, callback func(response *DescribeRecommendInstanceTypeResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DescribeRecommendInstanceTypeResponse
var err error
defer close(result)
response, err = client.DescribeRecommendInstanceType(request)
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

type DescribeRecommendInstanceTypeRequest struct {
*requests.RpcRequest
                ProxyId  string `position:"Query" name:"proxyId"`
                OwnerId  string `position:"Query" name:"OwnerId"`
                NetworkType  string `position:"Query" name:"NetworkType"`
                Operator  string `position:"Query" name:"operator"`
                InstanceType  string `position:"Query" name:"InstanceType"`
                ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
                Token  string `position:"Query" name:"token"`
                Scene  string `position:"Query" name:"Scene"`
                ResourceOwnerId  string `position:"Query" name:"ResourceOwnerId"`
                OwnerAccount  string `position:"Query" name:"OwnerAccount"`
                Channel  string `position:"Query" name:"channel"`
}


type DescribeRecommendInstanceTypeResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
                Data struct {
                    RecommendInstanceType []struct {
            RegionNo     string `json:"RegionNo" xml:"RegionNo"`
            CommodityCode     string `json:"CommodityCode" xml:"CommodityCode"`
            Scene     string `json:"Scene" xml:"Scene"`
            InstanceType struct {
            Generation     string `json:"Generation" xml:"Generation"`
            InstanceTypeFamily     string `json:"InstanceTypeFamily" xml:"InstanceTypeFamily"`
            InstanceType     string `json:"InstanceType" xml:"InstanceType"`
            SupportIoOptimized     string `json:"SupportIoOptimized" xml:"SupportIoOptimized"`
            Cores     requests.Integer `json:"Cores" xml:"Cores"`
            Memory     requests.Integer `json:"Memory" xml:"Memory"`
            }  `json:"InstanceType" xml:"InstanceType"`
                Zones struct {
                    Zone []struct {
            ZoneNo     string `json:"ZoneNo" xml:"ZoneNo"`
                NetworkTypes struct {
                NetworkType []    string `json:"NetworkType" xml:"NetworkType"`
                } `json:"NetworkTypes" xml:"NetworkTypes"`
                    }   `json:"zone" xml:"zone"`
                } `json:"Zones" xml:"Zones"`
                    }   `json:"RecommendInstanceType" xml:"RecommendInstanceType"`
                } `json:"Data" xml:"Data"`
}

func CreateDescribeRecommendInstanceTypeRequest() (request *DescribeRecommendInstanceTypeRequest) {
request = &DescribeRecommendInstanceTypeRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeRecommendInstanceType", "", "")
return
}

func CreateDescribeRecommendInstanceTypeResponse() (response *DescribeRecommendInstanceTypeResponse) {
response = &DescribeRecommendInstanceTypeResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

