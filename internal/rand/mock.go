package rand

type MockSource struct {
	values []int
	index  int
}

func NewMockSource(values []int) *MockSource {
	return &MockSource{values: values}
}

func (m *MockSource) Intn(n int) (int, error) {
	v := m.values[m.index%len(m.values)] % n
	m.index++
	return v, nil
}
