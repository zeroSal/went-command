package command

type BoolFlag struct {
	Name    string
	Default bool
	Usage   string
}

type IntFlag struct {
	Name    string
	Default int
	Usage   string
}

type Int64Flag struct {
	Name    string
	Default int64
	Usage   string
}

type StringFlag struct {
	Name    string
	Default string
	Usage   string
}

type Flags struct {
	Bool   []BoolFlag
	Int    []IntFlag
	Int64  []Int64Flag
	String []StringFlag
}

func NewFlags() *Flags {
	return &Flags{}
}

func (f *Flags) AddBool(name string, defaultValue bool, usage string) *Flags {
	f.Bool = append(f.Bool, BoolFlag{Name: name, Default: defaultValue, Usage: usage})
	return f
}

func (f *Flags) AddInt(name string, defaultValue int, usage string) *Flags {
	f.Int = append(f.Int, IntFlag{Name: name, Default: defaultValue, Usage: usage})
	return f
}

func (f *Flags) AddInt64(name string, defaultValue int64, usage string) *Flags {
	f.Int64 = append(f.Int64, Int64Flag{Name: name, Default: defaultValue, Usage: usage})
	return f
}

func (f *Flags) AddString(name string, defaultValue string, usage string) *Flags {
	f.String = append(f.String, StringFlag{Name: name, Default: defaultValue, Usage: usage})
	return f
}
