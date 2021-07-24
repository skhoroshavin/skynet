package mysql_storage

func newTestStorage() *Storage {
	storage, err := NewStorage(EnvConfig())
	if err != nil {
		panic(err)
	}

	if err := storage.CleanUp(); err != nil {
		panic(err)
	}

	if err := storage.Setup(); err != nil {
		panic(err)
	}

	return storage
}