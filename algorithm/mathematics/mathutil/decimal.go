package mathutil

import (
	"fmt"
	"strconv"

	"github.com/shopspring/decimal"
)

type DecimalOpt int8

const (
	DECIMAL_ADD DecimalOpt = 1 // 加
	DECIMAL_SUB DecimalOpt = 2 // 减
	DECIMAL_MUL DecimalOpt = 3 // 乘
	DECIMAL_DIV DecimalOpt = 4 // 除
)

func DecimalCalculate(opt DecimalOpt, f1, f2 float64, f3 ...float64) (float64, error) {

	d1 := decimal.NewFromFloat(f1)
	d2 := decimal.NewFromFloat(f2)

	var d decimal.Decimal
	switch opt {

	case DECIMAL_ADD: // 加
		d = d1.Add(d2)
		for _, v := range f3 {
			d = d.Add(decimal.NewFromFloat(v))
		}

	case DECIMAL_SUB: // 减
		d = d1.Sub(d2)
		for _, v := range f3 {
			d = d.Sub(decimal.NewFromFloat(v))
		}

	case DECIMAL_MUL: // 乘
		d = d1.Mul(d2)
		for _, v := range f3 {
			d = d.Mul(decimal.NewFromFloat(v))
		}

	case DECIMAL_DIV: // 除
		d = d1.Div(d2)
		for _, v := range f3 {
			d = d.Div(decimal.NewFromFloat(v))
		}
	}

	f, _ := d.Float64()
	return strconv.ParseFloat(fmt.Sprintf("%.8f", f), 64) // 保留 8 位小数
}
