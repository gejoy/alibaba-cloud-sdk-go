package vod

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

// DeleteTranscodeTemplates invokes the vod.DeleteTranscodeTemplates API synchronously
// api document: https://help.aliyun.com/api/vod/deletetranscodetemplates.html
func (client *Client) DeleteTranscodeTemplates(request *DeleteTranscodeTemplatesRequest) (response *DeleteTranscodeTemplatesResponse, err error) {
	response = CreateDeleteTranscodeTemplatesResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteTranscodeTemplatesWithChan invokes the vod.DeleteTranscodeTemplates API asynchronously
// api document: https://help.aliyun.com/api/vod/deletetranscodetemplates.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteTranscodeTemplatesWithChan(request *DeleteTranscodeTemplatesRequest) (<-chan *DeleteTranscodeTemplatesResponse, <-chan error) {
	responseChan := make(chan *DeleteTranscodeTemplatesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteTranscodeTemplates(request)
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

// DeleteTranscodeTemplatesWithCallback invokes the vod.DeleteTranscodeTemplates API asynchronously
// api document: https://help.aliyun.com/api/vod/deletetranscodetemplates.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteTranscodeTemplatesWithCallback(request *DeleteTranscodeTemplatesRequest, callback func(response *DeleteTranscodeTemplatesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteTranscodeTemplatesResponse
		var err error
		defer close(result)
		response, err = client.DeleteTranscodeTemplates(request)
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

// DeleteTranscodeTemplatesRequest is the request struct for api DeleteTranscodeTemplates
type DeleteTranscodeTemplatesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId          requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount     string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceRealOwnerId      requests.Integer `position:"Query" name:"ResourceRealOwnerId"`
	OwnerId                  requests.Integer `position:"Query" name:"OwnerId"`
	TranscodeTemplateGroupId string           `position:"Query" name:"TranscodeTemplateGroupId"`
	TranscodeTemplateIdList  string           `position:"Query" name:"TranscodeTemplateIdList"`
}

// DeleteTranscodeTemplatesResponse is the response struct for api DeleteTranscodeTemplates
type DeleteTranscodeTemplatesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteTranscodeTemplatesRequest creates a request to invoke DeleteTranscodeTemplates API
func CreateDeleteTranscodeTemplatesRequest() (request *DeleteTranscodeTemplatesRequest) {
	request = &DeleteTranscodeTemplatesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "DeleteTranscodeTemplates", "vod", "openAPI")
	return
}

// CreateDeleteTranscodeTemplatesResponse creates a response to parse from DeleteTranscodeTemplates response
func CreateDeleteTranscodeTemplatesResponse() (response *DeleteTranscodeTemplatesResponse) {
	response = &DeleteTranscodeTemplatesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
