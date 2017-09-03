# XKCDPass

xkcdpass generates secure random passwords that are (relatively) easy to memorize.

It uses a [dictionary provided specifically for passphrase generation by the EFF](https://www.eff.org/deeplinks/2016/07/new-wordlists-random-passphrases) and has an option for checking strength of the generated passwords using the [zxcvbn library](https://github.com/nbutton23/zxcvbn-go).

## Basic usage
```go
import (
  "github.com/happeens/xkcdpass"
)

// Create a 4-word password with strength score 4
password := xkcdpass.GenerateDefault()
```

For more detailed usage options, check the docs.
