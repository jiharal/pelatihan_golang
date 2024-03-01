package methods

func InterfaceValuesWithNil() {
	var i I1

	var t *T1
	i = t
	describe(i)
	i.M1()

	i = &T1{"hello"}
	describe(i)
	i.M1()
}
