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

// CreateSnatEntry invokes the vpc.CreateSnatEntry API synchronously
// api document: https://help.aliyun.com/api/vpc/createsnatentry.html
func (client *Client) CreateSnatEntry(request *CreateSnatEntryRequest) (response *CreateSnatEntryResponse, err error) {
	response = CreateCreateSnatEntryResponse()
	err = client.DoAction(request, response)
	return
}

// CreateSnatEntryWithChan invokes the vpc.CreateSnatEntry API asynchronously
// api document: https://help.aliyun.com/api/vpc/createsnatentry.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateSnatEntryWithChan(request *CreateSnatEntryRequest) (<-chan *CreateSnatEntryResponse, <-chan error) {
	responseChan := make(chan *CreateSnatEntryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateSnatEntry(request)
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

// CreateSnatEntryWithCallback invokes the vpc.CreateSnatEntry API asynchronously
// api document: https://help.aliyun.com/api/vpc/createsnatentry.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateSnatEntryWithCallback(request *CreateSnatEntryRequest, callback func(response *CreateSnatEntryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateSnatEntryResponse
		var err error
		defer close(result)
		response, err = client.CreateSnatEntry(request)
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

// CreateSnatEntryRequest is the request struct for api CreateSnatEntry
type CreateSnatEntryRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	SnatTableId          string           `position:"Query" name:"SnatTableId"`
	SourceVSwitchId      string           `position:"Query" name:"SourceVSwitchId"`
	SourceCIDR           string           `position:"Query" name:"SourceCIDR"`
	SnatIp               string           `position:"Query" name:"SnatIp"`
}

// CreateSnatEntryResponse is the response struct for api CreateSnatEntry
type CreateSnatEntryResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	SnatEntryId string `json:"SnatEntryId" xml:"SnatEntryId"`
}

// CreateCreateSnatEntryRequest creates a request to invoke CreateSnatEntry API
func CreateCreateSnatEntryRequest() (request *CreateSnatEntryRequest) {
	request = &CreateSnatEntryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "CreateSnatEntry", "vpc", "openAPI")
	return
}

// CreateCreateSnatEntryResponse creates a response to parse from CreateSnatEntry response
func CreateCreateSnatEntryResponse() (response *CreateSnatEntryResponse) {
	response = &CreateSnatEntryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
