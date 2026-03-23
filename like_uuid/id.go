package like_uuid

import (
	"fmt"
	"github.com/google/uuid"
)

// ---------------------------------------------------------------------------------------------------------------------
// Const
// ---------------------------------------------------------------------------------------------------------------------

var IdNil = Id{}

// ---------------------------------------------------------------------------------------------------------------------
// Struct
// ---------------------------------------------------------------------------------------------------------------------

type Id struct {
	value uuid.UUID
}

// ---------------------------------------------------------------------------------------------------------------------
// Create
// ---------------------------------------------------------------------------------------------------------------------

func IdFromInt128(value [16]byte) Id {
	// `uuid.FromBytes()` returns error only if `len != 16`,
	// but that is impossible in case of `[16]byte` input value,
	// i.e. any value is already correct here.
	return Id{value: uuid.Must(uuid.FromBytes(value[:]))}
}

// ---------------------------------------------------------------------------------------------------------------------

// IdFromString
//
// Accepts only lowercase 36 bytes length strings.
func IdFromString(value string) (Id, error) {
	_ = uuid.Parse // can parse more formats,
	// but this package and method is only about 36 bytes length strings.
	if len(value) != 36 {
		return IdNil, fmt.Errorf("%q is not a valid UUID string", value)
	}

	v, err := uuid.Parse(value)

	if err != nil {
		return IdNil, fmt.Errorf("%q is not a valid UUID string: %w", value, err)
	}

	_ = uuid.Parse // -> .String() returns a string of lowercased 36 bytes.
	// A difference between the original input and the String() method result in case of uppercased input may confuse,
	// so a strict validation rule must be defined -- accept only lowercased strings.
	if value != v.String() {
		return IdNil, fmt.Errorf("%q is not a valid UUID string", value)
	}

	return Id{value: v}, nil
}

// IdFromStringMust -- see IdFromString().
func IdFromStringMust(value string) Id {
	id, err := IdFromString(value)
	if err != nil {
		panic(err)
	}
	return id
}

// ---------------------------------------------------------------------------------------------------------------------
// State
// ---------------------------------------------------------------------------------------------------------------------

func (id Id) IsNil() bool {
	return id == IdNil
}

func (id Id) String() string {
	return id.value.String()
}

func (id Id) Int128() [16]byte {
	return id.value
}

// ---------------------------------------------------------------------------------------------------------------------
