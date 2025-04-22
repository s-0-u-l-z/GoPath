package utils

import _ "embed"
import "strings"

//go:embed wordlist.txt
var rawWordlist string

var Wordlist = strings.Split(strings.TrimSpace(rawWordlist), "\n")
