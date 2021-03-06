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

type AttachDiskRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* Instance ID  */
    InstanceId string `json:"instanceId"`

    /* 云硬盘ID  */
    DiskId string `json:"diskId"`

    /* 逻辑挂载点[vdb,vdc,vdd,vde,vdf,vdg,vdh]  */
    DeviceName string `json:"deviceName"`

    /* 当删除主机时，是否自动关联删除此硬盘，默认False，只支持按配置计费 (Optional) */
    AutoDelete *bool `json:"autoDelete"`
}

/*
 * param regionId: Region ID 
 * param instanceId: Instance ID 
 * param diskId: 云硬盘ID 
 * param deviceName: 逻辑挂载点[vdb,vdc,vdd,vde,vdf,vdg,vdh] 
 * param autoDelete: 当删除主机时，是否自动关联删除此硬盘，默认False，只支持按配置计费 (Optional)
 */
func NewAttachDiskRequest(
    regionId string,
    instanceId string,
    diskId string,
    deviceName string,
) *AttachDiskRequest {

	return &AttachDiskRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}:attachDisk",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        DiskId: diskId,
        DeviceName: deviceName,
	}
}

func (r *AttachDiskRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

func (r *AttachDiskRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

func (r *AttachDiskRequest) SetDiskId(diskId string) {
    r.DiskId = diskId
}

func (r *AttachDiskRequest) SetDeviceName(deviceName string) {
    r.DeviceName = deviceName
}

func (r *AttachDiskRequest) SetAutoDelete(autoDelete bool) {
    r.AutoDelete = &autoDelete
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AttachDiskRequest) GetRegionId() string {
    return r.RegionId
}

type AttachDiskResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AttachDiskResult `json:"result"`
}

type AttachDiskResult struct {
}