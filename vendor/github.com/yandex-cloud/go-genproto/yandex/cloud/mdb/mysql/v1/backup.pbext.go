// Code generated by protoc-gen-goext. DO NOT EDIT.

package mysql

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

func (m *Backup) SetId(v string) {
	m.Id = v
}

func (m *Backup) SetFolderId(v string) {
	m.FolderId = v
}

func (m *Backup) SetCreatedAt(v *timestamp.Timestamp) {
	m.CreatedAt = v
}

func (m *Backup) SetSourceClusterId(v string) {
	m.SourceClusterId = v
}

func (m *Backup) SetStartedAt(v *timestamp.Timestamp) {
	m.StartedAt = v
}
