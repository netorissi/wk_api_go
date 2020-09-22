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

	Users() UsersStore
	Sessions() SessionsStore
}

type UsersStore interface {
	Create(user *entities.User) StoreChannel
}

type SessionsStore interface {
	Create(session *entities.Session) StoreChannel
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

func (s *LayeredStore) Users() UsersStore {
	return s.DatabaseLayer.Users()
}

func (s *LayeredStore) Sessions() SessionsStore {
	return s.DatabaseLayer.Sessions()
}
