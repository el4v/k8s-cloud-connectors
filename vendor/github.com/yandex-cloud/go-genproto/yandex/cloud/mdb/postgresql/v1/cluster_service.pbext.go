// Code generated by protoc-gen-goext. DO NOT EDIT.

package postgresql

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	config "github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/postgresql/v1/config"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	timeofday "google.golang.org/genproto/googleapis/type/timeofday"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
)

func (m *GetClusterRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *ListClustersRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ListClustersRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListClustersRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListClustersRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListClustersResponse) SetClusters(v []*Cluster) {
	m.Clusters = v
}

func (m *ListClustersResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *CreateClusterRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *CreateClusterRequest) SetName(v string) {
	m.Name = v
}

func (m *CreateClusterRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreateClusterRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *CreateClusterRequest) SetEnvironment(v Cluster_Environment) {
	m.Environment = v
}

func (m *CreateClusterRequest) SetConfigSpec(v *ConfigSpec) {
	m.ConfigSpec = v
}

func (m *CreateClusterRequest) SetDatabaseSpecs(v []*DatabaseSpec) {
	m.DatabaseSpecs = v
}

func (m *CreateClusterRequest) SetUserSpecs(v []*UserSpec) {
	m.UserSpecs = v
}

func (m *CreateClusterRequest) SetHostSpecs(v []*HostSpec) {
	m.HostSpecs = v
}

func (m *CreateClusterRequest) SetNetworkId(v string) {
	m.NetworkId = v
}

func (m *CreateClusterRequest) SetSecurityGroupIds(v []string) {
	m.SecurityGroupIds = v
}

func (m *CreateClusterMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *UpdateClusterRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *UpdateClusterRequest) SetUpdateMask(v *field_mask.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateClusterRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdateClusterRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *UpdateClusterRequest) SetConfigSpec(v *ConfigSpec) {
	m.ConfigSpec = v
}

func (m *UpdateClusterRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdateClusterRequest) SetMaintenanceWindow(v *MaintenanceWindow) {
	m.MaintenanceWindow = v
}

func (m *UpdateClusterRequest) SetSecurityGroupIds(v []string) {
	m.SecurityGroupIds = v
}

func (m *UpdateClusterMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *DeleteClusterRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *DeleteClusterMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *StartClusterRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *StartClusterMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *StopClusterRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *StopClusterMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *MoveClusterRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *MoveClusterRequest) SetDestinationFolderId(v string) {
	m.DestinationFolderId = v
}

func (m *MoveClusterMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *MoveClusterMetadata) SetSourceFolderId(v string) {
	m.SourceFolderId = v
}

func (m *MoveClusterMetadata) SetDestinationFolderId(v string) {
	m.DestinationFolderId = v
}

func (m *BackupClusterRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *BackupClusterMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *RestoreClusterRequest) SetBackupId(v string) {
	m.BackupId = v
}

func (m *RestoreClusterRequest) SetTime(v *timestamp.Timestamp) {
	m.Time = v
}

func (m *RestoreClusterRequest) SetTimeInclusive(v bool) {
	m.TimeInclusive = v
}

func (m *RestoreClusterRequest) SetName(v string) {
	m.Name = v
}

func (m *RestoreClusterRequest) SetDescription(v string) {
	m.Description = v
}

func (m *RestoreClusterRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *RestoreClusterRequest) SetEnvironment(v Cluster_Environment) {
	m.Environment = v
}

func (m *RestoreClusterRequest) SetConfigSpec(v *ConfigSpec) {
	m.ConfigSpec = v
}

func (m *RestoreClusterRequest) SetHostSpecs(v []*HostSpec) {
	m.HostSpecs = v
}

func (m *RestoreClusterRequest) SetNetworkId(v string) {
	m.NetworkId = v
}

func (m *RestoreClusterRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *RestoreClusterRequest) SetSecurityGroupIds(v []string) {
	m.SecurityGroupIds = v
}

func (m *RestoreClusterMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *RestoreClusterMetadata) SetBackupId(v string) {
	m.BackupId = v
}

func (m *StartClusterFailoverRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *StartClusterFailoverRequest) SetHostName(v string) {
	m.HostName = v
}

func (m *StartClusterFailoverMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *RescheduleMaintenanceRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *RescheduleMaintenanceRequest) SetRescheduleType(v RescheduleMaintenanceRequest_RescheduleType) {
	m.RescheduleType = v
}

func (m *RescheduleMaintenanceRequest) SetDelayedUntil(v *timestamp.Timestamp) {
	m.DelayedUntil = v
}

func (m *RescheduleMaintenanceMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *RescheduleMaintenanceMetadata) SetDelayedUntil(v *timestamp.Timestamp) {
	m.DelayedUntil = v
}

func (m *LogRecord) SetTimestamp(v *timestamp.Timestamp) {
	m.Timestamp = v
}

func (m *LogRecord) SetMessage(v map[string]string) {
	m.Message = v
}

