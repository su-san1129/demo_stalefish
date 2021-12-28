type CharFilter interface {
	Filter(string) string
}

type MappingCharFilter struct {
	mapper map[string]string
}

func (c MappingCharFilter) Filter(s string) string {
	for k, v := range c.mapper {
		s = strings.Replace(s, k, v, -1)
	}
	return s;
}