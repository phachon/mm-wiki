package utils

import (
	"fmt"
	"testing"
)

func TestNewVersionCompare(t *testing.T) {

}

func TestVersionCompare_ConvertIntList(t *testing.T) {
	vc := NewVersionCompare("v")
	l := vc.ConvertIntList("v3.25.10")
	fmt.Println(l)
}

func TestVersionCompare_Eq(t *testing.T) {
	vc := NewVersionCompare("v")
	if vc.Eq("v3.25.10", "v1.10.9") {
		t.Fail()
	}
	if !vc.Eq("v1.10.9", "v1.10.9") {
		t.Fail()
	}
}

func TestVersionCompare_Neq(t *testing.T) {
	vc := NewVersionCompare("v")
	if !vc.Neq("v3.25.10", "v1.10.9") {
		t.Fail()
	}
	if vc.Neq("v1.10.9", "v1.10.9") {
		t.Fail()
	}
}

func TestVersionCompare_Gt(t *testing.T) {
	vc := NewVersionCompare("v")
	if !vc.Gt("v3.25.10", "v1.10.9") {
		t.Fatal()
	}
	if !vc.Gt("v3.25.10", "v3.10.9") {
		t.Fatal()
	}
	if !vc.Gt("v3.25.10", "v3.25.9") {
		t.Fatal()
	}
	if vc.Gt("v3.25.10", "v3.25.10") {
		t.Fatal()
	}
	if vc.Gt("v3.25.10", "v3.75.10") {
		t.Fail()
	}
}

func TestVersionCompare_Gte(t *testing.T) {
	vc := NewVersionCompare("v")
	if !vc.Gte("v3.25.10", "v1.10.9") {
		t.Fatal()
	}
	if !vc.Gte("v3.25.10", "v3.10.9") {
		t.Fatal()
	}
	if !vc.Gte("v3.25.10", "v3.25.9") {
		t.Fatal()
	}
	if !vc.Gte("v3.25.10", "v3.25.10") {
		t.Fatal()
	}
	if vc.Gte("v3.25.10", "v3.75.10") {
		t.Fail()
	}
}
