package language

import (
	"fmt"

	"github.com/tidwall/gjson"
)

// Language is used to hold information about a language
type Language struct {
	// name is the name of the language
	name string
	// icon is what would be displayed for the Language
	// if you do not want to have an icon or use one
	// then just leave this blank
	icon string
	// translations is a string of the json content
	// in the translations files
	translations string
}

// New will return a new Language with no translations,
// this should only be used when you are regis
func New(name string, icon string) Language {
	return Language{
		name:         name,
		icon:         icon,
		translations: "",
	}
}

// WithTranslations will set the languages translations to
// the given string. The string should be valid JSON as that
// is what is used to be translated
func (l Language) WithTranslations(tra string) Language {
	l.translations = tra
	return l
}

// Name ...
func (l Language) Name() string {
	return l.name
}

// Icon ...
func (l Language) Icon() string {
	return l.icon
}

// Translate will return the value of the given key unformatted
// if the given key is not found then it will simply return the key
func (l Language) Translate(key string) string {
	res := gjson.Get(l.translations, key)
	if res.String() == "" {
		return key
	}
	return res.String()
}

// Translatef will return a formatted version of the string
// if the given key is not found then it will simply return the key
func (l Language) Translatef(key string, args ...any) string {
	res := gjson.Get(l.translations, key)
	if res.String() == "" {
		return key
	}
	return fmt.Sprintf(res.String(), args...)
}
