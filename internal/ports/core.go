package ports

type ArithmeticPort interface {
	Add(a int32, b int32) (int32, error)
	Subtract(a int32, b int32) (int32, error)
	Multiply(a int32, b int32) (int32, error)
	Divide(a int32, b int32) (int32, error)
}
