
package ccc

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

func (client *Client) ModifyUser(request *ModifyUserRequest) (response *ModifyUserResponse, err error) {
response = CreateModifyUserResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) ModifyUserWithChan(request *ModifyUserRequest) (<-chan *ModifyUserResponse, <-chan error) {
responseChan := make(chan *ModifyUserResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.ModifyUser(request)
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

func (client *Client) ModifyUserWithCallback(request *ModifyUserRequest, callback func(response *ModifyUserResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *ModifyUserResponse
var err error
defer close(result)
response, err = client.ModifyUser(request)
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

type ModifyUserRequest struct {
*requests.RpcRequest
                Phone  string `position:"Query" name:"Phone"`
                SkillLevel  *[]string `position:"Query" name:"SkillLevel"  type:"Repeated"`
                RoleId  *[]string `position:"Query" name:"RoleId"  type:"Repeated"`
                Email  string `position:"Query" name:"Email"`
                SkillGroupId  *[]string `position:"Query" name:"SkillGroupId"  type:"Repeated"`
                UserId  string `position:"Query" name:"UserId"`
                InstanceId  string `position:"Query" name:"InstanceId"`
                DisplayName  string `position:"Query" name:"DisplayName"`
}


type ModifyUserResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            Success     requests.Boolean `json:"Success" xml:"Success"`
            Code     string `json:"Code" xml:"Code"`
            Message     string `json:"Message" xml:"Message"`
            HttpStatusCode     requests.Integer `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

func CreateModifyUserRequest() (request *ModifyUserRequest) {
request = &ModifyUserRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("CCC", "2017-07-05", "ModifyUser", "", "")
return
}

func CreateModifyUserResponse() (response *ModifyUserResponse) {
response = &ModifyUserResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

