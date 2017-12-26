
package cs

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

func (client *Client) CallBackAgilityCluster(request *CallBackAgilityClusterRequest) (response *CallBackAgilityClusterResponse, err error) {
response = CreateCallBackAgilityClusterResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) CallBackAgilityClusterWithChan(request *CallBackAgilityClusterRequest) (<-chan *CallBackAgilityClusterResponse, <-chan error) {
responseChan := make(chan *CallBackAgilityClusterResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.CallBackAgilityCluster(request)
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

func (client *Client) CallBackAgilityClusterWithCallback(request *CallBackAgilityClusterRequest, callback func(response *CallBackAgilityClusterResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *CallBackAgilityClusterResponse
var err error
defer close(result)
response, err = client.CallBackAgilityCluster(request)
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

type CallBackAgilityClusterRequest struct {
*requests.RoaRequest
                ReqOnce  string `position:"Path" name:"ReqOnce"`
                Token  string `position:"Path" name:"Token"`
}


type CallBackAgilityClusterResponse struct {
*responses.BaseResponse
}

func CreateCallBackAgilityClusterRequest() (request *CallBackAgilityClusterRequest) {
request = &CallBackAgilityClusterRequest{
RoaRequest: &requests.RoaRequest{},
}
request.InitWithApiInfo("CS", "2015-12-15", "CallBackAgilityCluster", "/agility/token/[Token]/req_once/[ReqOnce]/callback", "", "")
return
}

func CreateCallBackAgilityClusterResponse() (response *CallBackAgilityClusterResponse) {
response = &CallBackAgilityClusterResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

