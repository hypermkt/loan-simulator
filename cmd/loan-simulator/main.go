package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
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

func main() {
	y, ir, a := parseArgs()
	p := toParams(y, ir, a)

	loanTables := calcLoanTables(p)

	printParams(p)
	fmt.Println("---")
	printLoanTables(loanTables)
}

func parseArgs() (int, float64, int) {
	y := flag.Int("y", 35, "返済期間(年)")
	ir := flag.Float64("i", 1, "金利(%)")
	a := flag.Int("a", 0, "借入金額(万円)")
	flag.Parse()

	return *y, *ir, *a
}

func toParams(y int, ir float64, a int) Params {
	return Params{
		Year:                y,
		Months:              calcMonths(y),
		InterestRate:        ir,
		MonthlyInterestRate: ir / 100 / 12,
		AmountMan:           a,
		Amount:              a * 10000,
		CurrentBalance:      a * 10000,
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

func printLoanTables(lts [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"回数", "年/月", "返済額", "元金", "利息", "借入残高"})
	for _, lt := range lts {
		table.Append(lt)
	}
	table.Render() // Send output
}

func calcMonths(y int) int {
	return y * 12
}

func calcRepaidAmount(p Params) int {
	r := math.Abs(float64(p.Amount) * float64(p.MonthlyInterestRate) * math.Pow((1+p.MonthlyInterestRate), float64(p.Months)) / (math.Pow((1+p.MonthlyInterestRate), float64(p.Months)) - 1))

	return int(r)
}

func calcInterest(p Params) int {
	r := (float64(p.CurrentBalance) * p.MonthlyInterestRate * 1)

	return int(r)
}

func calcLoanTables(p Params) [][]string {
	t := time.Now()
	rows := [][]string{}

	for i := 0; i < p.Months; i++ {
		ra := calcRepaidAmount(p)
		int := calcInterest(p)
		pa := ra - int
		b := p.CurrentBalance - pa
		p.CurrentBalance = b

		if b < ra {
			ra = ra + b
			b = 0
		}

		row := []string{
			strconv.Itoa(i + 1),
			t.Format("2006-01"),
			strconv.Itoa(ra),
			strconv.Itoa(pa),
			strconv.Itoa(int),
			strconv.Itoa(b),
		}
		rows = append(rows, row)

		t = t.AddDate(0, 1, 0)
	}

	return rows
}
