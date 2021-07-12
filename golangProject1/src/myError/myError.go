package myErrorPack

type MyError struct{
	code int
	message string
}

func (MyError) Error() string {
	panic("implement me")
}

func (MyError) RuntimeError() {
	panic("implement me")
}

