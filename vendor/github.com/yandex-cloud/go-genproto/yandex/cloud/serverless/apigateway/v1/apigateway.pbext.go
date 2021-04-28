// Code generated by protoc-gen-goext. DO NOT EDIT.

package apigateway

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

func (m *ApiGateway) SetId(v string) {
	m.Id = v
}

func (m *ApiGateway) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ApiGateway) SetCreatedAt(v *timestamp.Timestamp) {
	m.CreatedAt = v
}

func (m *ApiGateway) SetName(v string) {
	m.Name = v
}

func (m *ApiGateway) SetDescription(v string) {
	m.Description = v
}

func (m *ApiGateway) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *ApiGateway) SetStatus(v ApiGateway_Status) {
	m.Status = v
}

func (m *ApiGateway) SetDomain(v string) {
	m.Domain = v
}

func (m *ApiGateway) SetLogGroupId(v string) {
	m.LogGroupId = v
}
