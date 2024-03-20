package utility

import (
	"path/filepath"

	"github.com/charliego3/assistant/enums"
	"github.com/progrium/macdriver/macos/foundation"
)

func SupportPath(sub ...string) string {
	url := foundation.FileManager_DefaultManager().
		URLForDirectoryInDomainAppropriateForURLCreateError(
			foundation.ApplicationSupportDirectory,
			foundation.UserDomainMask,
			nil,
			true,
			nil,
		)

	return filepath.Join(append([]string{url.Path(), enums.Identifier}, sub...)...)
}
