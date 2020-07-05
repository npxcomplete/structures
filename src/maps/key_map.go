package maps

import "github.com/cheekybits/genny/generic"

type K1 generic.Type
type K2 generic.Type
type V generic.Type

func KeyMapᐸK1_K2_Vᐳ(f func(x K1) K2) func(map[K1]V) map[K2]V {
	return func(m map[K1]V) map[K2]V {
		out := make(map[K2]V)
		for k, v := range m {
			out[f(k)] = v
		}
		return out
	}
}
