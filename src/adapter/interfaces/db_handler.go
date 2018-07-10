package interfaces

type DBHandler interface {
	Execute(string, ...interface{}) (Result, error)
	Query(string, ...interface{}) (Row, error)
	CloseConn()
}

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type Row interface {
	Scan(...interface{}) error
	Next() bool
	Close() error
}
