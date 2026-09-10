package core

import "strings"
import "fmt"

var subscriptionID int64 = 0

func Prompt() {
	fmt.Println()
	fmt.Println("Type an operation: add, ls, u, get, rm, q")
	fmt.Print("> ")
}

func Add() error {
	var name string
	var expense float64
	var expenseCurrency string

	fmt.Print("Subscription name: ")
	_, errName := fmt.Scanf("%s", &name)
	if errName != nil {
		return errName
	}

	fmt.Print("Subscription expense: ")
	_, errExpense := fmt.Scanf("%f", &expense)
	if errExpense != nil {
		return errExpense
	}

	fmt.Print("Subscription expense currency: ")

	_, errCurrency := fmt.Scanf("%s", &expenseCurrency)
	if errCurrency != nil {
		return errCurrency
	}

	currency, err := NewCurrency(expenseCurrency)
	if err != nil {
		return err
	}
	entry := AddSubscription(subscriptionID, name, expense, currency)
	subscriptionID++

	fmt.Println("Entry added.")
	fmt.Println(SubscriptiontoString(entry))
	return nil
}

func Get() error {
	var id int64
	fmt.Print("Subscription id: ")

	_, err := fmt.Scanf("%d", &id)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	if entry, exists := GetSubscription(id); exists {
		fmt.Println(SubscriptiontoString(entry))
	} else {
		fmt.Printf("Subscription with id %d does not exist!\n", id)
	}
	return nil
}

func Rm() error {
	var id int64
	fmt.Print("Subscription id you want to remove: ")

	_, err := fmt.Scanf("%d", &id)
	if err != nil {
		return err
	}

	if entry, exists := GetSubscription(id); exists {
		fmt.Println("Subscription entry will be removed:")
		fmt.Println(SubscriptiontoString(entry))
		fmt.Println("Continue?[Y/n]")

		var ans string = ""
		fmt.Scanf("%s", &ans)
		if ans == "" || strings.ToLower(ans) == "y" {
			RemoveSubscription(id)
			fmt.Println("Subscription removed.")
		} else {
			fmt.Println("Nothing is done.")
		}
	} else {
		fmt.Printf("Subscription with id %d does not exist!", id)
	}
	return nil
}

func U() error {
	var id int64
	fmt.Print("Subscription id you want to update: ")

	_, err := fmt.Scanf("%d", &id)
	if err != nil {
		return err
	}

	var entry Subscription
	var exists bool
	if entry, exists = GetSubscription(id); !exists {
		fmt.Printf("Subscription with id %d does not exist!", id)
		return nil
	}

	fmt.Println("You are updating: ", SubscriptiontoString(entry))
	fmt.Print("What do you want to update? (n for name, e for expense, c for currency): ")

	var ans string
	_, err = fmt.Scanf("%s", &ans)
	if err != nil {
		return err
	}

	switch ans {
	case "n":
		fmt.Print("New name: ")
		var newName string
		_, err := fmt.Scanf("%s", &newName)
		if err != nil {
			return err
		}
		UpdateSubscriptionName(id, newName)
	case "e":
		fmt.Print("New expense: ")
		var newExpense float64
		_, err := fmt.Scanf("%f", &newExpense)
		if err != nil {
			return err
		}
		UpdateSubscriptionExpense(id, newExpense)
	case "c":
		fmt.Print("New currency: ")
		var newCurrencyStr string
		_, err := fmt.Scanf("%s", &newCurrencyStr)
		if err != nil {
			return err
		}
		newCurrency, err := NewCurrency(newCurrencyStr)
		if err != nil {
			return err
		}
		UpdateSubscriptionCurrency(id, newCurrency)
	default:
		return fmt.Errorf("Invalid field to update. Got: %s", ans)
	}

	return nil
}
