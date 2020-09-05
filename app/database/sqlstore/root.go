package sqlstore

import "github.com/netorissi/wk_api_go/entities"

type StoreResult struct {
	Data interface{}
	Err  *entities.AppError
}

type StoreChannel chan StoreResult

type LayeredStoreDatabaseLayer interface {
	Store
}

type LayeredStore struct {
	DatabaseLayer LayeredStoreDatabaseLayer
}

type Store interface {
	Close()
}

func NewLayeredStore(db LayeredStoreDatabaseLayer) Store {
	store := &LayeredStore{
		DatabaseLayer: db,
	}

	return store
}

func Do(f func(result *StoreResult)) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		f(&result)
		storeChannel <- result
		close(storeChannel)
	}()
	return storeChannel
}

func (s *LayeredStore) Close() {
	s.DatabaseLayer.Close()
}
