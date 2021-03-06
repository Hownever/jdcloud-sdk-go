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

type ModifyCacheInstanceAttributeRequest struct {

    core.JDCloudRequest

    /* 缓存Redis实例所在区域的Region ID。目前缓存Redis有华北、华南、华东区域，对应Region ID为cn-north-1、cn-south-1、cn-east-2  */
    RegionId string `json:"regionId"`

    /* 缓存Redis实例ID  */
    CacheInstanceId string `json:"cacheInstanceId"`

    /* 缓存Redis实例资源名称 (Optional) */
    CacheInstanceName *string `json:"cacheInstanceName"`

    /* 缓存Redis实例资源描述 (Optional) */
    CacheInstanceDescription *string `json:"cacheInstanceDescription"`
}

/*
 * param regionId: 缓存Redis实例所在区域的Region ID。目前缓存Redis有华北、华南、华东区域，对应Region ID为cn-north-1、cn-south-1、cn-east-2 
 * param cacheInstanceId: 缓存Redis实例ID 
 * param cacheInstanceName: 缓存Redis实例资源名称 (Optional)
 * param cacheInstanceDescription: 缓存Redis实例资源描述 (Optional)
 */
func NewModifyCacheInstanceAttributeRequest(
    regionId string,
    cacheInstanceId string,
) *ModifyCacheInstanceAttributeRequest {

	return &ModifyCacheInstanceAttributeRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/cacheInstance/{cacheInstanceId}",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        CacheInstanceId: cacheInstanceId,
	}
}

func (r *ModifyCacheInstanceAttributeRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

func (r *ModifyCacheInstanceAttributeRequest) SetCacheInstanceId(cacheInstanceId string) {
    r.CacheInstanceId = cacheInstanceId
}

func (r *ModifyCacheInstanceAttributeRequest) SetCacheInstanceName(cacheInstanceName string) {
    r.CacheInstanceName = &cacheInstanceName
}

func (r *ModifyCacheInstanceAttributeRequest) SetCacheInstanceDescription(cacheInstanceDescription string) {
    r.CacheInstanceDescription = &cacheInstanceDescription
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyCacheInstanceAttributeRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyCacheInstanceAttributeResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyCacheInstanceAttributeResult `json:"result"`
}

type ModifyCacheInstanceAttributeResult struct {
}