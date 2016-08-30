package Bll
import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)
func TeatA(){
	file, err := os.Open("data/0600000.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.Read()//丢弃表头
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("%T\n",record)
		fmt.Println(record," end!") // record has the type []string
	}
}