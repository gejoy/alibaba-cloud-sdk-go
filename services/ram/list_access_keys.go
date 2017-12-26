
package ram

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

func (client *Client) ListAccessKeys(request *ListAccessKeysRequest) (response *ListAccessKeysResponse, err error) {
response = CreateListAccessKeysResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) ListAccessKeysWithChan(request *ListAccessKeysRequest) (<-chan *ListAccessKeysResponse, <-chan error) {
responseChan := make(chan *ListAccessKeysResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.ListAccessKeys(request)
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

func (client *Client) ListAccessKeysWithCallback(request *ListAccessKeysRequest, callback func(response *ListAccessKeysResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *ListAccessKeysResponse
var err error
defer close(result)
response, err = client.ListAccessKeys(request)
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

type ListAccessKeysRequest struct {
*requests.RpcRequest
                UserName  string `position:"Query" name:"UserName"`
}


type ListAccessKeysResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
                AccessKeys struct {
                    AccessKey []struct {
            AccessKeyId     string `json:"AccessKeyId" xml:"AccessKeyId"`
            Status     string `json:"Status" xml:"Status"`
            CreateDate     string `json:"CreateDate" xml:"CreateDate"`
                    }   `json:"AccessKey" xml:"AccessKey"`
                } `json:"AccessKeys" xml:"AccessKeys"`
}

func CreateListAccessKeysRequest() (request *ListAccessKeysRequest) {
request = &ListAccessKeysRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Ram", "2015-05-01", "ListAccessKeys", "", "")
return
}

func CreateListAccessKeysResponse() (response *ListAccessKeysResponse) {
response = &ListAccessKeysResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

