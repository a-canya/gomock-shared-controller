package gomocksharedcontroller

import (
	"testing"

	gomock "go.uber.org/mock/gomock"
)

// this test passes even though it seems the individual tests should not
func TestSharedMock(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockI(ctrl)

	t.Run("expect 2 calls, call  once", func(t *testing.T) {
		mock.EXPECT().M().Times(2)
		CallM(mock)
	})

	t.Run("no expectations, call once", func(t *testing.T) {
		CallM(mock)
	})
}

// only first test fails, even though both tests should
func TestSeparateMockSharedCtrl(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("expect 2 calls, call once", func(t *testing.T) {
		mock := NewMockI(ctrl)

		mock.EXPECT().M().Times(2)
		CallM(mock)
	})

	t.Run("no expectations, call once", func(t *testing.T) {
		mock := NewMockI(ctrl)

		CallM(mock)
	})
}

// both tests fail as expected
func TestSeparateMockAndCtrl(t *testing.T) {
	t.Run("expect 2 calls, call once", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock := NewMockI(ctrl)

		mock.EXPECT().M().Times(2)
		CallM(mock)
	})

	t.Run("no expectations, call once", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock := NewMockI(ctrl)

		CallM(mock)
	})
}
