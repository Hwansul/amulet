package constants

import "os"

var userHomeDir, _ = os.UserHomeDir()

// Base directory of snippet.
var BaseDir = userHomeDir + "/snippet"

// url of my digital garden.
const GardenBaseurl = "https://mindulle.github.io/garden"
