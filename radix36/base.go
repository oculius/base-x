package radix36

import (
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/martinlindhe/base36"
	baseX "github.com/oculius/base-x"
)

// Int is int64
// UUID is uint128
// Bytes inherit from the lastType or can be anything.

type (
	radix36 struct {
		data     []byte
		lastType R36Type
	}

	R36Type int
)

const (
	none R36Type = iota + 1
	integer
	bytes
	t_uuid
)

func (r *radix36) Bytes() []byte {
	return r.data
}

func NewRadix36(r36 string) (baseX.BaseX, error) {
	if !Validate(r36) {
		return nil, fmt.Errorf("%s is not a radix36", r36)
	}
	return &radix36{
		data:     base36.DecodeToBytes(r36),
		lastType: bytes,
	}, nil
}

func Radix36(r36 string) baseX.BaseX {
	if !Validate(r36) {
		panic(r36 + " is not a radix36")
	}
	return &radix36{
		data:     base36.DecodeToBytes(r36),
		lastType: bytes,
	}
}

func New() baseX.BaseX {
	return &radix36{
		data:     nil,
		lastType: none,
	}
}

func NewFromInt(val int64) baseX.BaseX {
	x := New()
	x.FromInt(val)
	return x
}

func NewFromUUID(val uuid.UUID) baseX.BaseX {
	x := New()
	x.FromUUID(val)
	return x
}

func NewFromBytes(val []byte) baseX.BaseX {
	x := New()
	x.FromBytes(val)
	return x
}

func NewFromUUIDString(val string) (baseX.BaseX, error) {
	x := New()
	if err := x.FromUUIDString(val); err != nil {
		return x, err
	}
	return x, nil
}

func NewRandomize() baseX.BaseX {
	return New().Randomize().(baseX.BaseX)
}

func (r *radix36) String() string {
	if r.data == nil {
		return ""
	}
	if len(r.data) == 0 {
		return ""
	}

	zeroPrefix := 0
	for i := range r.data {
		if r.data[i] == byte(0) {
			zeroPrefix++
		} else {
			break
		}
	}
	if len(r.data) == zeroPrefix {
		return "0"
	}
	return base36.EncodeBytes(r.data[zeroPrefix:])
}
