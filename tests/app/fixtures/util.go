package fixtures

func must[T any](result T, err error) T {
	if err != nil {
		panic(err)
	}
	return result 
}

func mustv(err error) {
	if err != nil {
		panic(err)
	}
}