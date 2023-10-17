package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/msft/bank/core"
)

type CustomAccount struct {
	*core.Account
}

// 重写 Account 的 Statement 方法
func (c *CustomAccount) Statement() string {
	json, err := json.Marshal(c)
	if err != nil {
		return err.Error()
	}

	return string(json)
}

var accounts = map[int32]*CustomAccount{}

func main() {
	accounts[1001] = &CustomAccount{
		Account: &core.Account{
			Customer: core.Customer{
				Name:    "John",
				Address: "Los Angeles, California",
				Phone:   "(213) 555 0147",
			},
			Number:  1001,
			Balance: 0,
		},
	}
	accounts[1002] = &CustomAccount{
		Account: &core.Account{
			Customer: core.Customer{
				Name:    "Mark",
				Address: "Irvine, California",
				Phone:   "(949) 555 0198",
			},
			Number:  1002,
			Balance: 0,
		},
	}

	http.HandleFunc("/statement", statement)
	http.HandleFunc("/deposit", deposit)
	http.HandleFunc("/withdraw", withdraw)
	http.HandleFunc("/transfer", transfer)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func statement(w http.ResponseWriter, req *http.Request) {
	numberOfQs := req.URL.Query().Get("number")

	if numberOfQs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}

	if number, err := strconv.ParseInt(numberOfQs, 10, 32); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else {
		account, ok := accounts[int32(number)]
		if !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", number)
		} else {
			json.NewEncoder(w).Encode(core.Statement(account))
		}
	}
}

func deposit(w http.ResponseWriter, req *http.Request) {
	numberOfQs := req.URL.Query().Get("number")
	amountOfQs := req.URL.Query().Get("amount")

	if numberOfQs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}

	if number, err := strconv.ParseInt(numberOfQs, 10, 32); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else if amount, err := strconv.ParseFloat(amountOfQs, 64); err != nil {
		fmt.Fprintf(w, "Invalid amount number!")
	} else {
		account, ok := accounts[int32(number)]
		if !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", number)
		} else {
			err := account.Deposit(amount)
			if err != nil {
				fmt.Fprintf(w, "%v", err)
			} else {
				fmt.Fprintf(w, account.Statement())
			}
		}
	}
}

func withdraw(w http.ResponseWriter, req *http.Request) {
	numberOfQs := req.URL.Query().Get("number")
	amountOfQs := req.URL.Query().Get("amount")

	if numberOfQs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}

	if number, err := strconv.ParseInt(numberOfQs, 10, 32); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else if amount, err := strconv.ParseFloat(amountOfQs, 64); err != nil {
		fmt.Fprintf(w, "Invalid amount number!")
	} else {
		account, ok := accounts[int32(number)]
		if !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", number)
		} else {
			err := account.Withdraw(amount)
			if err != nil {
				fmt.Fprintf(w, "%v", err)
			} else {
				fmt.Fprintf(w, account.Statement())
			}
		}
	}
}

func transfer(w http.ResponseWriter, req *http.Request) {
	numberOfQs := req.URL.Query().Get("number")
	amountOfQs := req.URL.Query().Get("amount")
	destOfQs := req.URL.Query().Get("dest")

	if numberOfQs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}

	if number, err := strconv.ParseInt(numberOfQs, 10, 32); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else if amount, err := strconv.ParseFloat(amountOfQs, 64); err != nil {
		fmt.Fprintf(w, "Invalid amount number!")
	} else if dest, err := strconv.ParseInt(destOfQs, 10, 32); err != nil {
		fmt.Fprintf(w, "Invalid dest account number!")
	} else {

		if accountA, ok := accounts[int32(number)]; !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", number)
		} else if accountB, ok := accounts[int32(dest)]; !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", dest)
		} else {
			err := accountA.Transfer(amount, accountB.Account)
			if err != nil {
				fmt.Fprintf(w, "%v", err)
			} else {
				fmt.Fprintf(w, accountA.Statement())
			}
		}
	}
}
