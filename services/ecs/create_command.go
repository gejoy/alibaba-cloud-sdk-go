
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

func (client *Client) CreateCommand(request *CreateCommandRequest) (response *CreateCommandResponse, err error) {
response = CreateCreateCommandResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) CreateCommandWithChan(request *CreateCommandRequest) (<-chan *CreateCommandResponse, <-chan error) {
responseChan := make(chan *CreateCommandResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.CreateCommand(request)
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

func (client *Client) CreateCommandWithCallback(request *CreateCommandRequest, callback func(response *CreateCommandResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *CreateCommandResponse
var err error
defer close(result)
response, err = client.CreateCommand(request)
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

type CreateCommandRequest struct {
*requests.RpcRequest
                WorkingDir  string `position:"Query" name:"WorkingDir"`
                Type  string `position:"Query" name:"Type"`
                ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
                Description  string `position:"Query" name:"Description"`
                Name  string `position:"Query" name:"Name"`
                ResourceOwnerId  string `position:"Query" name:"ResourceOwnerId"`
                OwnerAccount  string `position:"Query" name:"OwnerAccount"`
                CommandContent  string `position:"Query" name:"CommandContent"`
                OwnerId  string `position:"Query" name:"OwnerId"`
                Timeout  string `position:"Query" name:"Timeout"`
}


type CreateCommandResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            CommandId     string `json:"CommandId" xml:"CommandId"`
}

func CreateCreateCommandRequest() (request *CreateCommandRequest) {
request = &CreateCommandRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Ecs", "2014-05-26", "CreateCommand", "", "")
return
}

func CreateCreateCommandResponse() (response *CreateCommandResponse) {
response = &CreateCommandResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

