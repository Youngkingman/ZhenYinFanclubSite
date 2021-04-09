package datatrans

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

func RecordData(filename string, feasible [][]int, cost [][]string, step string, tract [][2]int, id int) {
	newFileName := "D:\\gotest\\src\\source\\newfile.csv"
	//nfs, err := os.Create(newFileName)

	nfs, err := os.OpenFile(newFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("can not create file, err is %+v", err)
	}
	defer nfs.Close()
	nfs.Seek(0, io.SeekEnd)

	w := csv.NewWriter(nfs)
	w.Comma = ','
	w.UseCRLF = true
	row := []string{"1", "2", "3", "4", "5,6"}
	err = w.Write(row)
	if err != nil {
		log.Fatalf("can not write, err is %+v", err)
	}
	w.Flush()

	var newContent [][]string
	newContent = append(newContent, []string{"1", "2", "3", "4", "5", "6"})
	newContent = append(newContent, []string{"11", "12", "13", "14", "15", "16"})
	newContent = append(newContent, []string{"21", "22", "23", "24", "25", "26"})
	w.WriteAll(newContent)
}

func OutputMat(filename string, feasible [][]int, cost [][]int, id int) {
	//save feasible
	f1, err := os.Create(filename + "_feasible_" + strconv.Itoa(id) + ".csv")
	if err != nil {
		panic(err)
	}
	f1.WriteString("\xEF\xBB\xBF")
	w1 := csv.NewWriter(f1)
	w1.WriteAll(formatStringmat(feasible))
	w1.Flush()
	f1.Close()

	//save cost
	f3, err := os.Create(filename + "_cost_" + strconv.Itoa(id) + ".csv")
	if err != nil {
		panic(err)
	}
	f3.WriteString("\xEF\xBB\xBF")
	w3 := csv.NewWriter(f3)
	w3.WriteAll(formatStringmat(cost))
	w3.Flush()
	f3.Close()

}

func OutputTract(filename string, tract [][2]int, id int) {
	//save tract
	f2, err := os.Create(filename + "_tract_" + strconv.Itoa(id) + ".csv")
	if err != nil {
		panic(err)
	}
	f2.WriteString("\xEF\xBB\xBF")
	w2 := csv.NewWriter(f2)
	w2.WriteAll(formatTrace(tract))
	w2.Flush()
	f2.Close()
}

func formatStringmat(InputMat [][]int) (OutputMat [][]string) {
	for _, arr := range InputMat {
		str := make([]string, len(arr))
		for i, v := range arr {
			str[i] = strconv.Itoa(v)
		}
		OutputMat = append(OutputMat, str)
	}
	return
}

func formatTrace(InputMat [][2]int) (OutputMat [][]string) {
	for _, arr := range InputMat {
		str := []string{strconv.Itoa(arr[0]), strconv.Itoa(arr[1])}
		OutputMat = append(OutputMat, str)
	}
	return
}
