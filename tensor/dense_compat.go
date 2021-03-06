package tensor

import (
	"fmt"
	"math"
	"math/cmplx"
	"reflect"

	"github.com/chewxy/math32"
	"github.com/gonum/matrix/mat64"
	"github.com/pkg/errors"
)

/*
GENERATED FILE. DO NOT EDIT
*/

func convFromFloat64s(to Dtype, data []float64) interface{} {
	switch to.Kind() {
	case reflect.Int:
		retVal := make([]int, len(data))
		for i, v := range data {
			switch {
			case math.IsNaN(v), math.IsInf(v, 0):
				retVal[i] = 0
			default:
				retVal[i] = int(v)
			}
		}
		return retVal
	case reflect.Int8:
		retVal := make([]int8, len(data))
		for i, v := range data {
			switch {
			case math.IsNaN(v), math.IsInf(v, 0):
				retVal[i] = 0
			default:
				retVal[i] = int8(v)
			}
		}
		return retVal
	case reflect.Int16:
		retVal := make([]int16, len(data))
		for i, v := range data {
			switch {
			case math.IsNaN(v), math.IsInf(v, 0):
				retVal[i] = 0
			default:
				retVal[i] = int16(v)
			}
		}
		return retVal
	case reflect.Int32:
		retVal := make([]int32, len(data))
		for i, v := range data {
			switch {
			case math.IsNaN(v), math.IsInf(v, 0):
				retVal[i] = 0
			default:
				retVal[i] = int32(v)
			}
		}
		return retVal
	case reflect.Int64:
		retVal := make([]int64, len(data))
		for i, v := range data {
			switch {
			case math.IsNaN(v), math.IsInf(v, 0):
				retVal[i] = 0
			default:
				retVal[i] = int64(v)
			}
		}
		return retVal
	case reflect.Uint:
		retVal := make([]uint, len(data))
		for i, v := range data {
			switch {
			case math.IsNaN(v), math.IsInf(v, 0):
				retVal[i] = 0
			default:
				retVal[i] = uint(v)
			}
		}
		return retVal
	case reflect.Uint8:
		retVal := make([]uint8, len(data))
		for i, v := range data {
			switch {
			case math.IsNaN(v), math.IsInf(v, 0):
				retVal[i] = 0
			default:
				retVal[i] = uint8(v)
			}
		}
		return retVal
	case reflect.Uint16:
		retVal := make([]uint16, len(data))
		for i, v := range data {
			switch {
			case math.IsNaN(v), math.IsInf(v, 0):
				retVal[i] = 0
			default:
				retVal[i] = uint16(v)
			}
		}
		return retVal
	case reflect.Uint32:
		retVal := make([]uint32, len(data))
		for i, v := range data {
			switch {
			case math.IsNaN(v), math.IsInf(v, 0):
				retVal[i] = 0
			default:
				retVal[i] = uint32(v)
			}
		}
		return retVal
	case reflect.Uint64:
		retVal := make([]uint64, len(data))
		for i, v := range data {
			switch {
			case math.IsNaN(v), math.IsInf(v, 0):
				retVal[i] = 0
			default:
				retVal[i] = uint64(v)
			}
		}
		return retVal
	case reflect.Float32:
		retVal := make([]float32, len(data))
		for i, v := range data {
			switch {
			case math.IsNaN(v):
				retVal[i] = math32.NaN()
			case math.IsInf(v, 1):
				retVal[i] = math32.Inf(1)
			case math.IsInf(v, -1):
				retVal[i] = math32.Inf(-1)
			default:
				retVal[i] = float32(v)
			}
		}
		return retVal
	case reflect.Float64:
		retVal := make([]float64, len(data))
		copy(retVal, data)
		return retVal
	case reflect.Complex64:
		retVal := make([]complex64, len(data))
		for i, v := range data {
			switch {
			case math.IsNaN(v):
				retVal[i] = complex64(cmplx.NaN())
			case math.IsInf(v, 0):
				retVal[i] = complex64(cmplx.Inf())
			default:
				retVal[i] = complex(float32(v), float32(0))
			}
		}
		return retVal
	case reflect.Complex128:
		retVal := make([]complex128, len(data))
		for i, v := range data {
			switch {
			case math.IsNaN(v):
				retVal[i] = cmplx.NaN()
			case math.IsInf(v, 0):
				retVal[i] = cmplx.Inf()
			default:
				retVal[i] = complex(v, float64(0))
			}
		}
		return retVal
	default:
		panic("Unsupported Dtype")
	}
	panic("Unreachable")
}

