package command

type Base struct {
	Params *Params
}

func (b Base) GetBoolFlag(fieldName string) bool {
	if b.Params == nil {
		return false
	}
	return b.Params.Bool(fieldName)
}

func (b Base) GetIntFlag(fieldName string) int {
	if b.Params == nil {
		return 0
	}
	return b.Params.Int(fieldName)
}

func (b Base) GetInt64Flag(fieldName string) int64 {
	if b.Params == nil {
		return 0
	}
	return b.Params.Int64(fieldName)
}

func (b Base) GetStringFlag(fieldName string) string {
	if b.Params == nil {
		return ""
	}
	return b.Params.String(fieldName)
}

func (b Base) GetStringFlagOr(fieldName string, fallback string) string {
	if v := b.GetStringFlag(fieldName); v != "" {
		return v
	}
	return fallback
}

func (b Base) GetArg(index int) string {
	return b.Params.Arg(index)
}

func (b Base) Args() []string {
	return b.Params.Args()
}

func (b Base) ArgsCount() int {
	return b.Params.ArgsCount()
}

func (b Base) HasFlag(fieldName string) bool {
	return b.Params.HasFlag(fieldName)
}
