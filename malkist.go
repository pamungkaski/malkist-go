package malkist

// Malkist represent struct of the library to save key and access other function.
type Malkist struct {
	Key string
}

// ChangeAPIKey change the API key of the Malkist struct.
func (m *Malkist) ChangeAPIKey(key string) {
	m.Key = key
}
