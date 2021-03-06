// RAINBOND, Application Management Platform
// Copyright (C) 2014-2017 Goodrain Co., Ltd.

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version. For any non-GPL usage of Rainbond,
// one or multiple Commercial Licenses authorized by Goodrain Co., Ltd.
// must be obtained first.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package model

//HttpRuleStruct -
type HttpRuleStruct struct {
	HttpRuleID      string                 `json:"http_rule_id" validate:"http_rule_id|required"`
	ServiceID       string                 `json:"service_id"`
	ContainerPort   int                    `json:"container_port"`
	Domain          string                 `json:"domain"`
	Path            string                 `json:"path"`
	Header          string                 `json:"header"`
	Cookie          string                 `json:"cookie"`
	IP              string                 `json:"ip"`
	CertificateID   string                 `json:"certificate_id"`
	CertificateName string                 `json:"certificate_name"`
	Certificate     string                 `json:"certificate"`
	PrivateKey      string                 `json:"private_key"`
	RuleExtensions  []*RuleExtensionStruct `json:"rule_extensions"`
}

type TcpRuleStruct struct {
	TcpRuleId      string                 `json:"tcp_rule_id" validate:"tcp_rule_id|required"`
	ServiceID      string                 `json:"service_id"`
	ContainerPort  int                    `json:"container_port"`
	IP             string                 `json:"ip"`
	Port           int                    `json:"port"`
	RuleExtensions []*RuleExtensionStruct `json:"rule_extensions"`
}

type RuleExtensionStruct struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
