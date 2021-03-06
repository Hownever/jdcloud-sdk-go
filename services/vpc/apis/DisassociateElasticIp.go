// Copyright 2018 JDCLOUD.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package apis

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
)

type DisassociateElasticIpRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* networkInterface ID  */
    NetworkInterfaceId string `json:"networkInterfaceId"`

    /* 指定解绑的弹性Ip Id (Optional) */
    ElasticIpId *string `json:"elasticIpId"`

    /* 指定解绑的弹性Ip地址 (Optional) */
    ElasticIpAddress *string `json:"elasticIpAddress"`
}

/*
 * param regionId: Region ID 
 * param networkInterfaceId: networkInterface ID 
 * param elasticIpId: 指定解绑的弹性Ip Id (Optional)
 * param elasticIpAddress: 指定解绑的弹性Ip地址 (Optional)
 */
func NewDisassociateElasticIpRequest(
    regionId string,
    networkInterfaceId string,
) *DisassociateElasticIpRequest {

	return &DisassociateElasticIpRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/networkInterfaces/{networkInterfaceId}:disassociateElasticIp",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        NetworkInterfaceId: networkInterfaceId,
	}
}

func (r *DisassociateElasticIpRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

func (r *DisassociateElasticIpRequest) SetNetworkInterfaceId(networkInterfaceId string) {
    r.NetworkInterfaceId = networkInterfaceId
}

func (r *DisassociateElasticIpRequest) SetElasticIpId(elasticIpId string) {
    r.ElasticIpId = &elasticIpId
}

func (r *DisassociateElasticIpRequest) SetElasticIpAddress(elasticIpAddress string) {
    r.ElasticIpAddress = &elasticIpAddress
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DisassociateElasticIpRequest) GetRegionId() string {
    return r.RegionId
}

type DisassociateElasticIpResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DisassociateElasticIpResult `json:"result"`
}

type DisassociateElasticIpResult struct {
}