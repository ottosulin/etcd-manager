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

// DeleteNetworkInterfacePermission invokes the ecs.DeleteNetworkInterfacePermission API synchronously
// api document: https://help.aliyun.com/api/ecs/deletenetworkinterfacepermission.html
func (client *Client) DeleteNetworkInterfacePermission(request *DeleteNetworkInterfacePermissionRequest) (response *DeleteNetworkInterfacePermissionResponse, err error) {
	response = CreateDeleteNetworkInterfacePermissionResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteNetworkInterfacePermissionWithChan invokes the ecs.DeleteNetworkInterfacePermission API asynchronously
// api document: https://help.aliyun.com/api/ecs/deletenetworkinterfacepermission.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteNetworkInterfacePermissionWithChan(request *DeleteNetworkInterfacePermissionRequest) (<-chan *DeleteNetworkInterfacePermissionResponse, <-chan error) {
	responseChan := make(chan *DeleteNetworkInterfacePermissionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteNetworkInterfacePermission(request)
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

// DeleteNetworkInterfacePermissionWithCallback invokes the ecs.DeleteNetworkInterfacePermission API asynchronously
// api document: https://help.aliyun.com/api/ecs/deletenetworkinterfacepermission.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteNetworkInterfacePermissionWithCallback(request *DeleteNetworkInterfacePermissionRequest, callback func(response *DeleteNetworkInterfacePermissionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteNetworkInterfacePermissionResponse
		var err error
		defer close(result)
		response, err = client.DeleteNetworkInterfacePermission(request)
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

// DeleteNetworkInterfacePermissionRequest is the request struct for api DeleteNetworkInterfacePermission
type DeleteNetworkInterfacePermissionRequest struct {
	*requests.RpcRequest
	ResourceOwnerId              requests.Integer `position:"Query" name:"ResourceOwnerId"`
	NetworkInterfacePermissionId string           `position:"Query" name:"NetworkInterfacePermissionId"`
	ResourceOwnerAccount         string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                 string           `position:"Query" name:"OwnerAccount"`
	OwnerId                      requests.Integer `position:"Query" name:"OwnerId"`
	Force                        requests.Boolean `position:"Query" name:"Force"`
}

// DeleteNetworkInterfacePermissionResponse is the response struct for api DeleteNetworkInterfacePermission
type DeleteNetworkInterfacePermissionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteNetworkInterfacePermissionRequest creates a request to invoke DeleteNetworkInterfacePermission API
func CreateDeleteNetworkInterfacePermissionRequest() (request *DeleteNetworkInterfacePermissionRequest) {
	request = &DeleteNetworkInterfacePermissionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DeleteNetworkInterfacePermission", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteNetworkInterfacePermissionResponse creates a response to parse from DeleteNetworkInterfacePermission response
func CreateDeleteNetworkInterfacePermissionResponse() (response *DeleteNetworkInterfacePermissionResponse) {
	response = &DeleteNetworkInterfacePermissionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
