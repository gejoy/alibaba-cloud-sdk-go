
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

func (client *Client) DescribeBandwidthPackages(request *DescribeBandwidthPackagesRequest) (response *DescribeBandwidthPackagesResponse, err error) {
response = CreateDescribeBandwidthPackagesResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) DescribeBandwidthPackagesWithChan(request *DescribeBandwidthPackagesRequest) (<-chan *DescribeBandwidthPackagesResponse, <-chan error) {
responseChan := make(chan *DescribeBandwidthPackagesResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DescribeBandwidthPackages(request)
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

func (client *Client) DescribeBandwidthPackagesWithCallback(request *DescribeBandwidthPackagesRequest, callback func(response *DescribeBandwidthPackagesResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DescribeBandwidthPackagesResponse
var err error
defer close(result)
response, err = client.DescribeBandwidthPackages(request)
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

type DescribeBandwidthPackagesRequest struct {
*requests.RpcRequest
                PageSize  string `position:"Query" name:"PageSize"`
                ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
                PageNumber  string `position:"Query" name:"PageNumber"`
                ResourceOwnerId  string `position:"Query" name:"ResourceOwnerId"`
                OwnerAccount  string `position:"Query" name:"OwnerAccount"`
                OwnerId  string `position:"Query" name:"OwnerId"`
                NatGatewayId  string `position:"Query" name:"NatGatewayId"`
                BandwidthPackageId  string `position:"Query" name:"BandwidthPackageId"`
}


type DescribeBandwidthPackagesResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            TotalCount     requests.Integer `json:"TotalCount" xml:"TotalCount"`
            PageNumber     requests.Integer `json:"PageNumber" xml:"PageNumber"`
            PageSize     requests.Integer `json:"PageSize" xml:"PageSize"`
                BandwidthPackages struct {
                    BandwidthPackage []struct {
            BandwidthPackageId     string `json:"BandwidthPackageId" xml:"BandwidthPackageId"`
            RegionId     string `json:"RegionId" xml:"RegionId"`
            Name     string `json:"Name" xml:"Name"`
            Description     string `json:"Description" xml:"Description"`
            ZoneId     string `json:"ZoneId" xml:"ZoneId"`
            NatGatewayId     string `json:"NatGatewayId" xml:"NatGatewayId"`
            Bandwidth     string `json:"Bandwidth" xml:"Bandwidth"`
            InstanceChargeType     string `json:"InstanceChargeType" xml:"InstanceChargeType"`
            InternetChargeType     string `json:"InternetChargeType" xml:"InternetChargeType"`
            BusinessStatus     string `json:"BusinessStatus" xml:"BusinessStatus"`
            IpCount     string `json:"IpCount" xml:"IpCount"`
            CreationTime     string `json:"CreationTime" xml:"CreationTime"`
            Status     string `json:"Status" xml:"Status"`
            ISP     string `json:"ISP" xml:"ISP"`
                PublicIpAddresses struct {
                    PublicIpAddresse []struct {
            AllocationId     string `json:"AllocationId" xml:"AllocationId"`
            IpAddress     string `json:"IpAddress" xml:"IpAddress"`
            UsingStatus     string `json:"UsingStatus" xml:"UsingStatus"`
            ApAccessEnabled     requests.Boolean `json:"ApAccessEnabled" xml:"ApAccessEnabled"`
                    }   `json:"PublicIpAddresse" xml:"PublicIpAddresse"`
                } `json:"PublicIpAddresses" xml:"PublicIpAddresses"`
                    }   `json:"BandwidthPackage" xml:"BandwidthPackage"`
                } `json:"BandwidthPackages" xml:"BandwidthPackages"`
}

func CreateDescribeBandwidthPackagesRequest() (request *DescribeBandwidthPackagesRequest) {
request = &DescribeBandwidthPackagesRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Vpc", "2016-04-28", "DescribeBandwidthPackages", "", "")
return
}

func CreateDescribeBandwidthPackagesResponse() (response *DescribeBandwidthPackagesResponse) {
response = &DescribeBandwidthPackagesResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

