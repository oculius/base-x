package baseX

import (
	"github.com/gofrs/uuid"
)

type (
	BaseX interface {
		IntegerTransformable
		BytesTransformable
		UUIDTransformable
		String() string
		Bytes() []byte
	}

	IntegerTransformable interface {
		FromInt(value int64)
		ToInt() int64
	}

	BytesTransformable interface {
		FromBytes(value []byte)
		ToBytes() []byte
	}

	UUIDTransformable interface {
		FromUUID(value uuid.UUID)
		FromUUIDString(value string) error
		ToUUID() uuid.UUID
		Randomize() UUIDTransformable
	}
)
