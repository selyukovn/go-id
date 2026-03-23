package like_uuid

import (
	"fmt"
	"github.com/google/uuid"
)

func GenerateUniqueRandom() (Id, error) {
	v, err := uuid.NewRandom() // NewRandom -- v4
	if err != nil {
		return IdNil, fmt.Errorf("uuid generation error : %w", err)
	}

	id, err := IdFromString(v.String())
	if err != nil {
		return IdNil, fmt.Errorf("id creation error : %w", err)
	}

	return id, nil
}
