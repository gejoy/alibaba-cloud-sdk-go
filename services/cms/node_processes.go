
package cms

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

func (client *Client) NodeProcesses(request *NodeProcessesRequest) (response *NodeProcessesResponse, err error) {
response = CreateNodeProcessesResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) NodeProcessesWithChan(request *NodeProcessesRequest) (<-chan *NodeProcessesResponse, <-chan error) {
responseChan := make(chan *NodeProcessesResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.NodeProcesses(request)
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

func (client *Client) NodeProcessesWithCallback(request *NodeProcessesRequest, callback func(response *NodeProcessesResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *NodeProcessesResponse
var err error
defer close(result)
response, err = client.NodeProcesses(request)
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

type NodeProcessesRequest struct {
*requests.RpcRequest
                InstanceId  string `position:"Query" name:"InstanceId"`
}


type NodeProcessesResponse struct {
*responses.BaseResponse
            ErrorCode     requests.Integer `json:"ErrorCode" xml:"ErrorCode"`
            ErrorMessage     string `json:"ErrorMessage" xml:"ErrorMessage"`
            Success     requests.Boolean `json:"Success" xml:"Success"`
            RequestId     string `json:"RequestId" xml:"RequestId"`
                NodeProcesses struct {
                    NodeProcess []struct {
            Id     requests.Integer `json:"Id" xml:"Id"`
            Name     string `json:"Name" xml:"Name"`
            InstanceId     string `json:"InstanceId" xml:"InstanceId"`
            ProcessName     string `json:"ProcessName" xml:"ProcessName"`
            ProcessUser     string `json:"ProcessUser" xml:"ProcessUser"`
            Command     string `json:"Command" xml:"Command"`
                    }   `json:"NodeProcess" xml:"NodeProcess"`
                } `json:"NodeProcesses" xml:"NodeProcesses"`
}

func CreateNodeProcessesRequest() (request *NodeProcessesRequest) {
request = &NodeProcessesRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Cms", "2017-03-01", "NodeProcesses", "", "")
return
}

func CreateNodeProcessesResponse() (response *NodeProcessesResponse) {
response = &NodeProcessesResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

