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
    vpc "github.com/jdcloud-api/jdcloud-sdk-go/services/vpc/models"
    common "github.com/jdcloud-api/jdcloud-sdk-go/services/common/models"
)

type DescribeElasticIpsRequest struct {

    JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 页码, 默认为1, 取值范围：[1,∞), 页码超过总页数时, 显示最后一页 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小，默认为20，取值范围：[10,500] (Optional) */
    PageSize *int `json:"pageSize"`

    /* elasticIpIds - elasticip id数组条件
elasticIpAddress - eip的IP地址 
chargeStatus	- eip的费用支付状态,normal(正常状态) or overdue(预付费已到期) or arrear(欠费状态)
 (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 * param regionId: Region ID 
 * param pageNumber: 页码, 默认为1, 取值范围：[1,∞), 页码超过总页数时, 显示最后一页 (Optional)
 * param pageSize: 分页大小，默认为20，取值范围：[10,500] (Optional)
 * param filters: elasticIpIds - elasticip id数组条件
elasticIpAddress - eip的IP地址 
chargeStatus	- eip的费用支付状态,normal(正常状态) or overdue(预付费已到期) or arrear(欠费状态)
 (Optional)
 */
func NewDescribeElasticIpsRequest(
    regionId string,
) *DescribeElasticIpsRequest {

	return &DescribeElasticIpsRequest{
        JDCloudRequest: JDCloudRequest{
			URL:     "/regions/{regionId}/elasticIps/",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

func (r *DescribeElasticIpsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

func (r *DescribeElasticIpsRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

func (r *DescribeElasticIpsRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

func (r *DescribeElasticIpsRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeElasticIpsRequest) GetRegionId() string {
    fieldName := "RegionId"
    reqType := reflect.TypeOf(r)
    value := reflect.ValueOf(r)
    _, ok := reqType.FieldByName(fieldName)
    if ok {
        return value.FieldByName(fieldName).String()
    }

    return ""
}

type DescribeElasticIpsResponse struct {
    RequestID string `json:"requestId"`
    Error ErrorResponse `json:"error"`
    Result DescribeElasticIpsResult `json:"result"`
}

type DescribeElasticIpsResult struct {
    ElasticIps []vpc.ElasticIp `json:"elasticIps"`
    TotalCount int `json:"totalCount"`
}