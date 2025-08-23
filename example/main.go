package main

import (
	"fmt"
	"time"

	"github.com/xeyossr/go-discordrpc/client"
)

func main() {
	// 1. Create a new Discord RPC client
	c := client.NewClient("DISCORD_APP_ID")

	// 2. Login to Discord via IPC
	if err := c.Login(); err != nil {
		panic(fmt.Errorf("login failed: %w", err))
	}
	fmt.Println("Logged in!")

	// 3. Get the current time for timestamps
	now := time.Now()

	// 4. Set Rich Presence activity
	if err := c.SetActivity(client.Activity{
		State:      "Playing Chess with Ai Hoshino",
		Details:    "If you like Ai Hoshino, we are 🤝",
		LargeImage: "largeimageid",
		LargeText:  "Large image tooltip",
		SmallImage: "smallimageid",
		SmallText:  "Small image tooltip",
		Party: &client.Party{
			ID:         "-1",
			Players:    5,
			MaxPlayers: 10,
		},
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

	// 5. Sleep to allow Discord to display the presence
	fmt.Println("Sleeping for 10 seconds...")
	time.Sleep(10 * time.Second)

	// 6. Logout from Discord and close the IPC socket
	if err := c.Logout(); err != nil {
		panic(fmt.Errorf("logout failed: %w", err))
	}
	fmt.Println("Logged out!")
}
