package inversemap

import "reflect"

func InverseMap(fwd interface{}) (rev interface{}) {
	fwdType := reflect.TypeOf(fwd)
	revType := reflect.MapOf(fwdType.Elem(), fwdType.Key())
	revValue := reflect.MakeMap(revType)

	fwdValue := reflect.ValueOf(fwd)
	for _, k := range fwdValue.MapKeys() {
		v := fwdValue.MapIndex(k)
		revValue.SetMapIndex(v, k)
	}
	rev = revValue.Interface()
	return
}