func convToFloat64s(t *Dense) (retVal []float64) {
	retVal = make([]float64, t.len())
	switch t.t.Kind() {
	case reflect.Int:
		for i, v := range t.ints() {
			retVal[i] = float64(v)
		}
		return retVal
	case reflect.Int8:
		for i, v := range t.int8s() {
			retVal[i] = float64(v)
		}
		return retVal
	case reflect.Int16:
		for i, v := range t.int16s() {
			retVal[i] = float64(v)
		}
		return retVal
	case reflect.Int32:
		for i, v := range t.int32s() {
			retVal[i] = float64(v)
		}
		return retVal
	case reflect.Int64:
		for i, v := range t.int64s() {
			retVal[i] = float64(v)
		}
		return retVal
	case reflect.Uint:
		for i, v := range t.uints() {
			retVal[i] = float64(v)
		}
		return retVal
	case reflect.Uint8:
		for i, v := range t.uint8s() {
			retVal[i] = float64(v)
		}
		return retVal
	case reflect.Uint16:
		for i, v := range t.uint16s() {
			retVal[i] = float64(v)
		}
		return retVal
	case reflect.Uint32:
		for i, v := range t.uint32s() {
			retVal[i] = float64(v)
		}
		return retVal
	case reflect.Uint64:
		for i, v := range t.uint64s() {
			retVal[i] = float64(v)
		}
		return retVal
	case reflect.Float32:
		for i, v := range t.float32s() {
			switch {
			case math32.IsNaN(v):
				retVal[i] = math.NaN()
			case math32.IsInf(v, 1):
				retVal[i] = math.Inf(1)
			case math32.IsInf(v, -1):
				retVal[i] = math.Inf(-1)
			default:
				retVal[i] = float64(v)
			}
		}
		return retVal
	case reflect.Float64:
		return t.float64s()
		return retVal
	case reflect.Complex64:
		for i, v := range t.complex64s() {
			switch {
			case cmplx.IsNaN(complex128(v)):
				retVal[i] = math.NaN()
			case cmplx.IsInf(complex128(v)):
				retVal[i] = math.Inf(1)
			default:
				retVal[i] = float64(real(v))
			}
		}
		return retVal
	case reflect.Complex128:
		for i, v := range t.complex128s() {
			switch {
			case cmplx.IsNaN(v):
				retVal[i] = math.NaN()
			case cmplx.IsInf(v):
				retVal[i] = math.Inf(1)
			default:
				retVal[i] = real(v)
			}
		}
		return retVal
	default:
		panic(fmt.Sprintf("Cannot convert *Dense of %v to []float64", t.t))
	}
	panic("Unreachable")
}

func convToFloat64(x interface{}) float64 {
	switch xt := x.(type) {
	case int:
		return float64(xt)
	case int8:
		return float64(xt)
	case int16:
		return float64(xt)
	case int32:
		return float64(xt)
	case int64:
		return float64(xt)
	case uint:
		return float64(xt)
	case uint8:
		return float64(xt)
	case uint16:
		return float64(xt)
	case uint32:
		return float64(xt)
	case uint64:
		return float64(xt)
	case float32:
		return float64(xt)
	case float64:
		return float64(xt)
	case complex64:
		return float64(real(xt))
	case complex128:
		return real(xt)
	default:
		panic("Cannot convert to float64")
	}
	panic("Unreachable")
}

