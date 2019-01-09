package main

type Headers []string

func (h *Headers) String() string {
	return ""
}

func (h *Headers) Set(v string) error {
	*h = append(*h, v)
	return nil
}
