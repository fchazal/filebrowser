package storage

import (
	"github.com/fchazal/filebrowser/auth"
	"github.com/fchazal/filebrowser/settings"
	"github.com/fchazal/filebrowser/share"
	"github.com/fchazal/filebrowser/users"
)

// Storage is a storage powered by a Backend which makes the necessary
// verifications when fetching and saving data to ensure consistency.
type Storage struct {
	Users    *users.Storage
	Share    *share.Storage
	Auth     *auth.Storage
	Settings *settings.Storage
}
