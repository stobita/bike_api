package gateway

type SqlHandler interface {
	Query(string, ...interface{}) (Rows, error)
}

type Rows interface {
	Next() bool
	StructScan(interface{}) error
}
