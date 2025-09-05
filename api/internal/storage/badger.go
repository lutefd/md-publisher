package storage

import (
	"encoding/json"

	"github.com/dgraph-io/badger/v4"
)

type Store interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte) error
	Delete(key string) error
	Close() error
	ListKeys() ([]string, error)
}

type BadgerStore struct {
	db *badger.DB
}

func NewBadgerStore(dataPath string) (*BadgerStore, error) {
	opts := badger.DefaultOptions(dataPath)
	opts.Logger = nil

	db, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}

	return &BadgerStore{db: db}, nil
}

func (s *BadgerStore) Get(key string) ([]byte, error) {
	var value []byte
	err := s.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}

		value, err = item.ValueCopy(nil)
		return err
	})

	return value, err
}

func (s *BadgerStore) Set(key string, value []byte) error {
	return s.db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(key), value)
	})
}

func (s *BadgerStore) Delete(key string) error {
	return s.db.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(key))
	})
}

func (s *BadgerStore) Close() error {
	return s.db.Close()
}

func (s *BadgerStore) ListKeys() ([]string, error) {
	var keys []string

	err := s.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = false

		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			key := string(item.Key())
			keys = append(keys, key)
		}

		return nil
	})

	return keys, err
}

type Note struct {
	ID       string                 `json:"id"`
	Content  string                 `json:"content"`
	Metadata map[string]interface{} `json:"metadata"`
}

type NoteStore struct {
	store Store
}

func NewNoteStore(store Store) *NoteStore {
	return &NoteStore{store: store}
}

func (ns *NoteStore) SaveNote(note Note) error {
	ExtractFrontmatter(&note)

	data, err := json.Marshal(note)
	if err != nil {
		return err
	}

	return ns.store.Set(note.ID, data)
}

func (ns *NoteStore) GetNote(id string) (Note, error) {
	var note Note

	data, err := ns.store.Get(id)
	if err != nil {
		return note, err
	}

	err = json.Unmarshal(data, &note)
	return note, err
}

func (ns *NoteStore) DeleteNote(id string) error {
	return ns.store.Delete(id)
}
func (ns *NoteStore) ListNotes() ([]Note, error) {
	keys, err := ns.store.ListKeys()
	if err != nil {
		return nil, err
	}

	var notes []Note
	for _, key := range keys {
		note, err := ns.GetNote(key)
		if err != nil {
			continue
		}
		notes = append(notes, note)
	}

	return notes, nil
}