func (m *ListClusterLogsRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *ListClusterLogsRequest) SetColumnFilter(v []string) {
	m.ColumnFilter = v
}

func (m *ListClusterLogsRequest) SetServiceType(v ListClusterLogsRequest_ServiceType) {
	m.ServiceType = v
}

func (m *ListClusterLogsRequest) SetFromTime(v *timestamp.Timestamp) {
	m.FromTime = v
}

func (m *ListClusterLogsRequest) SetToTime(v *timestamp.Timestamp) {
	m.ToTime = v
}

func (m *ListClusterLogsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListClusterLogsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListClusterLogsRequest) SetAlwaysNextPageToken(v bool) {
	m.AlwaysNextPageToken = v
}

func (m *ListClusterLogsResponse) SetLogs(v []*LogRecord) {
	m.Logs = v
}

func (m *ListClusterLogsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *StreamLogRecord) SetRecord(v *LogRecord) {
	m.Record = v
}

func (m *StreamLogRecord) SetNextRecordToken(v string) {
	m.NextRecordToken = v
}

func (m *StreamClusterLogsRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *StreamClusterLogsRequest) SetColumnFilter(v []string) {
	m.ColumnFilter = v
}

func (m *StreamClusterLogsRequest) SetServiceType(v StreamClusterLogsRequest_ServiceType) {
	m.ServiceType = v
}

func (m *StreamClusterLogsRequest) SetFromTime(v *timestamp.Timestamp) {
	m.FromTime = v
}

func (m *StreamClusterLogsRequest) SetToTime(v *timestamp.Timestamp) {
	m.ToTime = v
}

func (m *StreamClusterLogsRequest) SetRecordToken(v string) {
	m.RecordToken = v
}

func (m *StreamClusterLogsRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListClusterOperationsRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *ListClusterOperationsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListClusterOperationsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListClusterOperationsResponse) SetOperations(v []*operation.Operation) {
	m.Operations = v
}

func (m *ListClusterOperationsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *ListClusterBackupsRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *ListClusterBackupsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListClusterBackupsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListClusterBackupsResponse) SetBackups(v []*Backup) {
	m.Backups = v
}

func (m *ListClusterBackupsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *ListClusterHostsRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *ListClusterHostsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListClusterHostsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListClusterHostsResponse) SetHosts(v []*Host) {
	m.Hosts = v
}

func (m *ListClusterHostsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *AddClusterHostsRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *AddClusterHostsRequest) SetHostSpecs(v []*HostSpec) {
	m.HostSpecs = v
}

func (m *AddClusterHostsMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *AddClusterHostsMetadata) SetHostNames(v []string) {
	m.HostNames = v
}

func (m *DeleteClusterHostsRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *DeleteClusterHostsRequest) SetHostNames(v []string) {
	m.HostNames = v
}

func (m *DeleteClusterHostsMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *DeleteClusterHostsMetadata) SetHostNames(v []string) {
	m.HostNames = v
}

func (m *UpdateClusterHostsRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *UpdateClusterHostsRequest) SetUpdateHostSpecs(v []*UpdateHostSpec) {
	m.UpdateHostSpecs = v
}

func (m *UpdateClusterHostsMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *UpdateClusterHostsMetadata) SetHostNames(v []string) {
	m.HostNames = v
}

func (m *UpdateHostSpec) SetHostName(v string) {
	m.HostName = v
}

func (m *UpdateHostSpec) SetReplicationSource(v string) {
	m.ReplicationSource = v
}

func (m *UpdateHostSpec) SetPriority(v *wrappers.Int64Value) {
	m.Priority = v
}

func (m *UpdateHostSpec) SetConfigSpec(v *ConfigHostSpec) {
	m.ConfigSpec = v
}

func (m *HostSpec) SetZoneId(v string) {
	m.ZoneId = v
}

func (m *HostSpec) SetSubnetId(v string) {
	m.SubnetId = v
}

func (m *HostSpec) SetAssignPublicIp(v bool) {
	m.AssignPublicIp = v
}

func (m *HostSpec) SetReplicationSource(v string) {
	m.ReplicationSource = v
}

func (m *HostSpec) SetPriority(v *wrappers.Int64Value) {
	m.Priority = v
}

func (m *HostSpec) SetConfigSpec(v *ConfigHostSpec) {
	m.ConfigSpec = v
}

type ConfigSpec_PostgresqlConfig = isConfigSpec_PostgresqlConfig

func (m *ConfigSpec) SetPostgresqlConfig(v ConfigSpec_PostgresqlConfig) {
	m.PostgresqlConfig = v
}

func (m *ConfigSpec) SetVersion(v string) {
	m.Version = v
}

func (m *ConfigSpec) SetPostgresqlConfig_9_6(v *config.PostgresqlConfig9_6) {
	m.PostgresqlConfig = &ConfigSpec_PostgresqlConfig_9_6{
		PostgresqlConfig_9_6: v,
	}
}

