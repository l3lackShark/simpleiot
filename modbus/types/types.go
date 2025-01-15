package modbus

type RegisterType string

const (
	RegisterTypeCoil    RegisterType = "coil"
	RegisterTypeInteger RegisterType = "integer"
)

type ChangedRegisters struct {
	Registers []ChangedRegister
}

type ChangedRegister struct {
	Type    RegisterType
	Address uint16
	Value   interface{}
}
