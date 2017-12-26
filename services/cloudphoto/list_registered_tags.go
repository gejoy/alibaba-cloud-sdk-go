
package cloudphoto

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

func (client *Client) ListRegisteredTags(request *ListRegisteredTagsRequest) (response *ListRegisteredTagsResponse, err error) {
response = CreateListRegisteredTagsResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) ListRegisteredTagsWithChan(request *ListRegisteredTagsRequest) (<-chan *ListRegisteredTagsResponse, <-chan error) {
responseChan := make(chan *ListRegisteredTagsResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.ListRegisteredTags(request)
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

func (client *Client) ListRegisteredTagsWithCallback(request *ListRegisteredTagsRequest, callback func(response *ListRegisteredTagsResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *ListRegisteredTagsResponse
var err error
defer close(result)
response, err = client.ListRegisteredTags(request)
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

type ListRegisteredTagsRequest struct {
*requests.RpcRequest
                StoreName  string `position:"Query" name:"StoreName"`
                Lang  *[]string `position:"Query" name:"Lang"  type:"Repeated"`
}


type ListRegisteredTagsResponse struct {
*responses.BaseResponse
            Code     string `json:"Code" xml:"Code"`
            Message     string `json:"Message" xml:"Message"`
            RequestId     string `json:"RequestId" xml:"RequestId"`
            Action     string `json:"Action" xml:"Action"`
                RegisteredTags  []struct {
            TagKey     string `json:"TagKey" xml:"TagKey"`
                TagValues  []struct {
            Lang     string `json:"Lang" xml:"Lang"`
            Text     string `json:"Text" xml:"Text"`
                }  `json:"TagValues" xml:"TagValues"`
                }  `json:"RegisteredTags" xml:"RegisteredTags"`
}

func CreateListRegisteredTagsRequest() (request *ListRegisteredTagsRequest) {
request = &ListRegisteredTagsRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("CloudPhoto", "2017-07-11", "ListRegisteredTags", "", "")
return
}

func CreateListRegisteredTagsResponse() (response *ListRegisteredTagsResponse) {
response = &ListRegisteredTagsResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

