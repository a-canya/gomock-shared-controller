package gomocksharedcontroller

import (
	"testing"

	gomock "go.uber.org/mock/gomock"
)

//go:generate mockgen -destination=mock.go -package=gomocksharedcontroller . I
type I interface{ M() }

// Both subtests should fail but they pass.
func TestSharedMock(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockI(ctrl)

	t.Run("expect 2 calls, call  once", func(t *testing.T) {
		mock.EXPECT().M().Times(2)
		mock.M()
	})

	t.Run("no expectations, call once", func(t *testing.T) {
		mock.M()
	})
}

// Only the first test fails, even though both tests should fail. Error message
// is confusing.
//
// Test failure message:
//
//	`
//	--- FAIL: TestSeparateMockSharedCtrl (0.00s)
//	    a_test.go:52: Unexpected call to *gomocksharedcontroller.MockI.M([]) at /Users/arnau/test/gomock-shared-controller/a_test.go:52 because: there are no expected calls of the method "M" for that receiver
//	    --- FAIL: TestSeparateMockSharedCtrl/no_expectations,_call_once (0.00s)
//	        testing.go:1811: test executed panic(nil) or runtime.Goexit: subtest may have called FailNow on a parent test
//	    controller.go:97: missing call(s) to *gomocksharedcontroller.MockI.M() /Users/arnau/test/gomock-shared-controller/a_test.go:45
//	    controller.go:97: aborting test due to missing call(s)
//	`
func TestSeparateMockSharedCtrl(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("expect 2 calls, call once", func(t *testing.T) {
		mock := NewMockI(ctrl)

		mock.EXPECT().M().Times(2)
		mock.M()
	})

	t.Run("no expectations, call once", func(t *testing.T) {
		mock := NewMockI(ctrl)

		mock.M()
	})
}

// Both tests fail as expected. Error message is clear.
//
// Test failure message:
//
//	`
//	--- FAIL: TestSeparateMockAndCtrl (0.00s)
//	    --- FAIL: TestSeparateMockAndCtrl/expect_2_calls,_call_once (0.00s)
//	        controller.go:97: missing call(s) to *gomocksharedcontroller.MockI.M() /Users/arnau/test/gomock-shared-controller/a_test.go:73
//	        controller.go:97: aborting test due to missing call(s)
//	    --- FAIL: TestSeparateMockAndCtrl/no_expectations,_call_once (0.00s)
//	        a_test.go:81: Unexpected call to *gomocksharedcontroller.MockI.M([]) at /Users/arnau/test/gomock-shared-controller/a_test.go:81 because: there are no expected calls of the method "M" for that receiver
//	`
func TestSeparateMockAndCtrl(t *testing.T) {
	t.Run("expect 2 calls, call once", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock := NewMockI(ctrl)

		mock.EXPECT().M().Times(2)
		mock.M()
	})

	t.Run("no expectations, call once", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock := NewMockI(ctrl)

		mock.M()
	})
}
