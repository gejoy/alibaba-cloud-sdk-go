
package alidns

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

func (client *Client) DescribeDomainInfo(request *DescribeDomainInfoRequest) (response *DescribeDomainInfoResponse, err error) {
response = CreateDescribeDomainInfoResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) DescribeDomainInfoWithChan(request *DescribeDomainInfoRequest) (<-chan *DescribeDomainInfoResponse, <-chan error) {
responseChan := make(chan *DescribeDomainInfoResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DescribeDomainInfo(request)
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

func (client *Client) DescribeDomainInfoWithCallback(request *DescribeDomainInfoRequest, callback func(response *DescribeDomainInfoResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DescribeDomainInfoResponse
var err error
defer close(result)
response, err = client.DescribeDomainInfo(request)
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

type DescribeDomainInfoRequest struct {
*requests.RpcRequest
                NeedDetailAttributes  string `position:"Query" name:"NeedDetailAttributes"`
                DomainName  string `position:"Query" name:"DomainName"`
                UserClientIp  string `position:"Query" name:"UserClientIp"`
                Lang  string `position:"Query" name:"Lang"`
}


type DescribeDomainInfoResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            DomainId     string `json:"DomainId" xml:"DomainId"`
            DomainName     string `json:"DomainName" xml:"DomainName"`
            PunyCode     string `json:"PunyCode" xml:"PunyCode"`
            AliDomain     requests.Boolean `json:"AliDomain" xml:"AliDomain"`
            Remark     string `json:"Remark" xml:"Remark"`
            GroupId     string `json:"GroupId" xml:"GroupId"`
            GroupName     string `json:"GroupName" xml:"GroupName"`
            InstanceId     string `json:"InstanceId" xml:"InstanceId"`
            VersionCode     string `json:"VersionCode" xml:"VersionCode"`
            VersionName     string `json:"VersionName" xml:"VersionName"`
            MinTtl     requests.Integer `json:"MinTtl" xml:"MinTtl"`
            RecordLineTreeJson     string `json:"RecordLineTreeJson" xml:"RecordLineTreeJson"`
            LineType     string `json:"LineType" xml:"LineType"`
            RegionLines     requests.Boolean `json:"RegionLines" xml:"RegionLines"`
                DnsServers struct {
                DnsServer []    string `json:"DnsServer" xml:"DnsServer"`
                } `json:"DnsServers" xml:"DnsServers"`
                AvailableTtls struct {
                AvailableTtl []    string `json:"AvailableTtl" xml:"AvailableTtl"`
                } `json:"AvailableTtls" xml:"AvailableTtls"`
                RecordLines struct {
                    RecordLine []struct {
            LineCode     string `json:"LineCode" xml:"LineCode"`
            FatherCode     string `json:"FatherCode" xml:"FatherCode"`
            LineName     string `json:"LineName" xml:"LineName"`
            LineDisplayName     string `json:"LineDisplayName" xml:"LineDisplayName"`
                    }   `json:"RecordLine" xml:"RecordLine"`
                } `json:"RecordLines" xml:"RecordLines"`
}

func CreateDescribeDomainInfoRequest() (request *DescribeDomainInfoRequest) {
request = &DescribeDomainInfoRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Alidns", "2015-01-09", "DescribeDomainInfo", "", "")
return
}

func CreateDescribeDomainInfoResponse() (response *DescribeDomainInfoResponse) {
response = &DescribeDomainInfoResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

