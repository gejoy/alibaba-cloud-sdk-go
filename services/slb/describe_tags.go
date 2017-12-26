
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

func (client *Client) DescribeTags(request *DescribeTagsRequest) (response *DescribeTagsResponse, err error) {
response = CreateDescribeTagsResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) DescribeTagsWithChan(request *DescribeTagsRequest) (<-chan *DescribeTagsResponse, <-chan error) {
responseChan := make(chan *DescribeTagsResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DescribeTags(request)
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

func (client *Client) DescribeTagsWithCallback(request *DescribeTagsRequest, callback func(response *DescribeTagsResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DescribeTagsResponse
var err error
defer close(result)
response, err = client.DescribeTags(request)
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

type DescribeTagsRequest struct {
*requests.RpcRequest
                Tags  string `position:"Query" name:"Tags"`
                PageSize  string `position:"Query" name:"PageSize"`
                PageNumber  string `position:"Query" name:"PageNumber"`
                OwnerId  string `position:"Query" name:"OwnerId"`
                ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
                AccessKeyId  string `position:"Query" name:"access_key_id"`
                ResourceOwnerId  string `position:"Query" name:"ResourceOwnerId"`
                OwnerAccount  string `position:"Query" name:"OwnerAccount"`
                LoadBalancerId  string `position:"Query" name:"LoadBalancerId"`
                DistinctKey  string `position:"Query" name:"DistinctKey"`
}


type DescribeTagsResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            PageSize     requests.Integer `json:"PageSize" xml:"PageSize"`
            PageNumber     requests.Integer `json:"PageNumber" xml:"PageNumber"`
            TotalCount     requests.Integer `json:"TotalCount" xml:"TotalCount"`
                TagSets struct {
                    TagSet []struct {
            TagKey     string `json:"TagKey" xml:"TagKey"`
            TagValue     string `json:"TagValue" xml:"TagValue"`
            InstanceCount     requests.Integer `json:"InstanceCount" xml:"InstanceCount"`
                    }   `json:"TagSet" xml:"TagSet"`
                } `json:"TagSets" xml:"TagSets"`
}

func CreateDescribeTagsRequest() (request *DescribeTagsRequest) {
request = &DescribeTagsRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Slb", "2014-05-15", "DescribeTags", "", "")
return
}

func CreateDescribeTagsResponse() (response *DescribeTagsResponse) {
response = &DescribeTagsResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

