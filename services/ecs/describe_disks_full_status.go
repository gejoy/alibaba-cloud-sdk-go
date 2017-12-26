
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

func (client *Client) DescribeDisksFullStatus(request *DescribeDisksFullStatusRequest) (response *DescribeDisksFullStatusResponse, err error) {
response = CreateDescribeDisksFullStatusResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) DescribeDisksFullStatusWithChan(request *DescribeDisksFullStatusRequest) (<-chan *DescribeDisksFullStatusResponse, <-chan error) {
responseChan := make(chan *DescribeDisksFullStatusResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DescribeDisksFullStatus(request)
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

func (client *Client) DescribeDisksFullStatusWithCallback(request *DescribeDisksFullStatusRequest, callback func(response *DescribeDisksFullStatusResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DescribeDisksFullStatusResponse
var err error
defer close(result)
response, err = client.DescribeDisksFullStatus(request)
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

type DescribeDisksFullStatusRequest struct {
*requests.RpcRequest
                PageSize  string `position:"Query" name:"PageSize"`
                EventTimeEnd  string `position:"Query" name:"EventTime.End"`
                EventId  *[]string `position:"Query" name:"EventId"  type:"Repeated"`
                ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
                HealthStatus  string `position:"Query" name:"HealthStatus"`
                EventTimeStart  string `position:"Query" name:"EventTime.Start"`
                ResourceOwnerId  string `position:"Query" name:"ResourceOwnerId"`
                OwnerAccount  string `position:"Query" name:"OwnerAccount"`
                PageNumber  string `position:"Query" name:"PageNumber"`
                Status  string `position:"Query" name:"Status"`
                OwnerId  string `position:"Query" name:"OwnerId"`
                DiskId  *[]string `position:"Query" name:"DiskId"  type:"Repeated"`
                EventType  string `position:"Query" name:"EventType"`
}


type DescribeDisksFullStatusResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            TotalCount     requests.Integer `json:"TotalCount" xml:"TotalCount"`
            PageNumber     requests.Integer `json:"PageNumber" xml:"PageNumber"`
            PageSize     requests.Integer `json:"PageSize" xml:"PageSize"`
                DiskFullStatusSet struct {
                    DiskFullStatusType []struct {
            DiskId     string `json:"DiskId" xml:"DiskId"`
            Status struct {
            Code     requests.Integer `json:"Code" xml:"Code"`
            Name     string `json:"Name" xml:"Name"`
            }  `json:"Status" xml:"Status"`
            HealthStatus struct {
            Code     requests.Integer `json:"Code" xml:"Code"`
            Name     string `json:"Name" xml:"Name"`
            }  `json:"HealthStatus" xml:"HealthStatus"`
                DiskEventSet struct {
                    DiskEventType []struct {
            EventId     string `json:"EventId" xml:"EventId"`
            EventTime     string `json:"EventTime" xml:"EventTime"`
            EventType struct {
            Code     requests.Integer `json:"Code" xml:"Code"`
            Name     string `json:"Name" xml:"Name"`
            }  `json:"EventType" xml:"EventType"`
                    }   `json:"DiskEventType" xml:"DiskEventType"`
                } `json:"DiskEventSet" xml:"DiskEventSet"`
                    }   `json:"DiskFullStatusType" xml:"DiskFullStatusType"`
                } `json:"DiskFullStatusSet" xml:"DiskFullStatusSet"`
}

func CreateDescribeDisksFullStatusRequest() (request *DescribeDisksFullStatusRequest) {
request = &DescribeDisksFullStatusRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeDisksFullStatus", "", "")
return
}

func CreateDescribeDisksFullStatusResponse() (response *DescribeDisksFullStatusResponse) {
response = &DescribeDisksFullStatusResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

