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
    monitor "github.com/jdcloud-api/jdcloud-sdk-go/services/monitor/models"
)

type DescribeMetricsRequest struct {

    JDCloudRequest

    /* 资源的类型，取值vm, lb, ip, database 等  */
    ServiceCode string `json:"serviceCode"`
}

/*
 * param serviceCode: 资源的类型，取值vm, lb, ip, database 等 
 */
func NewDescribeMetricsRequest(
    serviceCode string,
) *DescribeMetricsRequest {

	return &DescribeMetricsRequest{
        JDCloudRequest: JDCloudRequest{
			URL:     "/metrics",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        ServiceCode: serviceCode,
	}
}

func (r *DescribeMetricsRequest) SetServiceCode(serviceCode string) {
    r.ServiceCode = serviceCode
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeMetricsRequest) GetRegionId() string {
    fieldName := "RegionId"
    reqType := reflect.TypeOf(r)
    value := reflect.ValueOf(r)
    _, ok := reqType.FieldByName(fieldName)
    if ok {
        return value.FieldByName(fieldName).String()
    }

    return ""
}

type DescribeMetricsResponse struct {
    RequestID string `json:"requestId"`
    Error ErrorResponse `json:"error"`
    Result DescribeMetricsResult `json:"result"`
}

type DescribeMetricsResult struct {
    Metrics []monitor.MetricDetail `json:"metrics"`
}