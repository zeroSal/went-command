package command

import (
	"fmt"
)

type Registry struct {
	commands map[string]Interface
}

func NewRegistry() *Registry {
    return &Registry{
		commands: make(map[string]Interface),
	}
}

func (r *Registry) Register(s Interface) {
    r.commands[s.GetHeader().Use] = s
}

func (r *Registry) Get(name string) (Interface, error) {
    s, ok := r.commands[name]
    if !ok {
        return nil, fmt.Errorf("command %q not found", name)
    }

    return s, nil
}

func (r *Registry) All() []Interface {
    out := make([]Interface, 0, len(r.commands))
    for _, s := range r.commands {
        out = append(out, s)
    }

    return out
}
