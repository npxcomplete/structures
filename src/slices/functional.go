package slices

import "github.com/cheekybits/genny/generic"

type OutType generic.Type
type InType generic.Type

func FMapᐸInType_OutTypeᐳ(f func(s InType) OutType) func([]InType) []OutType {
	return func(xs []InType) []OutType {
		ys := make([]OutType, len(xs))
		for i := range xs {
			ys[i] = f(xs[i])
		}
		return ys
	}
}

func BindᐸInType_OutTypeᐳ(f func(s InType) []OutType) func([]InType) []OutType {
	return func(xs []InType) []OutType {
		ys := make([]OutType, 0, 8)
		for i := range xs {
			ys = append(ys, f(xs[i])...)
		}
		return ys
	}
}

func FoldRightᐸInType_OutTypeᐳ(zero OutType, f func(lhs InType, rhs func() OutType) OutType) func([]InType) OutType {
	var g func(xs []InType) OutType
	g = func(xs []InType) OutType {
		if len(xs) == 0 {
			return zero
		}
		return f(xs[0], func() OutType { return g(xs[1:]) })
	}
	return g
}

func FoldLeftᐸInType_OutTypeᐳ(zero OutType, f func(lhs OutType, rhs InType) OutType) func([]InType) OutType {
	return func(xs []InType) OutType {
		var sum OutType = zero
		for i := 0; i < len(xs); i++ {
			sum = f(sum, xs[i])
		}
		return sum
	}
}
