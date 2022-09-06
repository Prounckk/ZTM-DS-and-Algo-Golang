package array

type MyArray struct {
	data []interface{}
	size int
}

func New(values ...interface{}) *MyArray {
	myArray := &MyArray{}
	length := len(values)
	if length != 0 {
		myArray.Add(values...)
	}
	return myArray
}

func (m *MyArray) Add(values ...interface{}) {
	m.resize(len(values))
	for _, value := range values {
		m.data[m.size] = value
		m.size++
	}
}

func (m *MyArray) Remove(index int) {
	m.data[index] = nil
	copy(m.data[index:], m.data[index+1:m.size])
	m.size--
}

func (m *MyArray) resize(cap int) {
	reSizedData := make([]interface{}, cap)
	copy(reSizedData, m.data)
	m.data = reSizedData
}
