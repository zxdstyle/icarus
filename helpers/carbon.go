package helpers

import "github.com/golang-module/carbon/v2"

var (
	Carbon     carbon.Carbon
	CarbonLang *carbon.Language
)

func init() {
	CarbonLang = carbon.NewLanguage()
	CarbonLang.SetLocale("en")
	Carbon.SetLanguage(CarbonLang)
}
