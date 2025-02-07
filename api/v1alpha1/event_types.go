package v1alpha1

const (
	// ReasonReplicationConfiguring indicates that replication is being configured.
	ReasonReplicationConfiguring = "ReplicationConfiguring"
	// ReasonReplicationConfigured indicates that replication has been configured.
	ReasonReplicationConfigured = "ReplicationConfigured"
	// ReasonReplicationPrimarySwitching indicates that primary is being switched.
	ReasonReplicationPrimarySwitching = "PrimarySwitching"
	// ReasonReplicationPrimarySwitched indicates that primary has been switched.
	ReasonReplicationPrimarySwitched = "PrimarySwitched"
	// ReasonReplicationPrimaryReadonly indicates that primary is being changed to readonly mode.
	ReasonReplicationPrimaryReadonly = "PrimaryReadonly"
	// ReasonReplicationReplicaSync indicates that replicas are being synced with primary.
	ReasonReplicationReplicaSync = "ReplicaSync"
	// ReasonReplicationReplicaSyncErr indicates that an error has happened while replicas were being synced with primary.
	ReasonReplicationReplicaSyncErr = "ReplicaSyncErr"
	// ReasonReplicationPrimaryNew indicates that a new primary is being configured.
	ReasonReplicationPrimaryNew = "PrimaryNew"
	// ReasonReplicationReplicaConn indicates that replicas are connecting to the new primary.
	ReasonReplicationReplicaConn = "ReplicaConn"
	// ReasonReplicationPrimaryToReplica indicates that current primary is becoming a replica.
	ReasonReplicationPrimaryToReplica = "PrimaryToReplica"
	// ReasonReplicationPrimarySvcUpdate indicates that primary service is being updated.
	ReasonReplicationPrimarySvcUpdate = "PrimarySvcUpdate"

	// ReasonGaleraClusterHealthy indicates that the cluster is healthy,
	ReasonGaleraClusterHealthy = "GaleraClusterHealthy"
	// ReasonGaleraClusterNotHealthy indicates that the cluster is not healthy.
	ReasonGaleraClusterNotHealthy = "GaleraClusterNotHealthy"
	// ReasonGaleraClusterBootstrap indicates that the cluster is being bootstrapped.
	ReasonGaleraClusterBootstrap = "GaleraClusterBootstrap"
	// ReasonGaleraClusterBootstrapTimeout indicates that the cluster bootstrap has timed out.
	ReasonGaleraClusterBootstrapTimeout = "GaleraClusterBootstrapTimeout"
	// ReasonGaleraPodStateFetched indicates that the Pod state has been fetched successfully.
	ReasonGaleraPodStateFetched = "GaleraPodStateFetched"
	// ReasonGaleraPodRecovered indicates that the Pod has successfully recovered the sequence.
	ReasonGaleraPodRecovered = "GaleraPodRecovered"
	// ReasonGaleraPodSyncTimeout indicates that the Pod has timed out reaching the Sync state.
	ReasonGaleraPodSyncTimeout = "GaleraPodSyncTimeout"
)
