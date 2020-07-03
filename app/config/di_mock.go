package config

type ContainerMock struct {
	SafeGetFn func(name string) (interface{}, error)
	GetFn     func(name string) interface{}
	FillFn    func(name string, dst interface{}) error
}

func (c ContainerMock) SafeGet(name string) (interface{}, error) {
	return c.SafeGetFn(name)
}
func (c ContainerMock) Get(name string) interface{} {
	return c.GetFn(name)
}
func (c ContainerMock) Fill(name string, dst interface{}) error {
	return c.Fill(name, dst)
}
