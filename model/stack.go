package model

type StrStack struct {
	s []string
}

func NewStrStack() *StrStack {
	return &StrStack{
		s: make([]string, 0),
	}
}

func (t *StrStack) Push(v ...string) {
	t.s = append(t.s, v...)
}

func (t *StrStack) Pop() {
	t.PopN(1)
}

func (t *StrStack) PopN(n int) {
	for i := 0; i < n; i++ {
		t.s = t.s[:len(t.s)-1]
	}
}

func (t *StrStack) S() []string {
	return t.s
}
