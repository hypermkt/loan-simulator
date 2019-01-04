package main

import (
	"flag"
	"fmt"
	"math"
	"strconv"
	"time"
)

type Params struct {
	Year                int     // 返済期間(年)
	Months              int     // 回数
	InterestRate        float64 // 金利(年利)
	MonthlyInterestRate float64 // 金利(月利)
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
	y, ir, a := parseArgs()
	p := toParams(y, ir, a)

	var loanTables LoanTables

	t := time.Now()
	for i := 0; i < p.Months; i++ {
		repaidAmount := calcRepaidAmount(p)
		interest := calcInterest(p)
		princepalAmount := repaidAmount - interest
		balance := p.CurrentBalance - princepalAmount
		p.CurrentBalance = balance

		if balance < repaidAmount {
			repaidAmount = repaidAmount + balance
			balance = 0
		}

		loanTables = append(loanTables, LoanTable{
			Count:           i + 1,
			Date:            t.Format("2006-01"),
			RepaidAmount:    repaidAmount,
			PrincipalAmount: princepalAmount,
			Interest:        interest,
			Balance:         balance,
		})
		t = t.AddDate(0, 1, 0)
	}
	printParams(p)
	fmt.Println("---")
	printLoanTables(loanTables)
}

func calcRepaidAmount(p Params) int {
	r := math.Abs(float64(p.Amount) * float64(p.MonthlyInterestRate) * math.Pow((1+p.MonthlyInterestRate), float64(p.Months)) / (math.Pow((1+p.MonthlyInterestRate), float64(p.Months)) - 1))

	return int(r)
}

func calcInterest(p Params) int {
	r := (float64(p.CurrentBalance) * p.MonthlyInterestRate * 1)

	return int(r)
}

func parseArgs() (string, string, string) {
	y := flag.String("y", "", "返済期間(年)")
	ir := flag.String("i", "", "金利(%)")
	a := flag.String("a", "", "借入金額(万円)")
	flag.Parse()

	return *y, *ir, *a
}

func toParams(y, ir, a string) Params {
	iy, _ := strconv.Atoi(y)
	iir, _ := strconv.ParseFloat(ir, 64)
	iam, _ := strconv.Atoi(a)
	ia, _ := strconv.Atoi(a)

	return Params{
		Year:                iy,
		Months:              calcMonths(iy),
		InterestRate:        iir,
		MonthlyInterestRate: iir / 100 / 12,
		AmountMan:           iam,
		Amount:              ia * 10000,
		CurrentBalance:      ia * 10000,
	}
}

func printParams(p Params) {
	fmt.Printf("返済期間: %v 年\n", p.Year)
	fmt.Printf("返済期間: %v ヶ月\n", p.Months)
	fmt.Printf("金利: %v ％\n", p.InterestRate)
	fmt.Printf("月利: %v\n", p.MonthlyInterestRate)
	fmt.Printf("借入金額: %v 万円\n", p.AmountMan)
	fmt.Printf("借入金額: %v 円\n", p.Amount)
}

func printLoanTables(lts LoanTables) {
	fmt.Printf("回数 年/月 返済額 元金 利息 借入残高\n")
	for _, lt := range lts {
		fmt.Printf("%3v  %v  %v  %v  %v   %v\n",
			lt.Count,
			lt.Date,
			lt.RepaidAmount,
			lt.PrincipalAmount,
			lt.Interest,
			lt.Balance,
		)
	}
}

func calcMonths(y int) int {
	return y * 12
}
