
package vpc

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

func (client *Client) DescribeSnatTableEntries(request *DescribeSnatTableEntriesRequest) (response *DescribeSnatTableEntriesResponse, err error) {
response = CreateDescribeSnatTableEntriesResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) DescribeSnatTableEntriesWithChan(request *DescribeSnatTableEntriesRequest) (<-chan *DescribeSnatTableEntriesResponse, <-chan error) {
responseChan := make(chan *DescribeSnatTableEntriesResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DescribeSnatTableEntries(request)
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

func (client *Client) DescribeSnatTableEntriesWithCallback(request *DescribeSnatTableEntriesRequest, callback func(response *DescribeSnatTableEntriesResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DescribeSnatTableEntriesResponse
var err error
defer close(result)
response, err = client.DescribeSnatTableEntries(request)
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

type DescribeSnatTableEntriesRequest struct {
*requests.RpcRequest
                SnatEntryId  string `position:"Query" name:"SnatEntryId"`
                PageSize  string `position:"Query" name:"PageSize"`
                ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
                PageNumber  string `position:"Query" name:"PageNumber"`
                ResourceOwnerId  string `position:"Query" name:"ResourceOwnerId"`
                OwnerAccount  string `position:"Query" name:"OwnerAccount"`
                OwnerId  string `position:"Query" name:"OwnerId"`
                SnatTableId  string `position:"Query" name:"SnatTableId"`
}


type DescribeSnatTableEntriesResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            TotalCount     requests.Integer `json:"TotalCount" xml:"TotalCount"`
            PageNumber     requests.Integer `json:"PageNumber" xml:"PageNumber"`
            PageSize     requests.Integer `json:"PageSize" xml:"PageSize"`
                SnatTableEntries struct {
                    SnatTableEntry []struct {
            SnatTableId     string `json:"SnatTableId" xml:"SnatTableId"`
            SnatEntryId     string `json:"SnatEntryId" xml:"SnatEntryId"`
            SourceVSwitchId     string `json:"SourceVSwitchId" xml:"SourceVSwitchId"`
            SourceCIDR     string `json:"SourceCIDR" xml:"SourceCIDR"`
            SnatIp     string `json:"SnatIp" xml:"SnatIp"`
            Status     string `json:"Status" xml:"Status"`
                    }   `json:"SnatTableEntry" xml:"SnatTableEntry"`
                } `json:"SnatTableEntries" xml:"SnatTableEntries"`
}

func CreateDescribeSnatTableEntriesRequest() (request *DescribeSnatTableEntriesRequest) {
request = &DescribeSnatTableEntriesRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Vpc", "2016-04-28", "DescribeSnatTableEntries", "", "")
return
}

func CreateDescribeSnatTableEntriesResponse() (response *DescribeSnatTableEntriesResponse) {
response = &DescribeSnatTableEntriesResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

