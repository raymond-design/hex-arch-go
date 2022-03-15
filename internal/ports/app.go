package ports

type API interface {
	GetAddition(a int32, b int32) (int32, error)
	GetSubtration(a int32, b int32) (int32, error)
	GetMultiplication(a int32, b int32) (int32, error)
	GetDivision(a int32, b int32) (int32, error)
}
