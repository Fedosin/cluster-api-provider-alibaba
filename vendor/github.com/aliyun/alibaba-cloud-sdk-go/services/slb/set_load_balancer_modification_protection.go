package slb

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

// SetLoadBalancerModificationProtection invokes the slb.SetLoadBalancerModificationProtection API synchronously
func (client *Client) SetLoadBalancerModificationProtection(request *SetLoadBalancerModificationProtectionRequest) (response *SetLoadBalancerModificationProtectionResponse, err error) {
	response = CreateSetLoadBalancerModificationProtectionResponse()
	err = client.DoAction(request, response)
	return
}

// SetLoadBalancerModificationProtectionWithChan invokes the slb.SetLoadBalancerModificationProtection API asynchronously
func (client *Client) SetLoadBalancerModificationProtectionWithChan(request *SetLoadBalancerModificationProtectionRequest) (<-chan *SetLoadBalancerModificationProtectionResponse, <-chan error) {
	responseChan := make(chan *SetLoadBalancerModificationProtectionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetLoadBalancerModificationProtection(request)
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

// SetLoadBalancerModificationProtectionWithCallback invokes the slb.SetLoadBalancerModificationProtection API asynchronously
func (client *Client) SetLoadBalancerModificationProtectionWithCallback(request *SetLoadBalancerModificationProtectionRequest, callback func(response *SetLoadBalancerModificationProtectionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetLoadBalancerModificationProtectionResponse
		var err error
		defer close(result)
		response, err = client.SetLoadBalancerModificationProtection(request)
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

// SetLoadBalancerModificationProtectionRequest is the request struct for api SetLoadBalancerModificationProtection
type SetLoadBalancerModificationProtectionRequest struct {
	*requests.RpcRequest
	ModificationProtectionReason string           `position:"Query" name:"ModificationProtectionReason"`
	ResourceOwnerId              requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount         string           `position:"Query" name:"ResourceOwnerAccount"`
	ModificationProtectionStatus string           `position:"Query" name:"ModificationProtectionStatus"`
	OwnerAccount                 string           `position:"Query" name:"OwnerAccount"`
	OwnerId                      requests.Integer `position:"Query" name:"OwnerId"`
	LoadBalancerId               string           `position:"Query" name:"LoadBalancerId"`
}

// SetLoadBalancerModificationProtectionResponse is the response struct for api SetLoadBalancerModificationProtection
type SetLoadBalancerModificationProtectionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetLoadBalancerModificationProtectionRequest creates a request to invoke SetLoadBalancerModificationProtection API
func CreateSetLoadBalancerModificationProtectionRequest() (request *SetLoadBalancerModificationProtectionRequest) {
	request = &SetLoadBalancerModificationProtectionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "SetLoadBalancerModificationProtection", "slb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSetLoadBalancerModificationProtectionResponse creates a response to parse from SetLoadBalancerModificationProtection response
func CreateSetLoadBalancerModificationProtectionResponse() (response *SetLoadBalancerModificationProtectionResponse) {
	response = &SetLoadBalancerModificationProtectionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
