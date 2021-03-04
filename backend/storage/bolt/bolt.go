package bolt

import (
	"github.com/asdine/storm"

	"github.com/fchazal/filebrowser/auth"
	"github.com/fchazal/filebrowser/settings"
	"github.com/fchazal/filebrowser/share"
	"github.com/fchazal/filebrowser/storage"
	"github.com/fchazal/filebrowser/users"
)

// NewStorage creates a storage.Storage based on Bolt DB.
func NewStorage(db *storm.DB) (*storage.Storage, error) {
	userStore := users.NewStorage(usersBackend{db: db})
	shareStore := share.NewStorage(shareBackend{db: db})
	settingsStore := settings.NewStorage(settingsBackend{db: db})
	authStore := auth.NewStorage(authBackend{db: db}, userStore)

	err := save(db, "version", 2)
	if err != nil {
		return nil, err
	}

	return &storage.Storage{
		Auth:     authStore,
		Users:    userStore,
		Share:    shareStore,
		Settings: settingsStore,
	}, nil
}
