package baas

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

// CreateChannelMember invokes the baas.CreateChannelMember API synchronously
// api document: https://help.aliyun.com/api/baas/createchannelmember.html
func (client *Client) CreateChannelMember(request *CreateChannelMemberRequest) (response *CreateChannelMemberResponse, err error) {
	response = CreateCreateChannelMemberResponse()
	err = client.DoAction(request, response)
	return
}

// CreateChannelMemberWithChan invokes the baas.CreateChannelMember API asynchronously
// api document: https://help.aliyun.com/api/baas/createchannelmember.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateChannelMemberWithChan(request *CreateChannelMemberRequest) (<-chan *CreateChannelMemberResponse, <-chan error) {
	responseChan := make(chan *CreateChannelMemberResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateChannelMember(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// CreateChannelMemberWithCallback invokes the baas.CreateChannelMember API asynchronously
// api document: https://help.aliyun.com/api/baas/createchannelmember.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateChannelMemberWithCallback(request *CreateChannelMemberRequest, callback func(response *CreateChannelMemberResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateChannelMemberResponse
		var err error
		defer close(result)
		response, err = client.CreateChannelMember(request)
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

// CreateChannelMemberRequest is the request struct for api CreateChannelMember
type CreateChannelMemberRequest struct {
	*requests.RpcRequest
	Organization *[]CreateChannelMemberOrganization `position:"Query" name:"Organization"  type:"Repeated"`
	ChannelId    string                             `position:"Query" name:"ChannelId"`
}

// CreateChannelMemberOrganization is a repeated param struct in CreateChannelMemberRequest
type CreateChannelMemberOrganization struct {
	Id string `name:"Id"`
}

// CreateChannelMemberResponse is the response struct for api CreateChannelMember
type CreateChannelMemberResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	Result    bool   `json:"Result" xml:"Result"`
}

// CreateCreateChannelMemberRequest creates a request to invoke CreateChannelMember API
func CreateCreateChannelMemberRequest() (request *CreateChannelMemberRequest) {
	request = &CreateChannelMemberRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Baas", "2018-07-31", "CreateChannelMember", "", "")
	return
}

// CreateCreateChannelMemberResponse creates a response to parse from CreateChannelMember response
func CreateCreateChannelMemberResponse() (response *CreateChannelMemberResponse) {
	response = &CreateChannelMemberResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}