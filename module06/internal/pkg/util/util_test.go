package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReversInt(t *testing.T) {
	req := require.New(t)
	res0, err := ReverseInt(123)
	req.NoError(err)
	req.Equal(321, res0)

	res1, err := ReverseInt(-649)
	req.NoError(err)
	req.Equal(-946, res1)

	res2, err := ReverseInt(0)
	req.NoError(err)
	req.Equal(0, res2) //	require.Equal(t, 0, res2)
}
