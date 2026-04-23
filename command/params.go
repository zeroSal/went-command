package command

type Params struct {
	flags map[string]any
	args  []string
}

func newParams() *Params {
	return &Params{
		flags: make(map[string]any),
	}
}

func (p *Params) SetFlag(name string, value any) {
	p.flags[name] = value
}

func (p *Params) SetArgs(args []string) {
	p.args = args
}

func (p *Params) Bool(name string) bool {
	if f, ok := p.flags[name]; ok {
		if b, ok := f.(*bool); ok {
			return *b
		}
	}
	return false
}

func (p *Params) Int(name string) int {
	if f, ok := p.flags[name]; ok {
		if i, ok := f.(*int); ok {
			return *i
		}
	}
	return 0
}

func (p *Params) Int64(name string) int64 {
	if f, ok := p.flags[name]; ok {
		if i, ok := f.(*int64); ok {
			return *i
		}
	}
	return 0
}

func (p *Params) String(name string) string {
	if f, ok := p.flags[name]; ok {
		if s, ok := f.(*string); ok {
			return *s
		}
	}
	return ""
}

func (p *Params) StringOr(name string, fallback string) string {
	if f, ok := p.flags[name]; ok {
		if s, ok := f.(*string); ok && *s != "" {
			return *s
		}
	}
	return fallback
}

func (p *Params) Arg(index int) string {
	if index >= 0 && index < len(p.args) {
		return p.args[index]
	}
	return ""
}

func (p *Params) Args() []string {
	return p.args
}

func (p *Params) ArgsCount() int {
	return len(p.args)
}

func (p *Params) HasFlag(name string) bool {
	_, ok := p.flags[name]
	return ok
}
