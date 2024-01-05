package utility

import (
	"github.com/progrium/macdriver/macos/appkit"
)

func ColorWithAppearance(light, dark appkit.Color) appkit.Color {
	effected := appkit.Application_SharedApplication().EffectiveAppearance()
	if effected.BestMatchFromAppearancesWithNames([]appkit.AppearanceName{
		appkit.AppearanceNameAqua,
		appkit.AppearanceNameDarkAqua,
	}) == appkit.AppearanceNameDarkAqua {
		return dark
	}
	return light
}

func ColorWithRGBA(r, g, b, a float64) appkit.Color {
	return appkit.Color_ColorWithRedGreenBlueAlpha(r/255, g/255, b/255, a)
}
