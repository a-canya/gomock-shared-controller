package gomocksharedcontroller

//go:generate mockgen -destination=mock_test.go -package=gomocksharedcontroller . I

type I interface {
	M()
}

func CallM(i I) {
	i.M()
}