// FromMat64 converts a *"gonum/matrix/mat64".Dense into a *tensorf64.Tensor.
func FromMat64(m *mat64.Dense, opts ...FuncOpt) *Dense {
	r, c := m.Dims()
	fo := parseFuncOpts(opts...)
	toCopy := fo.safe()
	as := fo.t
	if fo.t.Type == nil {
		as = Float64
	}

	switch as.Kind() {
	case reflect.Int:
		backing := convFromFloat64s(Int, m.RawMatrix().Data).([]int)
		retVal := New(WithBacking(backing), WithShape(r, c))
		return retVal
	case reflect.Int8:
		backing := convFromFloat64s(Int8, m.RawMatrix().Data).([]int8)
		retVal := New(WithBacking(backing), WithShape(r, c))
		return retVal
	case reflect.Int16:
		backing := convFromFloat64s(Int16, m.RawMatrix().Data).([]int16)
		retVal := New(WithBacking(backing), WithShape(r, c))
		return retVal
	case reflect.Int32:
		backing := convFromFloat64s(Int32, m.RawMatrix().Data).([]int32)
		retVal := New(WithBacking(backing), WithShape(r, c))
		return retVal
	case reflect.Int64:
		backing := convFromFloat64s(Int64, m.RawMatrix().Data).([]int64)
		retVal := New(WithBacking(backing), WithShape(r, c))
		return retVal
	case reflect.Uint:
		backing := convFromFloat64s(Uint, m.RawMatrix().Data).([]uint)
		retVal := New(WithBacking(backing), WithShape(r, c))
		return retVal
	case reflect.Uint8:
		backing := convFromFloat64s(Uint8, m.RawMatrix().Data).([]uint8)
		retVal := New(WithBacking(backing), WithShape(r, c))
		return retVal
	case reflect.Uint16:
		backing := convFromFloat64s(Uint16, m.RawMatrix().Data).([]uint16)
		retVal := New(WithBacking(backing), WithShape(r, c))
		return retVal
	case reflect.Uint32:
		backing := convFromFloat64s(Uint32, m.RawMatrix().Data).([]uint32)
		retVal := New(WithBacking(backing), WithShape(r, c))
		return retVal
	case reflect.Uint64:
		backing := convFromFloat64s(Uint64, m.RawMatrix().Data).([]uint64)
		retVal := New(WithBacking(backing), WithShape(r, c))
		return retVal
	case reflect.Float32:
		backing := convFromFloat64s(Float32, m.RawMatrix().Data).([]float32)
		retVal := New(WithBacking(backing), WithShape(r, c))
		return retVal
	case reflect.Float64:
		var backing []float64
		if toCopy {
			backing = make([]float64, len(m.RawMatrix().Data))
			copy(backing, m.RawMatrix().Data)
		} else {
			backing = m.RawMatrix().Data
		}
		retVal := New(WithBacking(backing), WithShape(r, c))
		return retVal
	case reflect.Complex64:
		backing := convFromFloat64s(Complex64, m.RawMatrix().Data).([]complex64)
		retVal := New(WithBacking(backing), WithShape(r, c))
		return retVal
	case reflect.Complex128:
		backing := convFromFloat64s(Complex128, m.RawMatrix().Data).([]complex128)
		retVal := New(WithBacking(backing), WithShape(r, c))
		return retVal
	default:
		panic(fmt.Sprintf("Unsupported Dtype - cannot convert float64 to %v", as))
	}
	panic("Unreachable")
}

// ToMat64 converts a *Dense to a *mat64.Dense. All the values are converted into float64s.
// This function will only convert matrices. Anything *Dense with dimensions larger than 2 will cause an error.
func ToMat64(t *Dense, opts ...FuncOpt) (retVal *mat64.Dense, err error) {
	// checks:
	if !t.IsMatrix() {
		// error
		err = errors.Errorf("Cannot convert *Dense to *mat64.Dense. Expected number of dimensions: <=2, T has got %d dimensions (Shape: %v)", t.Dims(), t.Shape())
		return
	}

	fo := parseFuncOpts(opts...)
	toCopy := fo.safe()

	// fix dims
	r := t.Shape()[0]
	c := t.Shape()[1]

	var data []float64
	switch {
	case t.t.Kind() == reflect.Float64 && toCopy && !t.IsMaterializable():
		data = make([]float64, t.len())
		copy(data, t.float64s())
	case !t.IsMaterializable():
		data = convToFloat64s(t)
	default:
		it := NewFlatIterator(t.AP)
		var next int
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			if err = handleNoOp(err); err != nil {
				return
			}
			data = append(data, convToFloat64(t.Get(next)))
		}
		err = nil

	}

	retVal = mat64.NewDense(r, c, data)
	return
}
