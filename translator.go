package translator

import (
	"fmt"
	"os"

	"github.com/Brougud/translator/language"

	lng "golang.org/x/text/language"
)

// path is where the translation files are located
var path string

// languages is a map containing the locale of the language
// and maps that to a string of the json value
var languages = make(map[lng.Tag]language.Language)

// Initalize will initalize the library, this is where
// you put all of your settings that you want, and
// how you want to use your languages
//
// dir is the directory containing all of the translation
// files
//
// if the directory does not exist then it will
// create the directory for you, if there is
// an error while creating the directory then
// it will panic with the error
func Initalize(dir string) {
  path = dir

  // the directory could not be found then make
  // it for the user
  if _, err := os.Stat(path); os.IsNotExist(err) {
    err := os.MkdirAll(path, os.ModePerm)
    if err != nil {
      panic(err)
    }
  }
}

// Register takes in the locale of the language,
// this should be the naming convention of your
// files
//
// ex: if your file name is "en_US.json" then 
// your locale should be "en_US"
func Register(locale string, lang language.Language) (language.Language, error) {
  dat, err := os.ReadFile(fmt.Sprintf("%v/%v.json", path, locale))
  if err != nil {
    return language.Language{}, err
  }
  lang = lang.WithTranslations(string(dat))
  languages[lng.MustParse(locale)]=lang
  return lang, nil
}

// All will return all of the languages that are registered
func All() []lng.Tag {
  final := []lng.Tag{}
  for l := range languages {
    final = append(final, l)
  }
  return final
}

// FromString returns a language from a locale string
// if the language could not be found it will return 
// false as the second value
func FromString(lang string) (language.Language, bool) {
  tag := lng.MustParse(lang)
  lan, ok := languages[tag]
  return lan, ok
}
