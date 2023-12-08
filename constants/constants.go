package constants

import "os"

// Base directory of snippet.
var userHomeDir, _ = os.UserHomeDir()
var BaseDir = userHomeDir + "/amulet"

// github username and repo in order to fetch snippets.
var Owner = "hwansul"
var Repository = "snippet"

// urls for hosted websites.
const GardenBaseurl = "https://mindulle.github.io/garden"
