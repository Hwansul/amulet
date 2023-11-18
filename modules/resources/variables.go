package resources

import "os"

var userHomeDir, _ = os.UserHomeDir()

// Base directory of snippet.
var BaseDir = userHomeDir + "/snippet"
