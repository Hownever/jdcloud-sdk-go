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

package client

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
    ipanti "github.com/jdcloud-api/jdcloud-sdk-go/services/ipanti/apis"
    "encoding/json"
    "errors"
)

type IpantiClient struct {
    core.JDCloudClient
}

func NewIpantiClient(credential *core.Credential) *IpantiClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("ipanti.jdcloud-api.com")

    return &IpantiClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "ipanti",
            Revision:    "1.0.0",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *IpantiClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *IpantiClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

/* 关闭实例CC防护 */
func (c *IpantiClient) DisableInstanceCC(request *ipanti.DisableInstanceCCRequest) (*ipanti.DisableInstanceCCResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ipanti.DisableInstanceCCResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 查询某条网站类规则 */
func (c *IpantiClient) DescribeWebRule(request *ipanti.DescribeWebRuleRequest) (*ipanti.DescribeWebRuleResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ipanti.DescribeWebRuleResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 非网站类规则切换成防御状态 */
func (c *IpantiClient) SwitchForwardRuleProtect(request *ipanti.SwitchForwardRuleProtectRequest) (*ipanti.SwitchForwardRuleProtectResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ipanti.SwitchForwardRuleProtectResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 非网站类规则切换成回源状态 */
func (c *IpantiClient) SwitchForwardRuleOrigin(request *ipanti.SwitchForwardRuleOriginRequest) (*ipanti.SwitchForwardRuleOriginResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ipanti.SwitchForwardRuleOriginResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 添加非网站类规则 */
func (c *IpantiClient) CreateForwardRule(request *ipanti.CreateForwardRuleRequest) (*ipanti.CreateForwardRuleResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ipanti.CreateForwardRuleResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 修改实例名称 */
func (c *IpantiClient) ModifyInstanceName(request *ipanti.ModifyInstanceNameRequest) (*ipanti.ModifyInstanceNameResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ipanti.ModifyInstanceNameResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 启用实例url白名单 */
func (c *IpantiClient) EnableInstanceUrlWhiteList(request *ipanti.EnableInstanceUrlWhiteListRequest) (*ipanti.EnableInstanceUrlWhiteListResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ipanti.EnableInstanceUrlWhiteListResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 关闭CC防护每ip的限速 */
func (c *IpantiClient) DisableCcObserverMode(request *ipanti.DisableCcObserverModeRequest) (*ipanti.DisableCcObserverModeResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ipanti.DisableCcObserverModeResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 开启CC防护每ip的限速 */
func (c *IpantiClient) EnableCcObserverMode(request *ipanti.EnableCcObserverModeRequest) (*ipanti.EnableCcObserverModeResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ipanti.EnableCcObserverModeResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 禁用实例ip白名单 */
func (c *IpantiClient) DisableInstanceIpWhiteList(request *ipanti.DisableInstanceIpWhiteListRequest) (*ipanti.DisableInstanceIpWhiteListResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ipanti.DisableInstanceIpWhiteListResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 查询实例 */
func (c *IpantiClient) DescribeInstance(request *ipanti.DescribeInstanceRequest) (*ipanti.DescribeInstanceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ipanti.DescribeInstanceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 查询某个实例下的非网站转发配置 */
func (c *IpantiClient) DescribeForwardRules(request *ipanti.DescribeForwardRulesRequest) (*ipanti.DescribeForwardRulesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ipanti.DescribeForwardRulesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* ddos防护报表 */
func (c *IpantiClient) DdosGraph(request *ipanti.DdosGraphRequest) (*ipanti.DdosGraphResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ipanti.DdosGraphResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 设置实例CC防护 */
func (c *IpantiClient) ModifyInstanceCC(request *ipanti.ModifyInstanceCCRequest) (*ipanti.ModifyInstanceCCResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ipanti.ModifyInstanceCCResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 查询某个实例下的网站类规则 */
func (c *IpantiClient) DescribeWebRules(request *ipanti.DescribeWebRulesRequest) (*ipanti.DescribeWebRulesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ipanti.DescribeWebRulesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 开启实例CC防护 */
func (c *IpantiClient) EnableInstanceCC(request *ipanti.EnableInstanceCCRequest) (*ipanti.EnableInstanceCCResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ipanti.EnableInstanceCCResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 禁用实例ip黑名单 */
func (c *IpantiClient) DisableInstanceIpBlackList(request *ipanti.DisableInstanceIpBlackListRequest) (*ipanti.DisableInstanceIpBlackListResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ipanti.DisableInstanceIpBlackListResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 更新某条非网站类规则 */
func (c *IpantiClient) ModifyForwardRule(request *ipanti.ModifyForwardRuleRequest) (*ipanti.ModifyForwardRuleResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ipanti.ModifyForwardRuleResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 网站类规则开启CC */
func (c *IpantiClient) EnableWebRuleCC(request *ipanti.EnableWebRuleCCRequest) (*ipanti.EnableWebRuleCCResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ipanti.EnableWebRuleCCResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 删除某条网站规则 */
func (c *IpantiClient) DeleteWebRule(request *ipanti.DeleteWebRuleRequest) (*ipanti.DeleteWebRuleResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ipanti.DeleteWebRuleResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 创建实例 */
func (c *IpantiClient) CreateInstance(request *ipanti.CreateInstanceRequest) (*ipanti.CreateInstanceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ipanti.CreateInstanceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 设置实例url白名单 */
func (c *IpantiClient) ModifyInstanceUrlWhiteList(request *ipanti.ModifyInstanceUrlWhiteListRequest) (*ipanti.ModifyInstanceUrlWhiteListResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ipanti.ModifyInstanceUrlWhiteListResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 添加网站类规则 */
func (c *IpantiClient) CreateWebRule(request *ipanti.CreateWebRuleRequest) (*ipanti.CreateWebRuleResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ipanti.CreateWebRuleResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 网站类规则切换成防御状态 */
func (c *IpantiClient) SwitchWebRuleProtect(request *ipanti.SwitchWebRuleProtectRequest) (*ipanti.SwitchWebRuleProtectResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ipanti.SwitchWebRuleProtectResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 查询某条非网站类规则 */
func (c *IpantiClient) DescribeForwardRule(request *ipanti.DescribeForwardRuleRequest) (*ipanti.DescribeForwardRuleResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ipanti.DescribeForwardRuleResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 网站类规则切换成回源状态 */
func (c *IpantiClient) SwitchWebRuleOrigin(request *ipanti.SwitchWebRuleOriginRequest) (*ipanti.SwitchWebRuleOriginResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ipanti.SwitchWebRuleOriginResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 转发流量报表 */
func (c *IpantiClient) FwdGraph(request *ipanti.FwdGraphRequest) (*ipanti.FwdGraphResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ipanti.FwdGraphResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 设置实例ip白名单 */
func (c *IpantiClient) ModifyInstanceIpWhiteList(request *ipanti.ModifyInstanceIpWhiteListRequest) (*ipanti.ModifyInstanceIpWhiteListResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ipanti.ModifyInstanceIpWhiteListResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 查询DDos攻击日志 */
func (c *IpantiClient) DescribeDDosAttackLogs(request *ipanti.DescribeDDosAttackLogsRequest) (*ipanti.DescribeDDosAttackLogsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ipanti.DescribeDDosAttackLogsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 设置实例ip黑名单 */
func (c *IpantiClient) ModifyInstanceIpBlackList(request *ipanti.ModifyInstanceIpBlackListRequest) (*ipanti.ModifyInstanceIpBlackListResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ipanti.ModifyInstanceIpBlackListResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 网站类规则禁用CC */
func (c *IpantiClient) DisableWebRuleCC(request *ipanti.DisableWebRuleCCRequest) (*ipanti.DisableWebRuleCCResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ipanti.DisableWebRuleCCResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 禁用实例url白名单 */
func (c *IpantiClient) DisableInstanceUrlWhiteList(request *ipanti.DisableInstanceUrlWhiteListRequest) (*ipanti.DisableInstanceUrlWhiteListResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ipanti.DisableInstanceUrlWhiteListResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 启用实例ip白名单 */
func (c *IpantiClient) EnableInstanceIpWhiteList(request *ipanti.EnableInstanceIpWhiteListRequest) (*ipanti.EnableInstanceIpWhiteListResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ipanti.EnableInstanceIpWhiteListResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 转发流量报表 */
func (c *IpantiClient) CcGraph(request *ipanti.CcGraphRequest) (*ipanti.CcGraphResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ipanti.CcGraphResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 删除某条非网站规则 */
func (c *IpantiClient) DeleteForwardRule(request *ipanti.DeleteForwardRuleRequest) (*ipanti.DeleteForwardRuleResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ipanti.DeleteForwardRuleResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 更新某条网站类规则 */
func (c *IpantiClient) ModifyWebRule(request *ipanti.ModifyWebRuleRequest) (*ipanti.ModifyWebRuleResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ipanti.ModifyWebRuleResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 设置实例CC防护每ip限速 */
func (c *IpantiClient) SetCcIpLimit(request *ipanti.SetCcIpLimitRequest) (*ipanti.SetCcIpLimitResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ipanti.SetCcIpLimitResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 查询cc攻击日志 */
func (c *IpantiClient) DescribeCcAttackLogs(request *ipanti.DescribeCcAttackLogsRequest) (*ipanti.DescribeCcAttackLogsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ipanti.DescribeCcAttackLogsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 启用实例ip黑名单 */
func (c *IpantiClient) EnableInstanceIpBlackList(request *ipanti.EnableInstanceIpBlackListRequest) (*ipanti.EnableInstanceIpBlackListResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ipanti.EnableInstanceIpBlackListResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 查询实例列表 */
func (c *IpantiClient) DescribeInstances(request *ipanti.DescribeInstancesRequest) (*ipanti.DescribeInstancesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ipanti.DescribeInstancesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 查询cc攻击日志详情 */
func (c *IpantiClient) DescribeCcAttackLogDetails(request *ipanti.DescribeCcAttackLogDetailsRequest) (*ipanti.DescribeCcAttackLogDetailsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ipanti.DescribeCcAttackLogDetailsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

