package mocks

import "fmt"
import "reflect"

type ISum interface {
	Sum(x, y int) int
}

type Sum struct{}

func (s *Sum) Sum(x, y int) int {
	return x + y
}

type SumMock struct {
	calls []*Call
}

type Call struct {
	Function string
	Args     []interface{}
	Returns  []interface{}
}

func (s *SumMock) EXPECT() *Call {
	c := new(Call)
	s.calls = append(s.calls, c)
	return c
}

func (c *Call) Sum(args ...interface{}) *Call {
	c.Function = "Sum"
	c.Args = args
	return c
}

func (c *Call) Return(args ...interface{}) *Call {
	c.Returns = args
	return c
}

func (s *SumMock) Sum(x, y int) int {
	for _, c := range s.calls {
		if c.Function == "Sum" {
			if len(c.Args) != 2 {
				panic("Wrong args count")
			}
			if !reflect.DeepEqual(x, c.Args[0]) {
				panic("Wrong argument at index 0")
			}
			if !reflect.DeepEqual(y, c.Args[1]) {
				panic("Wrong argument at index 1")
			}
			if len(c.Returns) != 0 {
				if len(c.Returns) != 1 {
					panic("Wrong return count")
				}
				if ret, ok := c.Returns[0].(int); ok {
					return ret
				}
			}
			return 0
		}
	}
	panic("Unexpected call")
}

func main() {
	var s ISum
	sMock := new(SumMock)
	sMock.EXPECT().Sum(2, 2).Return(4)
	s = sMock
	fmt.Println(s.Sum(2, 2))
}
