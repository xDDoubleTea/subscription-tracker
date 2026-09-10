package main

import (
	"fmt"
	"io"

	"github.com/xDDoubleTea/saas-payment-tracker/core"
)

func main() {
	fmt.Println("SAAS Payment Tracker v1.0")
	for {
		core.Prompt()

		operation := ""
		n, err := fmt.Scanf("%s", &operation)
		if err == io.EOF {
			fmt.Println("Exiting...")
			return
		}

		if n != 1 {
			fmt.Errorf("Invalid input.")
			continue
		}

		switch operation {
		case "add":
			if err := core.Add(); err != nil {
				fmt.Println(err.Error())
			}

		case "ls":
			fmt.Println()
			fmt.Println("------Subscriptions------")
			for _, val := range core.Subscriptions {
				fmt.Println(core.SubscriptiontoString(val))
			}
			fmt.Println("-------------------------")
			fmt.Println()
		case "u":
			if err := core.U(); err != nil {
				fmt.Println(err.Error())
			}

		case "get":
			if err := core.Get(); err != nil {
				fmt.Println(err.Error())
			}
		case "rm":
			if err := core.Rm(); err != nil {
				fmt.Println(err.Error())
			}

		case "q":
			return
		default:
			fmt.Println("Invalid operation.")
		}
	}
}
