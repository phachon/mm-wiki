package utils

import (
	"strconv"
	"strings"
)

// version compare
// version example: v0.1.1

var (
	defaultPrefix  = "v"
	VersionCompare = NewVersionCompare(defaultPrefix)
)

// struct
type versionCompare struct {
	prefix string
}

// new versionCompare
func NewVersionCompare(prefix string) *versionCompare {
	return &versionCompare{
		prefix: prefix,
	}
}

// version1 == version2
func (vc *versionCompare) Eq(version1 string, version2 string) bool {
	return version1 == version2
}

// version1 != version2
func (vc *versionCompare) Neq(version1 string, version2 string) bool {
	return version1 != version2
}

// version1 > version2
func (vc *versionCompare) Gt(version1 string, version2 string) bool {
	v1Number := vc.ConvertIntList(version1)
	v2Number := vc.ConvertIntList(version2)
	for i, v := range v1Number {
		if v > v2Number[i] {
			return true
		}
		if v == v2Number[i] {
			continue
		}
		if v < v2Number[i] {
			return false
		}
	}
	return false
}

// version1 < version2
func (vc *versionCompare) Lt(version1 string, version2 string) bool {
	return vc.Gt(version2, version1)
}

// version1 >= version2
func (vc *versionCompare) Gte(version1 string, version2 string) bool {
	v1Number := vc.ConvertIntList(version1)
	v2Number := vc.ConvertIntList(version2)
	for i, v := range v1Number {
		if v > v2Number[i] {
			return true
		}
		if v == v2Number[i] {
			continue
		}
		if v < v2Number[i] {
			return false
		}
	}
	return true
}

// version1 <= version2
func (vc *versionCompare) Lte(version1 string, version2 string) bool {
	return vc.Gte(version2, version1)
}

// version string convert int list
func (vc *versionCompare) ConvertIntList(version string) []int {
	var realVersion = version
	if version[0:len(vc.prefix)] == vc.prefix {
		realVersion = version[len(vc.prefix):]
	}
	versionList := strings.Split(realVersion, ".")
	var l = make([]int, len(versionList))
	for i, v := range versionList {
		intValue, _ := strconv.Atoi(v)
		l[i] = intValue
	}
	return l
}
