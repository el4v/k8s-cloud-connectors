package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type YandexHostClassRedis string

// YandexDiskType type for yandex disk type string
// +kubebuilder:validation:Enum=network-ssd;local-ssd
type YandexDiskType string

type YandexDiskSize int

// YandexZone type for yandex zone string
// +kubebuilder:validation:Enum=ru-central1-a;ru-central1-b;ru-central1-c
type YandexZone string

type YandexSubnet string

type YandexHostSpecsRedis struct {
	// ZoneID: ID of the availability zone where the host resides.
	// To get a list of available zones, use the [yandex.cloud.compute.v1.ZoneService.List] request.
	// +kubebuilder:validation:Required
	ZoneID YandexZone `json:"zoneID"`

	// SubnetID: ID of the subnet that the host should belong to. This subnet should be a part
	// of the network that the cluster belongs to.
	// The ID of the network is set in the field [Cluster.network_id].
	// +kubebuilder:validation:Required
	SubnetID YandexSubnet `json:"subnetID"`

	// ShardName: ID of the Redis shard the host belongs to.
	// To get the shard ID use a [ClusterService.ListShards] request.
	// +kubebuilder:validation:Optional
	ShardName string `json:"shardName"`
}

// YandexEnvironmentRedis Custom type for deployment environment of the Redis cluster
// +kubebuilder:validation:Enum=production;prestable
type YandexEnvironmentRedis string

type RedisClusterStatus string

const (
	Creating RedisClusterStatus = "CREATING"
	Active   RedisClusterStatus = "ACTIVE"
	Deleting RedisClusterStatus = "DELETING"
)

// YandexManagedRedisSpec defines the desired state of yandexmanagedredis
type YandexManagedRedisSpec struct {
	// FolderID: id of a folder in which redis is located. Must be immutable.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:
	FolderID string `json:"folderId"`

	// Must be unique within the Cloud. Valid length is from 1 to 63 characters.
	// Should satisfy the regular expression [a-zA-Z0-9_-]+
	// +kubebuilder:validation:MaxLength=63
	// +kubebuilder:validation:Pattern=[a-z0-9][a-z0-9-.]*[a-z0-9]
	// +kubebuilder:validation:Required
	Name string `json:"name"`

	// Description: Maximum length of the description is 256 characters.
	// +kubebuilder:validation:MaxLength=256
	// +kubebuilder:validation:Optional
	Description string `json:"description"`

	// Labels: Custom labels for the Redis cluster as `key:value` pairs. Maximum 64 per cluster.
	// +kubebuilder:validation:Optional
	Labels map[string]string `json:"labels"`

	// Environment: Deployment environment of the Redis cluster.
	// +kubebuilder:validation:Required
	Environment YandexEnvironmentRedis `json:"environment"`

	// Version: Redis version
	// +kubebuilder:validation:Required
	Version string `json:"version"`

	// Sharded: Redis cluster mode on/off.
	// +kubebuilder:default=false
	// +kubebuilder:validation:Optional
	Sharded bool `json:"sharded"`

	// TlsEnabled: TLS port and functionality on\off
	// +kubebuilder:default=false
	// +kubebuilder:validation:Optional
	TLSEnabled bool `json:"tlsEnabled"`

	// HostClass: ID of the preset for computational resources available to a host (CPU, memory etc.).
	// All available presets are listed in the [documentation](/docs/managed-redis/concepts/instance-types).
	// +kubebuilder:validation:Required
	HostClass YandexHostClassRedis `json:"hostClass"`

	// DiskType: Type of the storage environment for the host.
	// Possible values:
	// * network-ssd - network SSD drive,
	// * local-ssd - local SSD storage.
	DiskType YandexDiskType `json:"diskType"`

	// DiskSize: Volume of the storage available to a host, in GB.
	DiskSize YandexDiskSize `json:"diskSize"`

	// HostSpecs: Individual configurations for hosts that should be created for the Redis cluster.
	HostSpecs []YandexHostSpecsRedis `json:"hosts"`
}

// YandexManagedRedisStatus defines the observed state of yandexmanagedredis
type YandexManagedRedisStatus struct {
	// ID: id of redis cluster
	ID string `json:"id,omitempty"`

	// Status: status of redis cluster.
	// Valid values are:
	// - CREATING
	// - ACTIVE
	// - DELETING
	Status RedisClusterStatus `json:"status,omitempty"`

	// CreatedAt: RFC3339-formatted string, representing creation time of resource
	CreatedAt string `json:"createdAt,omitempty"`

	// Labels: redis cluster labels in key:value form. Maximum of 64 labels for resource is allowed
	Labels map[string]string `json:"labels,omitempty"`
}

// YandexManagedRedis is the Schema for the yandexmanagedredis API
// +kubebuilder:object:root=true
// +kubebuilder:resource:path=yandexmanagedrediss,singular=yandexmanagedredis,shortName=ymr
// +kubebuilder:printcolumn:name="Name",type="string",JSONPath=".spec.name",description="Redis Name"
// +kubebuilder:printcolumn:name="HostClass",type="string",JSONPath=".spec.hostClass",description="Redis host class"
// +kubebuilder:printcolumn:name="Version",type="string",JSONPath=".spec.version",description="Redis version"
// +kubebuilder:printcolumn:name="DiskType",type="string",JSONPath=".spec.diskType",description="Redis disk type"
type YandexManagedRedis struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   YandexManagedRedisSpec   `json:"spec,omitempty"`
	Status YandexManagedRedisStatus `json:"status,omitempty"`
}

// YandexManagedRedisList contains a list of yandexmanagedredis
// +kubebuilder:object:root=true
type YandexManagedRedisList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []YandexManagedRedis `json:"items"`
}

func init() {
	SchemeBuilder.Register(&YandexManagedRedis{}, &YandexManagedRedisList{})
}
