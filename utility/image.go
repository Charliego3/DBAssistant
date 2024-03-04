package utility

import "github.com/progrium/macdriver/macos/appkit"

func SymbolImage(name string, cfg ...appkit.ImageSymbolConfiguration) appkit.Image {
	image := appkit.Image_ImageWithSystemSymbolNameAccessibilityDescription(name, name)
	for _, configure := range cfg {
		image = image.ImageWithSymbolConfiguration(configure)
	}
	return image
}
