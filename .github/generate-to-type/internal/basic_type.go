package internal

type TestCase struct {
	Args       string
	ArgsType   string
	Want       string
	ErrContain string

	ArgsLite string
	WantLite string
}

type ToTypeReq struct {
	Type                string
	ZeroVal             string
	OneVal              string
	MaxVal              string
	ConvertTypes        []string
	OverflowTypes       []string
	GroupTypeMaxValType string

	TypeTitle string

	TestCases []*TestCase
}
