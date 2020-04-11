/*

The toggl command will display a user's Toggl account information.

Usage:
    toggl API_TOKEN

The API token can be retrieved from a user's account information page at toggl.com.

*/
package main

import (
	"tgl_cli/subcommand"
)

func main() {
	subcommand.Execute()
}