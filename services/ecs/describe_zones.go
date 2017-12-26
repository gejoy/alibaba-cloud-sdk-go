
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

func (client *Client) DescribeZones(request *DescribeZonesRequest) (response *DescribeZonesResponse, err error) {
response = CreateDescribeZonesResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) DescribeZonesWithChan(request *DescribeZonesRequest) (<-chan *DescribeZonesResponse, <-chan error) {
responseChan := make(chan *DescribeZonesResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DescribeZones(request)
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

func (client *Client) DescribeZonesWithCallback(request *DescribeZonesRequest, callback func(response *DescribeZonesResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DescribeZonesResponse
var err error
defer close(result)
response, err = client.DescribeZones(request)
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

type DescribeZonesRequest struct {
*requests.RpcRequest
                ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
                SpotStrategy  string `position:"Query" name:"SpotStrategy"`
                Verbose  string `position:"Query" name:"Verbose"`
                ResourceOwnerId  string `position:"Query" name:"ResourceOwnerId"`
                OwnerAccount  string `position:"Query" name:"OwnerAccount"`
                OwnerId  string `position:"Query" name:"OwnerId"`
                InstanceChargeType  string `position:"Query" name:"InstanceChargeType"`
}


type DescribeZonesResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
                Zones struct {
                    Zone []struct {
            ZoneId     string `json:"ZoneId" xml:"ZoneId"`
            LocalName     string `json:"LocalName" xml:"LocalName"`
                AvailableResourceCreation struct {
                ResourceTypes []    string `json:"ResourceTypes" xml:"ResourceTypes"`
                } `json:"AvailableResourceCreation" xml:"AvailableResourceCreation"`
                AvailableDiskCategories struct {
                DiskCategories []    string `json:"DiskCategories" xml:"DiskCategories"`
                } `json:"AvailableDiskCategories" xml:"AvailableDiskCategories"`
                AvailableInstanceTypes struct {
                InstanceTypes []    string `json:"InstanceTypes" xml:"InstanceTypes"`
                } `json:"AvailableInstanceTypes" xml:"AvailableInstanceTypes"`
                AvailableVolumeCategories struct {
                VolumeCategories []    string `json:"VolumeCategories" xml:"VolumeCategories"`
                } `json:"AvailableVolumeCategories" xml:"AvailableVolumeCategories"`
                AvailableResources struct {
                    ResourcesInfo []struct {
            IoOptimized     requests.Boolean `json:"IoOptimized" xml:"IoOptimized"`
                SystemDiskCategories struct {
                SupportedSystemDiskCategory []    string `json:"supportedSystemDiskCategory" xml:"supportedSystemDiskCategory"`
                } `json:"SystemDiskCategories" xml:"SystemDiskCategories"`
                DataDiskCategories struct {
                SupportedDataDiskCategory []    string `json:"supportedDataDiskCategory" xml:"supportedDataDiskCategory"`
                } `json:"DataDiskCategories" xml:"DataDiskCategories"`
                NetworkTypes struct {
                SupportedNetworkCategory []    string `json:"supportedNetworkCategory" xml:"supportedNetworkCategory"`
                } `json:"NetworkTypes" xml:"NetworkTypes"`
                InstanceTypes struct {
                SupportedInstanceType []    string `json:"supportedInstanceType" xml:"supportedInstanceType"`
                } `json:"InstanceTypes" xml:"InstanceTypes"`
                InstanceTypeFamilies struct {
                SupportedInstanceTypeFamily []    string `json:"supportedInstanceTypeFamily" xml:"supportedInstanceTypeFamily"`
                } `json:"InstanceTypeFamilies" xml:"InstanceTypeFamilies"`
                InstanceGenerations struct {
                SupportedInstanceGeneration []    string `json:"supportedInstanceGeneration" xml:"supportedInstanceGeneration"`
                } `json:"InstanceGenerations" xml:"InstanceGenerations"`
                    }   `json:"ResourcesInfo" xml:"ResourcesInfo"`
                } `json:"AvailableResources" xml:"AvailableResources"`
                    }   `json:"Zone" xml:"Zone"`
                } `json:"Zones" xml:"Zones"`
}

func CreateDescribeZonesRequest() (request *DescribeZonesRequest) {
request = &DescribeZonesRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeZones", "", "")
return
}

func CreateDescribeZonesResponse() (response *DescribeZonesResponse) {
response = &DescribeZonesResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

