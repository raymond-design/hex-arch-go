package api

import "github.com/raymond-design/hex-arch-go/internal/ports"

type Adapter struct {
	arith ports.ArithmeticPort
}

func NewAdapter(arith ports.ArithmeticPort) *Adapter {
	return &Adapter{arith: arith}
}

func (adapter Adapter) GetAddition(a int32, b int32) (int32, error) {
	answer, err := adapter.arith.Add(a, b)
	if err != nil {
		return 0, err
	}

	return answer, nil
}

func (adapter Adapter) GetSubtraction(a int32, b int32) (int32, error) {
	answer, err := adapter.arith.Subtract(a, b)
	if err != nil {
		return 0, err
	}

	return answer, nil
}

func (adapter Adapter) GetMultiplication(a int32, b int32) (int32, error) {
	answer, err := adapter.arith.Multiply(a, b)
	if err != nil {
		return 0, err
	}

	return answer, nil
}

func (adapter Adapter) GetDivision(a int32, b int32) (int32, error) {
	answer, err := adapter.arith.Divide(a, b)
	if err != nil {
		return 0, err
	}

	return answer, nil
}
