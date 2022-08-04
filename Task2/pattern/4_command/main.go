package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

func main() {
	fmt.Println(strings.Repeat("-", 65))

	// Создаем 2 аккаунта над которыми будут применяться команды
	rasul := NewAccount("Rasul", 0)
	ildar := NewAccount("Ildar", 0)

	cmdManager := CmdManager{}
	if err := cmdManager.
		Add(NewDeposit(820, rasul)).
		Add(NewWithdraw(139, rasul)).
		Add(NewDeposit(321, ildar)).
		Add(NewDeposit(132.3, rasul)).
		Add(NewWithdraw(192, ildar)).
		Run(); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(strings.Repeat("-", 65))

	fmt.Println(*rasul)
	fmt.Println(*ildar)
}

type Account struct {
	name  string
	money float64
}

func NewAccount(name string, money float64) *Account {
	return &Account{name, money}
}

type aCommand struct {
	account        *Account
	isCompleted    bool
	moneyToOperate float64
}

func (a *aCommand) IsCompleted() bool {
	return a.isCompleted
}

func newACommandAccount(account *Account, money float64) *aCommand {
	return &aCommand{account, false, money}
}

type Deposit struct {
	aCommand
}

func NewDeposit(toDeposit float64, account *Account) *Deposit {
	return &Deposit{*newACommandAccount(account, toDeposit)}
}

func (d *Deposit) Execute() error {
	if d.isCompleted {
		return fmt.Errorf("deposit +$%f was completed", d.moneyToOperate)
	}
	d.isCompleted = true
	d.account.money += d.moneyToOperate
	fmt.Printf("%s: put money in account +$%f, ", d.account.name, d.moneyToOperate)
	fmt.Printf("new balance $%f\n", d.account.money)
	return nil
}

type WithDraw struct {
	aCommand
}

func NewWithdraw(toWithDraw float64, account *Account) *WithDraw {
	return &WithDraw{*newACommandAccount(account, toWithDraw)}
}

func (w WithDraw) Execute() error {
	if w.isCompleted {
		return fmt.Errorf("withdraw -%f was completed", w.moneyToOperate)
	} else if w.account.money < w.moneyToOperate {
		return errors.New(w.account.name + ": haven't enough money not withdraw")
	}
	w.isCompleted = true
	w.account.money -= w.moneyToOperate
	fmt.Printf("%s: withdrawed from card -$%f, ", w.account.name, w.moneyToOperate)
	fmt.Printf("new balance $%f\n", w.account.money)
	return nil
}

type Executable interface {
	Execute() error
}

type CmdManager struct {
	commands []Executable
}

func (e *CmdManager) Add(execute Executable) *CmdManager {
	e.commands = append(e.commands, execute)
	return e
}

func (e *CmdManager) Run() (ok error) {
	for _, command := range e.commands {
		if ok = command.Execute(); ok != nil {
			return ok
		}
	}
	return nil
}
