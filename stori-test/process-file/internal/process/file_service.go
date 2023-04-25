package process

import (
	"bytes"
	"encoding/csv"
	"io"
	"log"
	"process-file/internal"
	"process-file/internal/storage/s3"
	"strconv"
	"strings"
)

type FileService struct {
	S3Repository *s3.Repository
	FileName     string
}

func NewFileService(s3Repository *s3.Repository, fileName string) *FileService {
	return &FileService{
		S3Repository: s3Repository,
		FileName:     fileName,
	}
}

func (f *FileService) read(fileToGet internal.File) ([]byte, error) {
	file, err := f.S3Repository.GetFileInBytes(fileToGet)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (f *FileService) GetSummary(fileToGet internal.File) (internal.Summary, error) {
	var summary internal.Summary
	file, err := f.read(fileToGet)
	if err != nil {
		return internal.Summary{}, err
	}
	r2 := bytes.NewReader(file)
	r := csv.NewReader(r2)
	_, err = r.Read()
	if err == io.EOF {
		return internal.Summary{}, err
	}
	summary.TransactionsByMonth = make([]int32, 12)
	lowestMonth := int64(12)
	highestMonth := int64(1)
	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		}
		averageDebitAmount, averageCreditAmount, err := getSumatory(line[2], summary.AverageDebitAmount, summary.AverageCreditAmount)
		if err != nil {
			return internal.Summary{}, err
		}
		month, err := getMonth(line[1])
		if err != nil {
			return internal.Summary{}, err
		}
		if month < lowestMonth {
			lowestMonth = month
		}
		if month > highestMonth {
			highestMonth = month
		}
		summary.AverageDebitAmount = averageDebitAmount
		summary.AverageCreditAmount = averageCreditAmount
		summary.TransactionsByMonth[month-1] += 1
	}
	summary.TotalBalance = summary.AverageCreditAmount + summary.AverageDebitAmount
	summary.AverageDebitAmount /= float64(highestMonth - lowestMonth + 1)
	summary.AverageCreditAmount /= float64(highestMonth - lowestMonth + 1)
	totalMonth := 0

	log.Println(lowestMonth, highestMonth, totalMonth, summary)
	return summary, nil
}

func getSumatory(line string, averageDebitAmount, averageCreditAmount float64) (float64, float64, error) {
	amount, err := strconv.ParseFloat(line, 64)
	if err != nil {
		return 0.0, 0.0, err
	}
	if amount < 0 {
		averageDebitAmount += amount
	} else {
		averageCreditAmount += amount
	}
	return averageDebitAmount, averageCreditAmount, nil
}

func getMonth(monthStr string) (int64, error) {
	date := strings.Split(monthStr, "/")[0]
	month, err := strconv.ParseInt(date, 10, 64)
	if err != nil {
		return 0, err
	}
	return month, nil
}
