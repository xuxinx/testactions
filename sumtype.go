package testactions

//go-sumtype:decl MySumType

type MySumType interface {
	sealed()
}

type VariantA struct{}

func (*VariantA) sealed() {}

type VariantB struct{}

func (*VariantB) sealed() {}

func init() {
	switch MySumType(nil).(type) {
	case *VariantA:
	}
}
