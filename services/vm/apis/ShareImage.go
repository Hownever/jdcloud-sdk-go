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

type ShareImageRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* Image ID  */
    ImageId string `json:"imageId"`

    /* 需要共享的帐户 (Optional) */
    Pins []string `json:"pins"`
}

/*
 * param regionId: Region ID 
 * param imageId: Image ID 
 * param pins: 需要共享的帐户 (Optional)
 */
func NewShareImageRequest(
    regionId string,
    imageId string,
) *ShareImageRequest {

	return &ShareImageRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/images/{imageId}:share",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ImageId: imageId,
	}
}

func (r *ShareImageRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

func (r *ShareImageRequest) SetImageId(imageId string) {
    r.ImageId = imageId
}

func (r *ShareImageRequest) SetPins(pins []string) {
    r.Pins = pins
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ShareImageRequest) GetRegionId() string {
    return r.RegionId
}

type ShareImageResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ShareImageResult `json:"result"`
}

type ShareImageResult struct {
}