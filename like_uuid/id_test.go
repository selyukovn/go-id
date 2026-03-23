package like_uuid

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func testDataProvider_Id_correctValuesStrings() []string {
	return []string{
		"aa902246-1f28-44c7-8452-04805999ac32",
		"592beeaa-e3d2-4768-bab3-7a5469d82bb3",
		"1f688344-cce7-4a60-a96e-768291c6d3e0",

		// Zero/nil value is also correct, but not for this data-provider.
		// "00000000-0000-0000-0000-000000000000",

		"aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
	}
}

func testDataProvider_Id_incorrectValuesStrings() []string {
	return []string{
		"",
		" ",
		"1",
		"123456789012345678901234567890123456",
		"a",
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa!!!",
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa ",
		" aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		" aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa ",
		"aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaa!!!",
		"ёёёёёёёё-ёёёё-ёёёё-ёёёё-ёёёёёёёёёёёё",

		"aaaaaaaa-aaaa-aaaa-aaaa-AAAAAAAAAAAA",
		"AAAAAAAA-AAAA-AAAA-AAAA-AAAAAAAAAAAA",
	}
}

// ---------------------------------------------------------------------------------------------------------------------
// IdFromInt128
// ---------------------------------------------------------------------------------------------------------------------

func Test_IdFromInt128(t *testing.T) {
	tCases := map[[16]byte]string{
		{0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa}: "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
		{0x1f, 0x68, 0x83, 0x44, 0xcc, 0xe7, 0x4a, 0x60, 0xa9, 0x6e, 0x76, 0x82, 0x91, 0xc6, 0xd3, 0xe0}: "1f688344-cce7-4a60-a96e-768291c6d3e0",
	}
	for tCase := range tCases {
		id := IdFromInt128(tCase)
		assert.False(t, id.IsNil())
		assert.Equal(t, tCase, id.Int128())
		assert.Equal(t, tCases[tCase], id.String())
	}

	// zero/nil
	nId := IdFromInt128([16]byte{ /* all zeroes */ })
	assert.Equal(t, [16]byte{ /* all zeroes */ }, nId.Int128())
	assert.Equal(t, "00000000-0000-0000-0000-000000000000", nId.String())
	assert.True(t, nId.IsNil())
}

// ---------------------------------------------------------------------------------------------------------------------
// IdFromString
// ---------------------------------------------------------------------------------------------------------------------

func Test_IdFromString(t *testing.T) {
	t.Run("correct", func(t *testing.T) {
		tCases := testDataProvider_Id_correctValuesStrings()
		for _, tCase := range tCases {
			id, err := IdFromString(tCase)
			// no error
			assert.NoError(t, err)
			// id not nil
			assert.NotEqual(t, IdNil, id)
			assert.False(t, id.IsNil())
			assert.Equal(t, tCase, id.String())
		}

		// zero/nil
		nId, err := IdFromString("00000000-0000-0000-0000-000000000000")
		assert.Equal(t, nId, IdNil)
		assert.True(t, nId.IsNil())
		assert.NoError(t, err)
	})

	t.Run("incorrect", func(t *testing.T) {
		tCases := testDataProvider_Id_incorrectValuesStrings()
		for _, tCase := range tCases {
			id, err := IdFromString(tCase)
			// error
			assert.Error(t, err)
			// id nil
			assert.Equal(t, IdNil, id)
			assert.True(t, id.IsNil())
		}
	})
}

// ---------------------------------------------------------------------------------------------------------------------
// IdFromStringMust
// ---------------------------------------------------------------------------------------------------------------------

func Test_IdFromStringMust(t *testing.T) {
	t.Run("correct", func(t *testing.T) {
		tCases := testDataProvider_Id_correctValuesStrings()
		for _, tCase := range tCases {
			// no panic
			assert.NotPanics(t, func() { IdFromStringMust(tCase) })
			// id not nil
			id := IdFromStringMust(tCase)
			assert.NotEqual(t, IdNil, id)
			assert.False(t, id.IsNil())
			assert.Equal(t, tCase, id.String())
		}
	})

	t.Run("incorrect", func(t *testing.T) {
		tCases := testDataProvider_Id_incorrectValuesStrings()
		for _, tCase := range tCases {
			assert.Panics(t, func() { IdFromStringMust(tCase) })
		}
	})
}

// ---------------------------------------------------------------------------------------------------------------------
// IsNil
// ---------------------------------------------------------------------------------------------------------------------

func Test_Id_IsNil(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		assert.True(t, IdNil.IsNil())
	})

	t.Run("false", func(t *testing.T) {
		tCases := testDataProvider_Id_correctValuesStrings()
		for _, tCase := range tCases {
			id, _ := IdFromString(tCase)
			assert.False(t, id.IsNil())
		}
	})
}

// ---------------------------------------------------------------------------------------------------------------------
// String
// ---------------------------------------------------------------------------------------------------------------------

func Test_Id_String(t *testing.T) {
	tCases := testDataProvider_Id_correctValuesStrings()
	for _, tCase := range tCases {
		id, _ := IdFromString(tCase)
		assert.Equal(t, tCase, id.String())
	}
}

// ---------------------------------------------------------------------------------------------------------------------
