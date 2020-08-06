package assert

func Equal(ok bool, panicMsg string) {
	if !ok {
		panic(panicMsg)
	}
}
