
package cdn

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

func (client *Client) StartMixStreamsService(request *StartMixStreamsServiceRequest) (response *StartMixStreamsServiceResponse, err error) {
response = CreateStartMixStreamsServiceResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) StartMixStreamsServiceWithChan(request *StartMixStreamsServiceRequest) (<-chan *StartMixStreamsServiceResponse, <-chan error) {
responseChan := make(chan *StartMixStreamsServiceResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.StartMixStreamsService(request)
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

func (client *Client) StartMixStreamsServiceWithCallback(request *StartMixStreamsServiceRequest, callback func(response *StartMixStreamsServiceResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *StartMixStreamsServiceResponse
var err error
defer close(result)
response, err = client.StartMixStreamsService(request)
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

type StartMixStreamsServiceRequest struct {
*requests.RpcRequest
                MixStreamName  string `position:"Query" name:"MixStreamName"`
                MixType  string `position:"Query" name:"MixType"`
                MainDomainName  string `position:"Query" name:"MainDomainName"`
                MixTemplate  string `position:"Query" name:"MixTemplate"`
                MixAppName  string `position:"Query" name:"MixAppName"`
                MainStreamName  string `position:"Query" name:"MainStreamName"`
                OwnerId  string `position:"Query" name:"OwnerId"`
                SecurityToken  string `position:"Query" name:"SecurityToken"`
                MixDomainName  string `position:"Query" name:"MixDomainName"`
                MainAppName  string `position:"Query" name:"MainAppName"`
}


type StartMixStreamsServiceResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
                MixStreamsInfoList struct {
                    MixStreamsInfo []struct {
            DomainName     string `json:"DomainName" xml:"DomainName"`
            AppName     string `json:"AppName" xml:"AppName"`
            StreamName     string `json:"StreamName" xml:"StreamName"`
                    }   `json:"MixStreamsInfo" xml:"MixStreamsInfo"`
                } `json:"MixStreamsInfoList" xml:"MixStreamsInfoList"`
}

func CreateStartMixStreamsServiceRequest() (request *StartMixStreamsServiceRequest) {
request = &StartMixStreamsServiceRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Cdn", "2014-11-11", "StartMixStreamsService", "", "")
return
}

func CreateStartMixStreamsServiceResponse() (response *StartMixStreamsServiceResponse) {
response = &StartMixStreamsServiceResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

