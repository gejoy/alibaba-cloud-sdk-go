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

// DescribeUserVvTopByDay invokes the vod.DescribeUserVvTopByDay API synchronously
// api document: https://help.aliyun.com/api/vod/describeuservvtopbyday.html
func (client *Client) DescribeUserVvTopByDay(request *DescribeUserVvTopByDayRequest) (response *DescribeUserVvTopByDayResponse, err error) {
	response = CreateDescribeUserVvTopByDayResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeUserVvTopByDayWithChan invokes the vod.DescribeUserVvTopByDay API asynchronously
// api document: https://help.aliyun.com/api/vod/describeuservvtopbyday.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeUserVvTopByDayWithChan(request *DescribeUserVvTopByDayRequest) (<-chan *DescribeUserVvTopByDayResponse, <-chan error) {
	responseChan := make(chan *DescribeUserVvTopByDayResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeUserVvTopByDay(request)
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

// DescribeUserVvTopByDayWithCallback invokes the vod.DescribeUserVvTopByDay API asynchronously
// api document: https://help.aliyun.com/api/vod/describeuservvtopbyday.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeUserVvTopByDayWithCallback(request *DescribeUserVvTopByDayRequest, callback func(response *DescribeUserVvTopByDayResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeUserVvTopByDayResponse
		var err error
		defer close(result)
		response, err = client.DescribeUserVvTopByDay(request)
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

// DescribeUserVvTopByDayRequest is the request struct for api DescribeUserVvTopByDay
type DescribeUserVvTopByDayRequest struct {
	*requests.RpcRequest
	VideoType     string           `position:"Query" name:"VideoType"`
	Bizdate       string           `position:"Query" name:"Bizdate"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeUserVvTopByDayResponse is the response struct for api DescribeUserVvTopByDay
type DescribeUserVvTopByDayResponse struct {
	*responses.BaseResponse
	RequestId               string                                          `json:"RequestId" xml:"RequestId"`
	UserPlayStatisticsInfos UserPlayStatisticsInfosInDescribeUserVvTopByDay `json:"UserPlayStatisticsInfos" xml:"UserPlayStatisticsInfos"`
}

// CreateDescribeUserVvTopByDayRequest creates a request to invoke DescribeUserVvTopByDay API
func CreateDescribeUserVvTopByDayRequest() (request *DescribeUserVvTopByDayRequest) {
	request = &DescribeUserVvTopByDayRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "DescribeUserVvTopByDay", "vod", "openAPI")
	return
}

// CreateDescribeUserVvTopByDayResponse creates a response to parse from DescribeUserVvTopByDay response
func CreateDescribeUserVvTopByDayResponse() (response *DescribeUserVvTopByDayResponse) {
	response = &DescribeUserVvTopByDayResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
