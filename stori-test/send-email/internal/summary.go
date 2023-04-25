package internal

type Summary struct {
	TotalBalance        float64
	AverageDebitAmount  float64
	AverageCreditAmount float64
	TransactionsByMonth []int32
}

type SummaryTemplate struct {
	TotalBalance        float64
	AverageDebitAmount  float64
	AverageCreditAmount float64
	TransactionsByMonth []TransactionByMonth
}
type TransactionByMonth struct {
	NameMonth    string
	Transactions int32
}

var Months = map[int]string{
	0:  "Jan",
	1:  "Feb",
	2:  "Mar",
	3:  "Apr",
	4:  "May",
	5:  "Jun",
	6:  "Jul",
	7:  "Aug",
	8:  "Sep",
	9:  "Oct",
	10: "Nov",
	11: "Dic",
}
