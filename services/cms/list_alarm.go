
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

func (client *Client) ListAlarm(request *ListAlarmRequest) (response *ListAlarmResponse, err error) {
response = CreateListAlarmResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) ListAlarmWithChan(request *ListAlarmRequest) (<-chan *ListAlarmResponse, <-chan error) {
responseChan := make(chan *ListAlarmResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.ListAlarm(request)
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

func (client *Client) ListAlarmWithCallback(request *ListAlarmRequest, callback func(response *ListAlarmResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *ListAlarmResponse
var err error
defer close(result)
response, err = client.ListAlarm(request)
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

type ListAlarmRequest struct {
*requests.RpcRequest
                Id  string `position:"Query" name:"Id"`
                PageSize  string `position:"Query" name:"PageSize"`
                Dimension  string `position:"Query" name:"Dimension"`
                PageNumber  string `position:"Query" name:"PageNumber"`
                Name  string `position:"Query" name:"Name"`
                State  string `position:"Query" name:"State"`
                IsEnable  string `position:"Query" name:"IsEnable"`
                Namespace  string `position:"Query" name:"Namespace"`
                CallbyCmsOwner  string `position:"Query" name:"callby_cms_owner"`
}


type ListAlarmResponse struct {
*responses.BaseResponse
            Success     requests.Boolean `json:"Success" xml:"Success"`
            Code     string `json:"Code" xml:"Code"`
            Message     string `json:"Message" xml:"Message"`
            NextToken     requests.Integer `json:"NextToken" xml:"NextToken"`
            Total     requests.Integer `json:"Total" xml:"Total"`
            RequestId     string `json:"RequestId" xml:"RequestId"`
                AlarmList struct {
                    Alarm []struct {
            Id     string `json:"Id" xml:"Id"`
            Name     string `json:"Name" xml:"Name"`
            Namespace     string `json:"Namespace" xml:"Namespace"`
            MetricName     string `json:"MetricName" xml:"MetricName"`
            Dimensions     string `json:"Dimensions" xml:"Dimensions"`
            Period     requests.Integer `json:"Period" xml:"Period"`
            Statistics     string `json:"Statistics" xml:"Statistics"`
            ComparisonOperator     string `json:"ComparisonOperator" xml:"ComparisonOperator"`
            Threshold     string `json:"Threshold" xml:"Threshold"`
            EvaluationCount     requests.Integer `json:"EvaluationCount" xml:"EvaluationCount"`
            StartTime     requests.Integer `json:"StartTime" xml:"StartTime"`
            EndTime     requests.Integer `json:"EndTime" xml:"EndTime"`
            SilenceTime     requests.Integer `json:"SilenceTime" xml:"SilenceTime"`
            NotifyType     requests.Integer `json:"NotifyType" xml:"NotifyType"`
            Enable     requests.Boolean `json:"Enable" xml:"Enable"`
            State     string `json:"State" xml:"State"`
            ContactGroups     string `json:"ContactGroups" xml:"ContactGroups"`
            Webhook     string `json:"Webhook" xml:"Webhook"`
                    }   `json:"Alarm" xml:"Alarm"`
                } `json:"AlarmList" xml:"AlarmList"`
}

func CreateListAlarmRequest() (request *ListAlarmRequest) {
request = &ListAlarmRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Cms", "2017-03-01", "ListAlarm", "", "")
return
}

func CreateListAlarmResponse() (response *ListAlarmResponse) {
response = &ListAlarmResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

