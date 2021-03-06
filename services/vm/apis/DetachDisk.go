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

type DetachDiskRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* Instance ID  */
    InstanceId string `json:"instanceId"`

    /* 云硬盘ID  */
    DiskId string `json:"diskId"`

    /* 强制缷载，默认False (Optional) */
    Force *bool `json:"force"`
}

/*
 * param regionId: Region ID 
 * param instanceId: Instance ID 
 * param diskId: 云硬盘ID 
 * param force: 强制缷载，默认False (Optional)
 */
func NewDetachDiskRequest(
    regionId string,
    instanceId string,
    diskId string,
) *DetachDiskRequest {

	return &DetachDiskRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}:detachDisk",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        DiskId: diskId,
	}
}

func (r *DetachDiskRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

func (r *DetachDiskRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

func (r *DetachDiskRequest) SetDiskId(diskId string) {
    r.DiskId = diskId
}

func (r *DetachDiskRequest) SetForce(force bool) {
    r.Force = &force
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DetachDiskRequest) GetRegionId() string {
    return r.RegionId
}

type DetachDiskResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DetachDiskResult `json:"result"`
}

type DetachDiskResult struct {
}