package weight

import (
	"reflect"
	"testing"
)

func TestFromString(t *testing.T) {
	str := "123.45lb"
	w := FromString(str)
	if !reflect.DeepEqual(w, Weight{Unit: "lb", Count: 123.45}) {
		t.Error("Expected a different weight!")
	}
}

func TestToString(t *testing.T) {
	wgt := Weight{Unit: "lb", Count: 123.45}
	if !reflect.DeepEqual(ToString(wgt), "123.45 lb") {
		t.Error("Expected a different string! Got ", ToString(wgt))
	}
}

func TestConvertToFromOzToLb(t *testing.T) {
	wgt := Weight{Unit: "oz", Count: 16}
	if converted, _ := ConvertTo(wgt, "lb"); converted.Count != 1 {
		t.Error("Expected a different weight! Got ", converted.Unit, converted.Count)
	}
}

func TestConvertToFromOzToGr(t *testing.T) {
	wgt := Weight{Unit: "oz", Count: 3.5274}
	if converted, _ := ConvertTo(wgt, "gr"); converted.Count != 100 {
		t.Error("Expected a different weight! Got ", converted.Unit, converted.Count)
	}
}

func TestConvertToFromOzToTons(t *testing.T) {
	wgt := Weight{Unit: "oz", Count: 32000}
	if converted, err := ConvertTo(wgt, "ton"); err == nil {
		t.Error("Expected an error! Got ", ToString(converted))
	}
}

func TestConvertToFromInvalidToOz(t *testing.T) {
	wgt := Weight{Unit: "invalid", Count: 1}
	if converted, err := ConvertTo(wgt, "oz"); err == nil {
		t.Error("Expected an error! Got ", ToString(converted))
	}
}
