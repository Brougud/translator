# Translator

This is a small library that can be used to translate JSON values, into strings, with multi language capability.

## Example

#### Directory
```
┏ main.go
┖ translations
  ┖ en_US.json
```

<small>en_US.json</small>

```json
{
    "value": "this is a value",
    "nested": {
        "value": "this is a nested value"
    },
    "formated": "this is a %v string"
}
```

<small>main.go</small>

```go
package main

import (
    "github.com/Brougud/translator"
    "github.com/Brougud/translator/language"
)

func main() {
    translator.Initalize("translations")
    lang, err := translator.Register("en_US", language.New("English", "icon/image"))
    if err != nil {
        panic(err)
    }
    fmt.Println(lang.Translate("value"))
    fmt.Println(lang.Translate("nested.value"))
    fmt.Println(lang.Translatef("formated", "cool"))
}
```

running this will output

```
this is a string
this is a nested value
this is a cool string
```