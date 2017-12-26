
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

func (client *Client) DescribeVRouters(request *DescribeVRoutersRequest) (response *DescribeVRoutersResponse, err error) {
response = CreateDescribeVRoutersResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) DescribeVRoutersWithChan(request *DescribeVRoutersRequest) (<-chan *DescribeVRoutersResponse, <-chan error) {
responseChan := make(chan *DescribeVRoutersResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DescribeVRouters(request)
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

func (client *Client) DescribeVRoutersWithCallback(request *DescribeVRoutersRequest, callback func(response *DescribeVRoutersResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DescribeVRoutersResponse
var err error
defer close(result)
response, err = client.DescribeVRouters(request)
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

type DescribeVRoutersRequest struct {
*requests.RpcRequest
                PageSize  string `position:"Query" name:"PageSize"`
                ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
                PageNumber  string `position:"Query" name:"PageNumber"`
                VRouterId  string `position:"Query" name:"VRouterId"`
                ResourceOwnerId  string `position:"Query" name:"ResourceOwnerId"`
                OwnerAccount  string `position:"Query" name:"OwnerAccount"`
                OwnerId  string `position:"Query" name:"OwnerId"`
}


type DescribeVRoutersResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            TotalCount     requests.Integer `json:"TotalCount" xml:"TotalCount"`
            PageNumber     requests.Integer `json:"PageNumber" xml:"PageNumber"`
            PageSize     requests.Integer `json:"PageSize" xml:"PageSize"`
                VRouters struct {
                    VRouter []struct {
            RegionId     string `json:"RegionId" xml:"RegionId"`
            VpcId     string `json:"VpcId" xml:"VpcId"`
            VRouterName     string `json:"VRouterName" xml:"VRouterName"`
            Description     string `json:"Description" xml:"Description"`
            VRouterId     string `json:"VRouterId" xml:"VRouterId"`
            CreationTime     string `json:"CreationTime" xml:"CreationTime"`
                RouteTableIds struct {
                RouteTableId []    string `json:"RouteTableId" xml:"RouteTableId"`
                } `json:"RouteTableIds" xml:"RouteTableIds"`
                    }   `json:"VRouter" xml:"VRouter"`
                } `json:"VRouters" xml:"VRouters"`
}

func CreateDescribeVRoutersRequest() (request *DescribeVRoutersRequest) {
request = &DescribeVRoutersRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Vpc", "2016-04-28", "DescribeVRouters", "", "")
return
}

func CreateDescribeVRoutersResponse() (response *DescribeVRoutersResponse) {
response = &DescribeVRoutersResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

