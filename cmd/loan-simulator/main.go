package main

import (
	"flag"
	"fmt"
	"math"
	"strconv"
	"time"
)

type Params struct {
	Year                int     //
	Months              int     // 回数
	InterestRate        float64 // 金利(年利)
	MonthlyInterestRate float64 // 月利
	AmountMan           int     // 金額(万円)
	Amount              int     // 金額
	CurrentBalance      int     // 残高
}

type LoanTable struct {
	Count           int    // 回数
	Date            string // 年/月
	RepaidAmount    int    // 返済額
	PrincipalAmount int    // 元金
	Interest        int    // 利息
	Balance         int    // 残高
}

type LoanTables []LoanTable

func main() {
	year, interestRate, amount := parseArgs()
	params := toParams(year, interestRate, amount)

	var loanTables LoanTables

	t := time.Now()
	for i := 0; i < params.Months; i++ {
		repaidAmount := calcRepaidAmount(params)
		interest := calcInterest(params)
		princepalAmount := repaidAmount - interest
		balance := params.CurrentBalance - princepalAmount
		params.CurrentBalance = balance

		if balance < repaidAmount {
			repaidAmount = repaidAmount + balance
			balance = 0
		}

		loanTables = append(loanTables, LoanTable{
			Count:           i + 1,
			Date:            t.Format("2006-01"),
			RepaidAmount:    repaidAmount,    // 返済金額
			PrincipalAmount: princepalAmount, // 元金
			Interest:        interest,        // 利息
			Balance:         balance,         // 残高
		})
		t = t.AddDate(0, 1, 0)
	}
	printParams(params)
	fmt.Println("---")
	printLoanTables(loanTables)
}

func calcRepaidAmount(params Params) int {
	r := math.Abs(float64(params.Amount) * float64(params.MonthlyInterestRate) * math.Pow((1+params.MonthlyInterestRate), float64(params.Months)) / (math.Pow((1+params.MonthlyInterestRate), float64(params.Months)) - 1))

	return int(r)
}

func calcInterest(params Params) int {
	r := (float64(params.CurrentBalance) * params.MonthlyInterestRate * 1)

	return int(r)
}

func parseArgs() (string, string, string) {
	year := flag.String("y", "", "返済期間(年)")
	interestRate := flag.String("i", "", "金利(%)")
	amount := flag.String("a", "", "借入金額(万円)")
	flag.Parse()

	return *year, *interestRate, *amount
}

func toParams(year, interestRate, amount string) Params {
	iYear, _ := strconv.Atoi(year)
	iInterestRate, _ := strconv.ParseFloat(interestRate, 64)
	iAmountMan, _ := strconv.Atoi(amount)
	iAmount, _ := strconv.Atoi(amount)

	return Params{
		Year:                iYear,
		Months:              calcMonths(iYear),
		InterestRate:        iInterestRate,
		MonthlyInterestRate: iInterestRate / 100 / 12,
		AmountMan:           iAmountMan,
		Amount:              iAmount * 10000,
		CurrentBalance:      iAmount * 10000,
	}
}

func printParams(params Params) {
	fmt.Printf("返済期間: %v 年\n", params.Year)
	fmt.Printf("返済期間: %v ヶ月\n", params.Months)
	fmt.Printf("金利: %v ％\n", params.InterestRate)
	fmt.Printf("月利: %v\n", params.MonthlyInterestRate)
	fmt.Printf("借入金額: %v 万円\n", params.AmountMan)
	fmt.Printf("借入金額: %v 円\n", params.Amount)
}

func printLoanTables(loanTables LoanTables) {
	fmt.Printf("回数 年/月 返済額 元金 利息 借入残高\n")
	for _, loan := range loanTables {
		fmt.Printf("%3v  %v  %v  %v  %v   %v\n",
			loan.Count,
			loan.Date,
			loan.RepaidAmount,
			loan.PrincipalAmount,
			loan.Interest,
			loan.Balance,
		)
	}
}

func calcMonths(year int) int {
	return year * 12
}
