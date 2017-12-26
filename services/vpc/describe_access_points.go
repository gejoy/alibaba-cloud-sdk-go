
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

func (client *Client) DescribeAccessPoints(request *DescribeAccessPointsRequest) (response *DescribeAccessPointsResponse, err error) {
response = CreateDescribeAccessPointsResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) DescribeAccessPointsWithChan(request *DescribeAccessPointsRequest) (<-chan *DescribeAccessPointsResponse, <-chan error) {
responseChan := make(chan *DescribeAccessPointsResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DescribeAccessPoints(request)
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

func (client *Client) DescribeAccessPointsWithCallback(request *DescribeAccessPointsRequest, callback func(response *DescribeAccessPointsResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DescribeAccessPointsResponse
var err error
defer close(result)
response, err = client.DescribeAccessPoints(request)
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

type DescribeAccessPointsRequest struct {
*requests.RpcRequest
                PageSize  string `position:"Query" name:"PageSize"`
                ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
                PageNumber  string `position:"Query" name:"PageNumber"`
                Name  string `position:"Query" name:"Name"`
                ResourceOwnerId  string `position:"Query" name:"ResourceOwnerId"`
                Type  string `position:"Query" name:"Type"`
                OwnerId  string `position:"Query" name:"OwnerId"`
                Filter  *[]DescribeAccessPointsFilter `position:"Query" name:"Filter"  type:"Repeated"`
                HostOperator  string `position:"Query" name:"HostOperator"`
}

type DescribeAccessPointsFilter struct{
        Key string `name:"Key"`
        Value *[]string `name:"Value" type:"Repeated"`
}

type DescribeAccessPointsResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            PageNumber     requests.Integer `json:"PageNumber" xml:"PageNumber"`
            PageSize     requests.Integer `json:"PageSize" xml:"PageSize"`
            TotalCount     requests.Integer `json:"TotalCount" xml:"TotalCount"`
                AccessPointSet struct {
                    AccessPointType []struct {
            AccessPointId     string `json:"AccessPointId" xml:"AccessPointId"`
            Status     string `json:"Status" xml:"Status"`
            Type     string `json:"Type" xml:"Type"`
            AttachedRegionNo     string `json:"AttachedRegionNo" xml:"AttachedRegionNo"`
            Location     string `json:"Location" xml:"Location"`
            HostOperator     string `json:"HostOperator" xml:"HostOperator"`
            Name     string `json:"Name" xml:"Name"`
            Description     string `json:"Description" xml:"Description"`
                    }   `json:"AccessPointType" xml:"AccessPointType"`
                } `json:"AccessPointSet" xml:"AccessPointSet"`
}

func CreateDescribeAccessPointsRequest() (request *DescribeAccessPointsRequest) {
request = &DescribeAccessPointsRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Vpc", "2016-04-28", "DescribeAccessPoints", "", "")
return
}

func CreateDescribeAccessPointsResponse() (response *DescribeAccessPointsResponse) {
response = &DescribeAccessPointsResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

