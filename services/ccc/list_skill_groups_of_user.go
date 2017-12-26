
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

func (client *Client) ListSkillGroupsOfUser(request *ListSkillGroupsOfUserRequest) (response *ListSkillGroupsOfUserResponse, err error) {
response = CreateListSkillGroupsOfUserResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) ListSkillGroupsOfUserWithChan(request *ListSkillGroupsOfUserRequest) (<-chan *ListSkillGroupsOfUserResponse, <-chan error) {
responseChan := make(chan *ListSkillGroupsOfUserResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.ListSkillGroupsOfUser(request)
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

func (client *Client) ListSkillGroupsOfUserWithCallback(request *ListSkillGroupsOfUserRequest, callback func(response *ListSkillGroupsOfUserResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *ListSkillGroupsOfUserResponse
var err error
defer close(result)
response, err = client.ListSkillGroupsOfUser(request)
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

type ListSkillGroupsOfUserRequest struct {
*requests.RpcRequest
                UserId  string `position:"Query" name:"UserId"`
                InstanceId  string `position:"Query" name:"InstanceId"`
}


type ListSkillGroupsOfUserResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            Success     requests.Boolean `json:"Success" xml:"Success"`
            Code     string `json:"Code" xml:"Code"`
            Message     string `json:"Message" xml:"Message"`
            HttpStatusCode     requests.Integer `json:"HttpStatusCode" xml:"HttpStatusCode"`
                SkillLevels struct {
                    SkillLevel []struct {
            SkillLevelId     string `json:"SkillLevelId" xml:"SkillLevelId"`
            Level     requests.Integer `json:"Level" xml:"Level"`
            Skill struct {
            SkillGroupId     string `json:"SkillGroupId" xml:"SkillGroupId"`
            InstanceId     string `json:"InstanceId" xml:"InstanceId"`
            SkillGroupName     string `json:"SkillGroupName" xml:"SkillGroupName"`
            SkillGroupDescription     string `json:"SkillGroupDescription" xml:"SkillGroupDescription"`
                OutboundPhoneNumbers struct {
                    PhoneNumber []struct {
            PhoneNumberId     string `json:"PhoneNumberId" xml:"PhoneNumberId"`
            InstanceId     string `json:"InstanceId" xml:"InstanceId"`
            Number     string `json:"Number" xml:"Number"`
            PhoneNumberDescription     string `json:"PhoneNumberDescription" xml:"PhoneNumberDescription"`
            TestOnly     requests.Boolean `json:"TestOnly" xml:"TestOnly"`
            RemainingTime     requests.Integer `json:"RemainingTime" xml:"RemainingTime"`
            AllowOutbound     requests.Boolean `json:"AllowOutbound" xml:"AllowOutbound"`
            Usage     string `json:"Usage" xml:"Usage"`
            Trunks     requests.Integer `json:"Trunks" xml:"Trunks"`
                    }   `json:"PhoneNumber" xml:"PhoneNumber"`
                } `json:"OutboundPhoneNumbers" xml:"OutboundPhoneNumbers"`
            }  `json:"Skill" xml:"Skill"`
                    }   `json:"SkillLevel" xml:"SkillLevel"`
                } `json:"SkillLevels" xml:"SkillLevels"`
}

func CreateListSkillGroupsOfUserRequest() (request *ListSkillGroupsOfUserRequest) {
request = &ListSkillGroupsOfUserRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("CCC", "2017-07-05", "ListSkillGroupsOfUser", "", "")
return
}

func CreateListSkillGroupsOfUserResponse() (response *ListSkillGroupsOfUserResponse) {
response = &ListSkillGroupsOfUserResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

