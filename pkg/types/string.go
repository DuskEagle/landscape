package types

type StringInput interface {
	Await() string
}

type StringOutput struct{ *stringOutputInternal }

type stringOutputInternal struct {
	f func() string
}

func (s *stringOutputInternal) Await() string {
	return s.f()
}

func String(s string) StringInput {
	return &stringOutputInternal{
		f: func() string { return s },
	}
}

// TODO(joel): Make this hidden?
func NewStringOutput(f func() string) StringOutput {
	return StringOutput{&stringOutputInternal{f: f}}
}
