package db

const (
	LEVEL_DB = "leveldb"
	FILE_DB  = "filedb"
)

func Defaults() DbConfig {
	return DbConfig{
		Type: LEVEL_DB,
	}
}
