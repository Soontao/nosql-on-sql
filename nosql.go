package nosql

// NoSQL interface
type NoSQL struct {
	DBType string
}

func (n *NoSQL) insert(row interface{}) interface{} {
	return nil
}

func (n *NoSQL) update(filter interface{}, row interface{}) interface{} {
	return nil
}

func (n *NoSQL) delete(filter interface{}) interface{} {
	return nil
}

func (n *NoSQL) query(row interface{}) interface{} {
	return nil
}

// New NoSQL instance
func New(DBType string) *NoSQL {
	return &NoSQL{DBType}
}
