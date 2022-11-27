package goutil_loc

import (
	"github.com/21888/go-tagexpr-new/ameda-loc"
)

// CopyStrings creates a copy of the string slice.
func CopyStrings(a []string) []string {
	return ameda_loc.StringsCopy(a)
}

// StringsToBools converts string slice to bool slice.
func StringsToBools(a []string, emptyAsZero ...bool) ([]bool, error) {
	return ameda_loc.StringsToBools(a, emptyAsZero...)
}

// StringsToFloat32s converts string slice to float32 slice.
func StringsToFloat32s(a []string, emptyAsZero ...bool) ([]float32, error) {
	return ameda_loc.StringsToFloat32s(a, emptyAsZero...)
}

// StringsToFloat64s converts string slice to float64 slice.
func StringsToFloat64s(a []string, emptyAsZero ...bool) ([]float64, error) {
	return ameda_loc.StringsToFloat64s(a, emptyAsZero...)
}

// StringsToInts converts string slice to int slice.
func StringsToInts(a []string, emptyAsZero ...bool) ([]int, error) {
	return ameda_loc.StringsToInts(a, emptyAsZero...)
}

// StringsToInt64s converts string slice to int64 slice.
func StringsToInt64s(a []string, emptyAsZero ...bool) ([]int64, error) {
	return ameda_loc.StringsToInt64s(a, emptyAsZero...)
}

// StringsToInt32s converts string slice to int32 slice.
func StringsToInt32s(a []string, emptyAsZero ...bool) ([]int32, error) {
	return ameda_loc.StringsToInt32s(a, emptyAsZero...)
}

// StringsToInt16s converts string slice to int16 slice.
func StringsToInt16s(a []string, emptyAsZero ...bool) ([]int16, error) {
	return ameda_loc.StringsToInt16s(a, emptyAsZero...)
}

// StringsToInt8s converts string slice to int8 slice.
func StringsToInt8s(a []string, emptyAsZero ...bool) ([]int8, error) {
	return ameda_loc.StringsToInt8s(a, emptyAsZero...)
}

// StringsToUint8s converts string slice to uint8 slice.
func StringsToUint8s(a []string, emptyAsZero ...bool) ([]uint8, error) {
	return ameda_loc.StringsToUint8s(a, emptyAsZero...)
}

// StringsToUint16s converts string slice to uint16 slice.
func StringsToUint16s(a []string, emptyAsZero ...bool) ([]uint16, error) {
	return ameda_loc.StringsToUint16s(a, emptyAsZero...)
}

// StringsToUint32s converts string slice to uint32 slice.
func StringsToUint32s(a []string, emptyAsZero ...bool) ([]uint32, error) {
	return ameda_loc.StringsToUint32s(a, emptyAsZero...)
}

// StringsToUint64s converts string slice to uint64 slice.
func StringsToUint64s(a []string, emptyAsZero ...bool) ([]uint64, error) {
	return ameda_loc.StringsToUint64s(a, emptyAsZero...)
}

// StringsToUints converts string slice to uint slice.
func StringsToUints(a []string, emptyAsZero ...bool) ([]uint, error) {
	return ameda_loc.StringsToUints(a, emptyAsZero...)
}

// StringsConvert converts the string slice to a new slice using fn.
// If fn returns error, exit the conversion and return the error.
func StringsConvert(a []string, fn func(string) (string, error)) ([]string, error) {
	ret := make([]string, len(a))
	for i, s := range a {
		r, err := fn(s)
		if err != nil {
			return nil, err
		}
		ret[i] = r
	}
	return ret, nil
}

// StringsConvertMap converts the string slice to a new map using fn.
// If fn returns error, exit the conversion and return the error.
func StringsConvertMap(a []string, fn func(string) (string, error)) (map[string]string, error) {
	ret := make(map[string]string, len(a))
	for _, s := range a {
		r, err := fn(s)
		if err != nil {
			return nil, err
		}
		ret[s] = r
	}
	return ret, nil
}

// IntersectStrings calculate intersection of two sets.
func IntersectStrings(set1, set2 []string) []string {
	return ameda_loc.StringSetIntersect(set1, set2)
}

