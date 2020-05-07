package records

type IRecord interface {
	ToBytes() []byte
}

