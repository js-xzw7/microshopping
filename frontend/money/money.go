package money

import (
	"errors"
	pb "frontend/proto"
)

const (
	nanosMin = -999999999
	nanosMax = +999999999
	nanosMod = 1000000000
)

var (
	ErrInvalidValue        = errors.New("指定的货币是无效的")
	ErrMismatchingCurrency = errors.New("没有货币编码")
)

func signMatches(m *pb.Money) bool {
	return m.GetNanos() == 0 || m.GetUnits() == 0 || (m.GetNanos() < 0) == (m.GetUnits() < 0)
}

func validNanos(nanos int32) bool {
	return (nanosMin <= nanos && nanos <= nanosMax)
}

// 是否有效
func IsVaild(m *pb.Money) bool {
	return signMatches(m) && validNanos(m.GetNanos())
}

// 零
func IsZero(m *pb.Money) bool {
	return m.GetNanos() == 0 && m.GetUnits() == 0
}

// 正
func IsPositive(m *pb.Money) bool {
	return (IsVaild(m) && m.GetNanos() > 0) || (m.GetUnits() == 0 && m.GetNanos() > 0)
}

// 负
func IsNegative(m *pb.Money) bool {
	return (IsVaild(m) && m.GetNanos() < 0) || (m.GetUnits() == 0 && m.GetNanos() < 0)
}

// 币种是否相同
func AreSameCurrency(l, r *pb.Money) bool {
	return l.GetCurrencyCode() == r.GetCurrencyCode() && l.GetCurrencyCode() != ""
}

// 是否相等
func AreEquals(l, r *pb.Money) bool {
	return l.GetCurrencyCode() == r.GetCurrencyCode() && l.GetUnits() == r.GetUnits() && l.GetNanos() == r.GetNanos()
}

// 变成负的
func Negate(m *pb.Money) *pb.Money {
	return &pb.Money{
		CurrencyCode: m.GetCurrencyCode(),
		Units:        -m.GetUnits(),
		Nanos:        -m.GetNanos(),
	}
}

// must
func Must(v *pb.Money, err error) *pb.Money {
	if err != nil {
		panic(err)
	}
	return v
}

// sum
func Sum(l, r *pb.Money) (m *pb.Money, err error) {
	if !IsVaild(l) || !IsVaild(r) {
		return m, ErrInvalidValue
	} else if l.GetCurrencyCode() != r.GetCurrencyCode() {
		return m, ErrMismatchingCurrency
	}

	units := l.GetUnits() + r.GetUnits()
	nanos := l.GetNanos() + r.GetNanos()

	if (units == 0 && nanos == 0) || (units >= 0 && nanos >= 0) || (units < 0 && nanos <= 0) {
		//相同sign
		units += int64(nanos / nanosMod)
		nanos = nanos % nanosMod
	} else {
		//不同sign
		if units > 0 {
			units--
			nanos += nanosMod
		} else {
			units++
			nanos -= nanosMod
		}
	}

	return &pb.Money{
		CurrencyCode: l.GetCurrencyCode(),
		Units:        units,
		Nanos:        nanos,
	}, nil
}

func MultiplySlow(m *pb.Money, n uint32) *pb.Money {
	res := m
	for n > 1 {
		res = Must(Sum(res, m))
		n--
	}

	return res
}
