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

package models


type Alarm struct {

    /* 统计方法：平均值=avg、最大值=max、最小值=min (Optional) */
    Calculation string `json:"calculation"`

    /* 通知的联系组，如 [“联系组1”,”联系组2”] (Optional) */
    ContactGroups []string `json:"contactGroups"`

    /* 通知的联系人，如 [“联系人1”,”联系人2”] (Optional) */
    ContactPersons []string `json:"contactPersons"`

    /* 创建时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 启用禁用 1启用，0禁用 (Optional) */
    Enabled int `json:"enabled"`

    /* 规则id (Optional) */
    Id string `json:"id"`

    /* 监控项 (Optional) */
    Metric string `json:"metric"`

    /* 规则id监控项名称 (Optional) */
    MetricName string `json:"metricName"`

    /* 通知周期 单位：小时 (Optional) */
    NoticePeriod int `json:"noticePeriod"`

    /* 报警的时间  , 查询正在报警规则时该字段有效 (Optional) */
    NoticeTime string `json:"noticeTime"`

    /* >=、>、<、<=、==、！= (Optional) */
    Operation string `json:"operation"`

    /* 统计周期（单位：分钟） (Optional) */
    Period int `json:"period"`

    /* 地域信息 (Optional) */
    Region string `json:"region"`

    /* 此规则所应用的资源id (Optional) */
    ResourceId string `json:"resourceId"`

    /* 报警规则对应的产品 (Optional) */
    ServiceCode string `json:"serviceCode"`

    /* 监控项状态:1正常，2告警，4数据不足 (Optional) */
    Status int `json:"status"`

    /* 监控项附属信息 (Optional) */
    Tag string `json:"tag"`

    /* 阈值 (Optional) */
    Threshold float64 `json:"threshold"`

    /* 连续多少次后报警 (Optional) */
    Times int `json:"times"`

    /* 报警值 , 查询正在报警规则时该字段有效 (Optional) */
    Value float64 `json:"value"`
}