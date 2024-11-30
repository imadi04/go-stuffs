package prices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxincludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) LoadData() {
	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Println("File not found!")
		fmt.Println(err)
	}
	// NewScanner() is used to create a Scanner, which is a convenient tool for reading input from a variety of sources
	//(like files, standard input, or strings) token by token.
	scanner := bufio.NewScanner(file)
	var lines []string // to store scanned prices.
	//scanner.Scan() // read one line at a time from file.
	for scanner.Scan() {
		lines = append(lines, scanner.Text()) // text() to get the content of scanned line by scan()
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println("Reading the file content failed!!")
		fmt.Println(err)
		file.Close()
		return
	}
	prices := make([]float64, len(lines))

	for lineIndex, line := range lines {
		floatPrice, err := strconv.ParseFloat(line, 64) // func to convert string to float ; 64 here means float64
		if err != nil {
			fmt.Println("Converting price to float failed ")
			fmt.Println(err)
			file.Close()
			return
		}
		prices[lineIndex] = floatPrice
	}
	job.InputPrices = prices

}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()
	result := make(map[string]string) // as we declared map TaxincludedPrices
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}
	fmt.Println(result)
}

//constructor function

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:     taxRate,
		InputPrices: []float64{10, 20, 30},
	}
}
