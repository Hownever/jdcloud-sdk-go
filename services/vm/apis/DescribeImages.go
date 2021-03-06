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
    vm "github.com/jdcloud-api/jdcloud-sdk-go/services/vm/models"
)

type DescribeImagesRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 镜像来源：public、shared、thirdparty、private，如果没有指定ids参数，此参数必传 (Optional) */
    ImageSource *string `json:"imageSource"`

    /* 操作系统平台: Windows Server、CentOS、Ubuntu (Optional) */
    Platform *string `json:"platform"`

    /* 镜像ID列表，如果指定了此参数，其它参数可为空 (Optional) */
    Ids []string `json:"ids"`
}

/*
 * param regionId: Region ID 
 * param imageSource: 镜像来源：public、shared、thirdparty、private，如果没有指定ids参数，此参数必传 (Optional)
 * param platform: 操作系统平台: Windows Server、CentOS、Ubuntu (Optional)
 * param ids: 镜像ID列表，如果指定了此参数，其它参数可为空 (Optional)
 */
func NewDescribeImagesRequest(
    regionId string,
) *DescribeImagesRequest {

	return &DescribeImagesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/images",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

func (r *DescribeImagesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

func (r *DescribeImagesRequest) SetImageSource(imageSource string) {
    r.ImageSource = &imageSource
}

func (r *DescribeImagesRequest) SetPlatform(platform string) {
    r.Platform = &platform
}

func (r *DescribeImagesRequest) SetIds(ids []string) {
    r.Ids = ids
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeImagesRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeImagesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeImagesResult `json:"result"`
}

type DescribeImagesResult struct {
    Images []vm.Image `json:"images"`
    TotalCount int `json:"totalCount"`
}