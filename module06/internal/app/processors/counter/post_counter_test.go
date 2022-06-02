package counter

import (
	"module06/test/gomock/mocks/postmock"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestPostCount(t *testing.T) {
	req := require.New(t)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	client01 := postmock.NewMockClient(mockCtrl)
	client01.EXPECT().GetList().AnyTimes()
	_, err := PostCount(client01)
	req.NoError(err)

}
