package facade

import (
	"fmt"
	"reflect"
)

/*
The facade pattern is fairly simple. When you want to obfuscate or keep from
writing repetitive code for functions that a user might take, or a program
might invoke.

An example might be a financial transaction, where the user wants to make a
purchase.  The transaction may take several steps:
		- Verify User Account Id
		- Verify Security Code
		- Check if funcds are available
		- Send funds
		- Subtract funds from account
		...

We do not want to have to call each of these steps in our code each time our
application has a point where a user can purchase something.  So we hide the
subsequent steps behind a facade, which we can easily call.
*/

type Transaction struct {
	cost         float32
	verified     bool
	userId       string
	securityCode []string
	error        []string
}

type Account struct {
	funds        float32
	userId       string
	securityCode []string
}

// we have some accounts
var (
	user1 = Account{200.00, "f09h3209hf02h1k2lj23F", []string{"1", "0", "9", "1"}}
	user2 = Account{200.00, "093h09h1029h8392g9g32", []string{"3", "3", "3", "2"}}
)

// instead of writing the following every time we have a new transaction
func doStuff() {

	// some business logic for user login or something

	t := Transaction{
		10.15,
		false,
		"f09h3209hf02h1k2lj23F",
		[]string{"1", "0", "9", "1"},
		[]string{},
	}

	if user1.userId != t.userId {
		t.verified = false
		t.error = append(t.error, "Incorrect user id")
	}

	if reflect.DeepEqual(user1.securityCode, t.securityCode) == false {
		t.verified = false
		t.error = append(t.error, "Invalid security code")
	}

	if user1.funds < t.cost {
		t.verified = false
		t.error = append(t.error, "Insufficient funds")
	}

	user1.funds = user1.funds - t.cost

	t.verified = true

	// further logic for other functionality

}

// we can write it in a facade such as

func (t *Transaction) verify(user Account) (bool, *Transaction) {
	if t.userId != user.userId {
		t.error = append(t.error, "Invalid User Id")
		return false, t
	}

	if reflect.DeepEqual(t.securityCode, user.securityCode) == false {
		t.error = append(t.error, "Incorrect security code")
		return false, t
	}

	if user.funds < t.cost {
		t.error = append(t.error, "Insufficient funds")
		return false, t
	}

	t.verified = true
	return true, t
}

func (u *Account) sendFunds(t *Transaction) *Account {
	u.funds = u.funds - t.cost
	return u
}

func doStuffRight() {

	// some business logic for user login or something

	// here is where we can have a trasnaction
	t := Transaction{
		10.15,
		false,
		"f09h3209hf02h1k2lj23F",
		[]string{"1", "0", "9", "1"},
		[]string{},
	}

	if ok, t := t.verify(user1); ok {
		user1.sendFunds(t)
	} else {
		for _, err := range t.error {
			fmt.Printf("Transaction Error: %s", err)
		}
	}

	// further logic for other functionality

}

/*
this allows us to write far less, and easier to maintain code at the transaction
point
*/
