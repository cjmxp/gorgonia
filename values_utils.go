package gorgonia

import (
	"fmt"

	"github.com/chewxy/gorgonia/tensor"
	"github.com/chewxy/hm"
	"github.com/pkg/errors"
)

type Dtyper interface {
	Dtype() tensor.Dtype
}

type Typer interface {
	Type() hm.Type
}

type ValueEqualer interface {
	ValueEq(Value) bool
}

type Cloner interface {
	Clone() interface{}
}

type CopierTo interface {
	CopyTo(dest interface{}) error
}

type CopierFrom interface {
	CopyFrom(src interface{}) error
}

// TypeOf returns the Type of the value
func TypeOf(v Value) hm.Type {
	switch t := v.(type) {
	case tensor.Tensor:
		dt, dim := tensorInfo(t)
		return newTensorType(dim, dt)
	case Scalar:
		return DtypeOf(t)
	case Typer:
		return t.Type()

	default:
		panic(fmt.Sprintf("TypeOf Not yet implemented for %v %T", v, v))
	}
}

// DtypeOf returns the Dtype of a Value
func DtypeOf(v Value) tensor.Dtype {
	switch vt := v.(type) {
	case F64:
		return tensor.Float64
	case F32:
		return tensor.Float32
	case I:
		return tensor.Int
	case I32:
		return tensor.Int32
	case I64:
		return tensor.Int64
	case U8:
		return tensor.Byte
	case B:
		return tensor.Bool
	case *tensor.Dense:
		return vt.Dtype()
	case Dtyper:
		return vt.Dtype()
	default:
		panic(fmt.Sprintf(nyiTypeFail, "DtypeOf", v))
	}
}

// ValueEq is the equality function for values
func ValueEq(a, b Value) bool {
	switch at := a.(type) {
	case Scalar:
		if bt, ok := b.(Scalar); ok {
			return scalarEq(at, bt)
		}
		return false
	case tensor.Tensor:
		if bt, ok := b.(tensor.Tensor); ok {
			return at.Eq(bt)
		}
		return false
	case ValueEqualer:
		return at.ValueEq(b)
	default:
		panic("Not implemented yet")
	}
}

// CloneValue clones a value. For scalars, since Go copies scalars, it returns itself
func CloneValue(v Value) (Value, error) {
	switch vt := v.(type) {
	case F64:
		return vt, nil
	case F32:
		return vt, nil
	case I:
		return vt, nil
	case I32:
		return vt, nil
	case I64:
		return vt, nil
	case U8:
		return vt, nil
	case B:
		return vt, nil
	case tensor.Tensor:
		return vt.Clone().(tensor.Tensor), nil
	case Cloner:
		return vt.Clone().(Value), nil
	default:
		return nil, errors.Errorf("Unable to clone value of type %T", v)
	}

}

// ZeroValue returns the zero value of a type
func ZeroValue(v Value) Value {
	switch vt := v.(type) {
	case F64:
		return F64(0)
	case F32:
		return F32(0)
	case I:
		return I(0)
	case I32:
		return I32(0)
	case I64:
		return I64(0)
	case U8:
		return U8(0)
	case B:
		return B(false)
	case tensor.Tensor:
		vt.Zero()
		return vt
	case ZeroValuer:
		return vt.ZeroValue()
	default:
		panic(fmt.Sprintf("Cannot return zero value of %T", v))
	}
}

// Copy copies the src values into dest values. For scalars, it just returns itself
func Copy(dest, src Value) (Value, error) {
	var ok bool
	switch srcT := src.(type) {
	case Scalar:
		return src, nil
	case tensor.Tensor:
		var destT tensor.Tensor
		if destT, ok = dest.(tensor.Tensor); !ok {
			return nil, errors.Errorf("Expected dest to be a tensor.Tensor. Got %T instead", dest)
		}
		err := tensor.Copy(destT, srcT)
		return dest, err
	case CopierTo:
		err := srcT.CopyTo(dest)
		return dest, err
	default:
		var copyFrom CopierFrom
		if copyFrom, ok = dest.(CopierFrom); ok {
			err := copyFrom.CopyFrom(src)
			return dest, err
		}
		return nil, errors.Errorf("Unable to copy value of type %T into value of type %T", src, dest)
	}
}
