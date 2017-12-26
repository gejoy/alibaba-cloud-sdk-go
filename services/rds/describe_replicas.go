
package rds

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

func (client *Client) DescribeReplicas(request *DescribeReplicasRequest) (response *DescribeReplicasResponse, err error) {
response = CreateDescribeReplicasResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) DescribeReplicasWithChan(request *DescribeReplicasRequest) (<-chan *DescribeReplicasResponse, <-chan error) {
responseChan := make(chan *DescribeReplicasResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DescribeReplicas(request)
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

func (client *Client) DescribeReplicasWithCallback(request *DescribeReplicasRequest, callback func(response *DescribeReplicasResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DescribeReplicasResponse
var err error
defer close(result)
response, err = client.DescribeReplicas(request)
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

type DescribeReplicasRequest struct {
*requests.RpcRequest
                PageSize  string `position:"Query" name:"PageSize"`
                ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
                PageNumber  string `position:"Query" name:"PageNumber"`
                ReplicaId  string `position:"Query" name:"ReplicaId"`
                ResourceOwnerId  string `position:"Query" name:"ResourceOwnerId"`
                OwnerAccount  string `position:"Query" name:"OwnerAccount"`
                OwnerId  string `position:"Query" name:"OwnerId"`
                SecurityToken  string `position:"Query" name:"SecurityToken"`
}


type DescribeReplicasResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            PageNumber     string `json:"PageNumber" xml:"PageNumber"`
            TotalRecordCount     requests.Integer `json:"TotalRecordCount" xml:"TotalRecordCount"`
            PageRecordCount     requests.Integer `json:"PageRecordCount" xml:"PageRecordCount"`
                Replicas  []struct {
            ReplicaId     string `json:"ReplicaId" xml:"ReplicaId"`
            ReplicaDescription     string `json:"ReplicaDescription" xml:"ReplicaDescription"`
            ReplicaStatus     string `json:"ReplicaStatus" xml:"ReplicaStatus"`
            ReplicaMode     string `json:"ReplicaMode" xml:"ReplicaMode"`
            DomainMode     string `json:"DomainMode" xml:"DomainMode"`
                DBInstances  []struct {
            DBInstanceId     string `json:"DBInstanceId" xml:"DBInstanceId"`
            Role     string `json:"Role" xml:"Role"`
            ReadWriteType     string `json:"ReadWriteType" xml:"ReadWriteType"`
                }  `json:"DBInstances" xml:"DBInstances"`
                }  `json:"Replicas" xml:"Replicas"`
}

func CreateDescribeReplicasRequest() (request *DescribeReplicasRequest) {
request = &DescribeReplicasRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2014-08-15", "DescribeReplicas", "", "")
return
}

func CreateDescribeReplicasResponse() (response *DescribeReplicasResponse) {
response = &DescribeReplicasResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

