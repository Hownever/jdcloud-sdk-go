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
    oss "github.com/jdcloud-api/jdcloud-sdk-go/services/oss/models"
)

type ListBucketsRequest struct {

    core.JDCloudRequest

    /* Region ID，例如：cn-north-1  */
    RegionId string `json:"regionId"`
}

/*
 * param regionId: Region ID，例如：cn-north-1 
 */
func NewListBucketsRequest(
    regionId string,
) *ListBucketsRequest {

	return &ListBucketsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/buckets",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

func (r *ListBucketsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ListBucketsRequest) GetRegionId() string {
    return r.RegionId
}

type ListBucketsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ListBucketsResult `json:"result"`
}

type ListBucketsResult struct {
    Owner oss.User `json:"owner"`
    Buckets []oss.Bucket `json:"buckets"`
}