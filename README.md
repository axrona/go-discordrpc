# go-discordrpc

⚡ Cross-Platform Discord Rich Presence implementation in Go

[![Go Reference](https://pkg.go.dev/badge/github.com/xeyossr/go-discordrpc.svg)](https://pkg.go.dev/github.com/xeyossr/go-discordrpc)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](LICENSE)

`go-discordrpc` is a Go module for interacting with Discord Rich Presence (RPC), inspired by and forked from [rich-go](https://github.com/hugolgst/rich-go).  
It provides a simple, cross-platform way to set custom activities, buttons, timestamps, and more, all from Go applications.

--- 

## Installation

You can install `go-discordrpc` using `go get`:

```bash
go get github.com/xeyossr/go-discordrpc/client
```

---

## Usage

Import the client and create a new Discord RPC client instance. Then log in, set your activity, and optionally logout:

```golang
package main

import (
    "fmt"
    "time"

    "github.com/xeyossr/go-discordrpc/client"
)

func main() {
    // Create a new Discord RPC client
    c := client.NewClient("YOUR_DISCORD_APP_ID")

    // Login to Discord
    if err := c.Login(); err != nil {
        panic(fmt.Errorf("login failed: %w", err))
    }
    fmt.Println("Logged in!")

    // Set activity
    now := time.Now()
    if err := c.SetActivity(client.Activity{
        State:      "Playing Go Modules",
        Details:    "Building cross-platform Discord RPC",
        LargeImage: "largeimageid",
        LargeText:  "Large image tooltip",
        SmallImage: "smallimageid",
        SmallText:  "Small image tooltip",
        Timestamps: &client.Timestamps{
            Start: &now,
        },
        Buttons: []*client.Button{
            {
                Label: "GitHub",
                Url:   "https://github.com/xeyossr/go-discordrpc",
            },
        },
    }); err != nil {
        panic(fmt.Errorf("set activity failed: %w", err))
    }
    fmt.Println("Activity set!")

    // Sleep to allow Discord to display the presence
    fmt.Println("Sleeping for 10 seconds...")
    time.Sleep(10 * time.Second)

    // Logout
    if err := c.Logout(); err != nil {
        panic(fmt.Errorf("logout failed: %w", err))
    }
    fmt.Println("Logged out!")
}
```

--- 

## Contributing / Support

If you want to contribute, report a bug, or suggest a feature, please open an issue or a pull request.
All contributions are welcome!

--- 

## License
This project is licensed under the [GNU General Public License v3](https://opensource.org/license/GPL-3.0). See the [LICENSE](LICENSE) file for details.

---

## Thanks

Special thanks to [rich-go](https://github.com/hugolgst/rich-go) for the original idea and inspiration!
