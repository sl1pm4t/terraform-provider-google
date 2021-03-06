package google

import (
	computeBeta "google.golang.org/api/compute/v0.beta"
	"google.golang.org/api/compute/v1"
)

func computeSharedOperationWaitZone(config *Config, op interface{}, project string, zone, activity string) error {
	return computeSharedOperationWaitZoneTime(config, op, project, zone, 4, activity)
}

func computeSharedOperationWaitZoneTime(config *Config, op interface{}, project string, zone string, minutes int, activity string) error {
	switch op.(type) {
	case *compute.Operation:
		return computeOperationWaitTime(config, op.(*compute.Operation), project, activity, minutes)
	case *computeBeta.Operation:
		return computeBetaOperationWaitZoneTime(config, op.(*computeBeta.Operation), project, zone, minutes, activity)
	case nil:
		panic("Attempted to wait on an Operation that was nil.")
	default:
		panic("Attempted to wait on an Operation of unknown type.")
	}
}
