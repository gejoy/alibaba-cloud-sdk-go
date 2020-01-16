package r_kvstore

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

// DeleteDedicatedUserCluster invokes the r_kvstore.DeleteDedicatedUserCluster API synchronously
// api document: https://help.aliyun.com/api/r-kvstore/deletededicatedusercluster.html
func (client *Client) DeleteDedicatedUserCluster(request *DeleteDedicatedUserClusterRequest) (response *DeleteDedicatedUserClusterResponse, err error) {
	response = CreateDeleteDedicatedUserClusterResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteDedicatedUserClusterWithChan invokes the r_kvstore.DeleteDedicatedUserCluster API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/deletededicatedusercluster.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteDedicatedUserClusterWithChan(request *DeleteDedicatedUserClusterRequest) (<-chan *DeleteDedicatedUserClusterResponse, <-chan error) {
	responseChan := make(chan *DeleteDedicatedUserClusterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteDedicatedUserCluster(request)
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

// DeleteDedicatedUserClusterWithCallback invokes the r_kvstore.DeleteDedicatedUserCluster API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/deletededicatedusercluster.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteDedicatedUserClusterWithCallback(request *DeleteDedicatedUserClusterRequest, callback func(response *DeleteDedicatedUserClusterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteDedicatedUserClusterResponse
		var err error
		defer close(result)
		response, err = client.DeleteDedicatedUserCluster(request)
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

// DeleteDedicatedUserClusterRequest is the request struct for api DeleteDedicatedUserCluster
type DeleteDedicatedUserClusterRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	Engine               string           `position:"Query" name:"Engine"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	ClusterId            string           `position:"Query" name:"ClusterId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ZoneId               string           `position:"Query" name:"ZoneId"`
}

// DeleteDedicatedUserClusterResponse is the response struct for api DeleteDedicatedUserCluster
type DeleteDedicatedUserClusterResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteDedicatedUserClusterRequest creates a request to invoke DeleteDedicatedUserCluster API
func CreateDeleteDedicatedUserClusterRequest() (request *DeleteDedicatedUserClusterRequest) {
	request = &DeleteDedicatedUserClusterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "DeleteDedicatedUserCluster", "R-kvstore", "openAPI")
	return
}

// CreateDeleteDedicatedUserClusterResponse creates a response to parse from DeleteDedicatedUserCluster response
func CreateDeleteDedicatedUserClusterResponse() (response *DeleteDedicatedUserClusterResponse) {
	response = &DeleteDedicatedUserClusterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
