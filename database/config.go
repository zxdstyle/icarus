package database

type Config struct {
	MaxOpenConns    int
	MaxIdleConns    int
	DataSource      string
	SlaveDataSource string
}
