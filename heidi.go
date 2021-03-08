package heidi

type Reducer func(res Row)

type Heidi interface {
	Tag(tag string) Heidi
	Flush() Heidi
	Filter(tag string, p Filter) Heidi
	Reduce(r Reducer, data ...interface{}) error
}

type heidi struct {
	tag        string
	processors map[string]Filter
}

func (h *heidi) Tag(tag string) Heidi {
	h.tag = tag

	return h
}

func (h *heidi) Flush() Heidi {
	h.processors = map[string]Filter{}

	return h
}

func (h *heidi) Filter(tag string, p Filter) Heidi {
	h.processors[tag] = p

	return h
}

func (h *heidi) Reduce(r Reducer, data ...interface{}) error {
	panic("implement me")
}

func New() Heidi {
	return &heidi{
		tag: "heidi",
		processors: map[string]Filter{
			"suppress": Suppress,
			"censor":   Censor,
		},
	}
}
