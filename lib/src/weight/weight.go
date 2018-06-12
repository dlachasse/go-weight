package weight

import (
	"errors"
	"math"
	"regexp"
	"strconv"
)

var r = regexp.MustCompile(`\A(?P<count>\d*\.\d*)\s*(?P<unit>\w*)`)

type Weight struct {
	Unit  string
	Count float64
}

var conversions = map[string]map[string]float64{
	"gr": map[string]float64{
		"gr": 1.0,
		"lb": 0.002204623,
		"kg": 0.001,
		"oz": 0.03527399072294,
	},
	"kg": map[string]float64{
		"gr": 1000.0,
		"lb": 2.20462,
		"kg": 1.0,
		"oz": 35.274,
	},
	"lb": map[string]float64{
		"gr": 453.59237,
		"lb": 1.0,
		"kg": 0.45359237,
		"oz": 16.0,
	},
	"oz": map[string]float64{
		"gr": 28.3495,
		"lb": 0.0625,
		"kg": 0.0283495,
		"oz": 1.0,
	},
}

func FromString(str string) Weight {
	var unit string
	var count float64

	match := r.FindStringSubmatch(str)
	for i, name := range r.SubexpNames() {
		switch name {
		case "count":
			count, _ = strconv.ParseFloat(match[i], 64)
		case "unit":
			unit = match[i]
		}
	}
	return Weight{Unit: unit, Count: count}
}

func ToString(wgt Weight) string {
	return floatToString(wgt.Count) + " " + wgt.Unit
}

func ConvertTo(inWgt Weight, outUnit string) (Weight, error) {
	var err error
	var outCount float64
	if convertUnit, inExists := conversions[inWgt.Unit]; inExists {
		if conversionRate, outExists := convertUnit[outUnit]; outExists {
			outCount = roundPrecision(conversionRate*inWgt.Count, 4)
		} else {
			err = errors.New("Requested conversion unit unavailable")
		}
	} else {
		err = errors.New("Requested conversion unit unavailable")
	}

	return Weight{Count: outCount, Unit: outUnit}, err
}

func floatToString(input_num float64) string {
	return strconv.FormatFloat(roundPrecision(input_num, 4), 'f', -1, 64)
}

func round(f float64) float64 {
	return math.Floor(f + .5)
}

func roundPrecision(f float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return round(f*shift) / shift
}