func (m *ConfigSpec) SetPostgresqlConfig_10_1C(v *config.PostgresqlConfig10_1C) {
	m.PostgresqlConfig = &ConfigSpec_PostgresqlConfig_10_1C{
		PostgresqlConfig_10_1C: v,
	}
}

func (m *ConfigSpec) SetPostgresqlConfig_10(v *config.PostgresqlConfig10) {
	m.PostgresqlConfig = &ConfigSpec_PostgresqlConfig_10{
		PostgresqlConfig_10: v,
	}
}

func (m *ConfigSpec) SetPostgresqlConfig_11(v *config.PostgresqlConfig11) {
	m.PostgresqlConfig = &ConfigSpec_PostgresqlConfig_11{
		PostgresqlConfig_11: v,
	}
}

func (m *ConfigSpec) SetPostgresqlConfig_11_1C(v *config.PostgresqlConfig11_1C) {
	m.PostgresqlConfig = &ConfigSpec_PostgresqlConfig_11_1C{
		PostgresqlConfig_11_1C: v,
	}
}

func (m *ConfigSpec) SetPostgresqlConfig_12(v *config.PostgresqlConfig12) {
	m.PostgresqlConfig = &ConfigSpec_PostgresqlConfig_12{
		PostgresqlConfig_12: v,
	}
}

func (m *ConfigSpec) SetPostgresqlConfig_12_1C(v *config.PostgresqlConfig12_1C) {
	m.PostgresqlConfig = &ConfigSpec_PostgresqlConfig_12_1C{
		PostgresqlConfig_12_1C: v,
	}
}

func (m *ConfigSpec) SetPostgresqlConfig_13(v *config.PostgresqlConfig13) {
	m.PostgresqlConfig = &ConfigSpec_PostgresqlConfig_13{
		PostgresqlConfig_13: v,
	}
}

func (m *ConfigSpec) SetPoolerConfig(v *ConnectionPoolerConfig) {
	m.PoolerConfig = v
}

func (m *ConfigSpec) SetResources(v *Resources) {
	m.Resources = v
}

func (m *ConfigSpec) SetAutofailover(v *wrappers.BoolValue) {
	m.Autofailover = v
}

func (m *ConfigSpec) SetBackupWindowStart(v *timeofday.TimeOfDay) {
	m.BackupWindowStart = v
}

func (m *ConfigSpec) SetAccess(v *Access) {
	m.Access = v
}

func (m *ConfigSpec) SetPerformanceDiagnostics(v *PerformanceDiagnostics) {
	m.PerformanceDiagnostics = v
}

type ConfigHostSpec_PostgresqlConfig = isConfigHostSpec_PostgresqlConfig

func (m *ConfigHostSpec) SetPostgresqlConfig(v ConfigHostSpec_PostgresqlConfig) {
	m.PostgresqlConfig = v
}

func (m *ConfigHostSpec) SetPostgresqlConfig_9_6(v *config.PostgresqlHostConfig9_6) {
	m.PostgresqlConfig = &ConfigHostSpec_PostgresqlConfig_9_6{
		PostgresqlConfig_9_6: v,
	}
}

func (m *ConfigHostSpec) SetPostgresqlConfig_10_1C(v *config.PostgresqlHostConfig10_1C) {
	m.PostgresqlConfig = &ConfigHostSpec_PostgresqlConfig_10_1C{
		PostgresqlConfig_10_1C: v,
	}
}

func (m *ConfigHostSpec) SetPostgresqlConfig_10(v *config.PostgresqlHostConfig10) {
	m.PostgresqlConfig = &ConfigHostSpec_PostgresqlConfig_10{
		PostgresqlConfig_10: v,
	}
}

func (m *ConfigHostSpec) SetPostgresqlConfig_11(v *config.PostgresqlHostConfig11) {
	m.PostgresqlConfig = &ConfigHostSpec_PostgresqlConfig_11{
		PostgresqlConfig_11: v,
	}
}

func (m *ConfigHostSpec) SetPostgresqlConfig_11_1C(v *config.PostgresqlHostConfig11_1C) {
	m.PostgresqlConfig = &ConfigHostSpec_PostgresqlConfig_11_1C{
		PostgresqlConfig_11_1C: v,
	}
}

func (m *ConfigHostSpec) SetPostgresqlConfig_12(v *config.PostgresqlHostConfig12) {
	m.PostgresqlConfig = &ConfigHostSpec_PostgresqlConfig_12{
		PostgresqlConfig_12: v,
	}
}

func (m *ConfigHostSpec) SetPostgresqlConfig_12_1C(v *config.PostgresqlHostConfig12_1C) {
	m.PostgresqlConfig = &ConfigHostSpec_PostgresqlConfig_12_1C{
		PostgresqlConfig_12_1C: v,
	}
}

func (m *ConfigHostSpec) SetPostgresqlConfig_13(v *config.PostgresqlHostConfig13) {
	m.PostgresqlConfig = &ConfigHostSpec_PostgresqlConfig_13{
		PostgresqlConfig_13: v,
	}
}
