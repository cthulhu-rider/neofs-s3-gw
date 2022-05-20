package data

import (
	"time"

	oid "github.com/nspcc-dev/neofs-sdk-go/object/id"
	"github.com/nspcc-dev/neofs-sdk-go/user"
)

// NodeVersion represent node from tree service.
type NodeVersion struct {
	BaseNodeVersion
	DeleteMarker  *DeleteMarkerInfo
	IsUnversioned bool
}

// DeleteMarkerInfo is used to save object info if node in the tree service is delete marker.
// We need this information because the "delete marker" object is no longer stored in NeoFS.
type DeleteMarkerInfo struct {
	FilePath string
	Created  time.Time
	Owner    user.ID
}

// ExtendedObjectInfo contains additional node info to be able to sort versions by timestamp.
type ExtendedObjectInfo struct {
	ObjectInfo  *ObjectInfo
	NodeVersion *NodeVersion
}

// BaseNodeVersion is minimal node info from tree service.
// Basically used for "system" object.
type BaseNodeVersion struct {
	ID        uint64
	OID       oid.ID
	Timestamp uint64
}