// StringsDistinct creates a string set that
// removes the same elements and returns them in their original order.
func StringsDistinct(a []string) (set []string) {
	set = ameda_loc.StringsCopy(a)
	ameda_loc.StringsDistinct(&set, true)
	return set
}

// SetToStrings sets a element to the string set.
func SetToStrings(set []string, a ...string) []string {
	return ameda_loc.StringsPushDistinct(set, a...)
}

// RemoveFromStrings removes the first element from the string set.
func RemoveFromStrings(set []string, a string) []string {
	ameda_loc.StringsRemoveFirst(&set, a)
	return set
}

// RemoveAllFromStrings removes all the a element from the string set.
func RemoveAllFromStrings(set []string, a string) []string {
	ameda_loc.StringsRemoveEvery(&set, a)
	return set
}

// IntsDistinct creates a int set that
// removes the same elements and returns them in their original order.
func IntsDistinct(a []int) (set []int) {
	set = ameda_loc.IntsCopy(a)
	ameda_loc.IntsDistinct(&set, true)
	return set
}

// SetToInts sets a element to the int set.
func SetToInts(set []int, a int) []int {
	return ameda_loc.IntsPushDistinct(set, a)
}

// RemoveFromInts removes the first element from the int set.
func RemoveFromInts(set []int, a int) []int {
	ameda_loc.IntsRemoveFirst(&set, a)
	return set
}

// RemoveAllFromInts removes all the a element from the int set.
func RemoveAllFromInts(set []int, a int) []int {
	ameda_loc.IntsRemoveEvery(&set, a)
	return set
}

// Int32sDistinct creates a int32 set that
// removes the same element32s and returns them in their original order.
func Int32sDistinct(a []int32) (set []int32) {
	set = ameda_loc.Int32sCopy(a)
	ameda_loc.Int32sDistinct(&set, true)
	return set
}

// SetToInt32s sets a element to the int32 set.
func SetToInt32s(set []int32, a int32) []int32 {
	return ameda_loc.Int32sPushDistinct(set, a)
}

// RemoveFromInt32s removes the first element from the int32 set.
func RemoveFromInt32s(set []int32, a int32) []int32 {
	ameda_loc.Int32sRemoveFirst(&set, a)
	return set
}

// RemoveAllFromInt32s removes all the a element from the int32 set.
func RemoveAllFromInt32s(set []int32, a int32) []int32 {
	ameda_loc.Int32sRemoveEvery(&set, a)
	return set
}

// Int64sDistinct creates a int64 set that
// removes the same element64s and returns them in their original order.
func Int64sDistinct(a []int64) (set []int64) {
	set = ameda_loc.Int64sCopy(a)
	ameda_loc.Int64sDistinct(&set, true)
	return set
}

// SetToInt64s sets a element to the int64 set.
func SetToInt64s(set []int64, a int64) []int64 {
	return ameda_loc.Int64sPushDistinct(set, a)
}

// RemoveFromInt64s removes the first element from the int64 set.
func RemoveFromInt64s(set []int64, a int64) []int64 {
	ameda_loc.Int64sRemoveFirst(&set, a)
	return set
}

// RemoveAllFromInt64s removes all the a element from the int64 set.
func RemoveAllFromInt64s(set []int64, a int64) []int64 {
	ameda_loc.Int64sRemoveEvery(&set, a)
	return set
}

// InterfacesDistinct creates a interface{} set that
// removes the same elements and returns them in their original order.
func InterfacesDistinct(a []interface{}) (set []interface{}) {
	set = ameda_loc.InterfacesCopy(a)
	ameda_loc.InterfacesDistinct(&set, true)
	return set
}

// SetToInterfaces sets a element to the interface{} set.
func SetToInterfaces(set []interface{}, a interface{}) []interface{} {
	return ameda_loc.InterfacesPushDistinct(set, a)
}

// RemoveFromInterfaces removes the first element from the interface{} set.
func RemoveFromInterfaces(set []interface{}, a interface{}) []interface{} {
	ameda_loc.InterfacesRemoveFirst(&set, a)
	return set
}

// RemoveAllFromInterfaces removes all the a element from the interface{} set.
func RemoveAllFromInterfaces(set []interface{}, a interface{}) []interface{} {
	ameda_loc.InterfacesRemoveEvery(&set, a)
	return set
}
