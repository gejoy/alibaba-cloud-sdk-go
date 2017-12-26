
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

func (client *Client) QueryMetricLast(request *QueryMetricLastRequest) (response *QueryMetricLastResponse, err error) {
response = CreateQueryMetricLastResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) QueryMetricLastWithChan(request *QueryMetricLastRequest) (<-chan *QueryMetricLastResponse, <-chan error) {
responseChan := make(chan *QueryMetricLastResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.QueryMetricLast(request)
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

func (client *Client) QueryMetricLastWithCallback(request *QueryMetricLastRequest, callback func(response *QueryMetricLastResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *QueryMetricLastResponse
var err error
defer close(result)
response, err = client.QueryMetricLast(request)
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

type QueryMetricLastRequest struct {
*requests.RpcRequest
                EndTime  string `position:"Query" name:"EndTime"`
                StartTime  string `position:"Query" name:"StartTime"`
                Cursor  string `position:"Query" name:"Cursor"`
                Express  string `position:"Query" name:"Express"`
                Period  string `position:"Query" name:"Period"`
                Project  string `position:"Query" name:"Project"`
                Page  string `position:"Query" name:"Page"`
                Metric  string `position:"Query" name:"Metric"`
                Length  string `position:"Query" name:"Length"`
                Dimensions  string `position:"Query" name:"Dimensions"`
                ResourceOwnerId  string `position:"Query" name:"ResourceOwnerId"`
                CallbyCmsOwner  string `position:"Query" name:"callby_cms_owner"`
}


type QueryMetricLastResponse struct {
*responses.BaseResponse
            Code     string `json:"Code" xml:"Code"`
            Message     string `json:"Message" xml:"Message"`
            Success     requests.Boolean `json:"Success" xml:"Success"`
            RequestId     string `json:"RequestId" xml:"RequestId"`
            Cursor     string `json:"Cursor" xml:"Cursor"`
            Datapoints     string `json:"Datapoints" xml:"Datapoints"`
            Period     string `json:"Period" xml:"Period"`
}

func CreateQueryMetricLastRequest() (request *QueryMetricLastRequest) {
request = &QueryMetricLastRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Cms", "2017-03-01", "QueryMetricLast", "", "")
return
}

func CreateQueryMetricLastResponse() (response *QueryMetricLastResponse) {
response = &QueryMetricLastResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

