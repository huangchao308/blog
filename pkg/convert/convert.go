package convert

import "strconv"

type StrTo string

func (s StrTo) String() string {
	return string(s)
}

func (s StrTo) Int() (int, error) {
	return strconv.Atoi(s.String())
}

func (s StrTo) MustInt() int {
	i, _ := s.Int()
	return i
}

func (s StrTo) Uint32() (uint32, error) {
	v, err := strconv.ParseUint(s.String(), 10, 32)
	return uint32(v), err
}

func (s StrTo) MustUint32() uint32 {
	v, _ := s.Uint32()
	return v
}
