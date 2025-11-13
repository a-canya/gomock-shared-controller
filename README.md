# GoMock Shared Controller

This repos shows the issue with sharing gomock.Controller and gomock-generated
mocks between different subtests (`t.Run`).

Take a look at `a.go` and then `a_test.go`. All tests defined should fail but
some don't. If we run `go test .` this is the output:

```txt
--- FAIL: TestSeparateMockSharedCtrl (0.00s)
    a.go:10: Unexpected call to *gomocksharedcontroller.MockI.M([]) at /Users/arnau/test/gomock-shared-controller/a.go:10 because: there are no expected calls of the method "M" for that receiver
    --- FAIL: TestSeparateMockSharedCtrl/no_expectations,_call_once (0.00s)
        testing.go:1811: test executed panic(nil) or runtime.Goexit: subtest may have called FailNow on a parent test
    controller.go:251: missing call(s) to *gomocksharedcontroller.MockI.M() /Users/arnau/test/gomock-shared-controller/a_test.go:33
    controller.go:251: aborting test due to missing call(s)
--- FAIL: TestSeparateMockAndCtrl (0.00s)
    --- FAIL: TestSeparateMockAndCtrl/expect_2_calls,_call_once (0.00s)
        controller.go:251: missing call(s) to *gomocksharedcontroller.MockI.M() /Users/arnau/test/gomock-shared-controller/a_test.go:51
        controller.go:251: aborting test due to missing call(s)
    --- FAIL: TestSeparateMockAndCtrl/no_expectations,_call_once (0.00s)
        a.go:10: Unexpected call to *gomocksharedcontroller.MockI.M([]) at /Users/arnau/test/gomock-shared-controller/a.go:10 because: there are no expected calls of the method "M" for that receiver
FAIL
exit status 1
FAIL	gomock-shared-controller	0.201s
```
