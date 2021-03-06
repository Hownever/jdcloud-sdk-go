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
    streambus "github.com/jdcloud-api/jdcloud-sdk-go/services/streambus/models"
)

type UpdateTopicRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /*   */
    TopicModel *streambus.AddTopic `json:"topicModel"`
}

/*
 * param regionId: Region ID 
 * param topicModel:  
 */
func NewUpdateTopicRequest(
    regionId string,
    topicModel *streambus.AddTopic,
) *UpdateTopicRequest {

	return &UpdateTopicRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/topic",
			Method:  "PUT",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        TopicModel: topicModel,
	}
}

func (r *UpdateTopicRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

func (r *UpdateTopicRequest) SetTopicModel(topicModel *streambus.AddTopic) {
    r.TopicModel = topicModel
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpdateTopicRequest) GetRegionId() string {
    return r.RegionId
}

type UpdateTopicResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpdateTopicResult `json:"result"`
}

type UpdateTopicResult struct {
    Status bool `json:"status"`
    Message string `json:"message"`
}