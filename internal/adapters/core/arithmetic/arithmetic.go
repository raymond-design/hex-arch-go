package arithmetic

type Adapter struct {
}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (adapter Adapter) Add(a int32, b int32) (int32, error) {
	return a + b, nil
}

func (adapter Adapter) Subtract(a int32, b int32) (int32, error) {
	return a - b, nil
}

func (adapter Adapter) Multiply(a int32, b int32) (int32, error) {
	return a * b, nil
}

func (adapter Adapter) Divide(a int32, b int32) (int32, error) {
	return a / b, nil
}
