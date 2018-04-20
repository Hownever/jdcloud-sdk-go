// Copyright 2018-2025 JDCLOUD.COM
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
    . "github.com/jdcloud-api/jdcloud-sdk-go/core"
    "reflect"
)

type ModifyInstanceSpecRequest struct {

    JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* Instance ID  */
    InstanceId string `json:"instanceId"`

    /* 实例规格，包年包月不允许小于当前规格。  */
    InstanceClass string `json:"instanceClass"`

    /* 存储空间，包年包月不允许小于当前规格。  */
    InstanceStorageGB int `json:"instanceStorageGB"`
}

/*
 * param regionId: Region ID 
 * param instanceId: Instance ID 
 * param instanceClass: 实例规格，包年包月不允许小于当前规格。 
 * param instanceStorageGB: 存储空间，包年包月不允许小于当前规格。 
 */
func NewModifyInstanceSpecRequest(
    regionId string,
    instanceId string,
    instanceClass string,
    instanceStorageGB int,
) *ModifyInstanceSpecRequest {

	return &ModifyInstanceSpecRequest{
        JDCloudRequest: JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}:modifyInstanceSpec",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        InstanceClass: instanceClass,
        InstanceStorageGB: instanceStorageGB,
	}
}

func (r *ModifyInstanceSpecRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

func (r *ModifyInstanceSpecRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

func (r *ModifyInstanceSpecRequest) SetInstanceClass(instanceClass string) {
    r.InstanceClass = instanceClass
}

func (r *ModifyInstanceSpecRequest) SetInstanceStorageGB(instanceStorageGB int) {
    r.InstanceStorageGB = instanceStorageGB
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyInstanceSpecRequest) GetRegionId() string {
    fieldName := "RegionId"
    reqType := reflect.TypeOf(r)
    value := reflect.ValueOf(r)
    _, ok := reqType.FieldByName(fieldName)
    if ok {
        return value.FieldByName(fieldName).String()
    }

    return ""
}

type ModifyInstanceSpecResponse struct {
    RequestID string `json:"requestId"`
    Error ErrorResponse `json:"error"`
    Result ModifyInstanceSpecResult `json:"result"`
}

type ModifyInstanceSpecResult struct {
    InstanceId string `json:"instanceId"`
    OrderId string `json:"orderId"`
}