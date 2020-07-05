package maps

import "github.com/cheekybits/genny/generic"

type A generic.Type
type B generic.Type

func KeySliceᐸA_Bᐳ(m map[A]B) []A {
	out := make([]A, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	return out
}

func ValueSliceᐸA_Bᐳ(m map[A]B) []B {
	out := make([]B, 0, len(m))
	for _, v := range m {
		out = append(out, v)
	}
	return out
}
