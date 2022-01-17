package forms

type errors map[string][]string

func (e errors) Add(k, v string) {
	e[k] = append(e[k], v)
}

func (e errors) Get(k string) string {
	es := e[k]
	if len(es) == 0 {
		return ""
	}

	return es[0]
}
