package constants

import "os"

// Base directory of snippet.
var userHomeDir, _ = os.UserHomeDir()
var BaseDir = userHomeDir + "/meok"

// github username and repo in order to fetch snippets.
var Owner = "hoehwa"
var Repository = "snippet"

// urls for hosted websites.
const GardenBaseurl = "https://mindulle.github.io/garden"
