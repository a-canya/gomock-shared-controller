package gomocksharedcontroller

//go:generate mockgen -destination=mock.go -package=gomocksharedcontroller . I

type I interface {
	M()
}

func CallM(i I) {
	i.M()
}
