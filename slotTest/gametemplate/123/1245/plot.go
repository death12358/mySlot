package main

import (
	"fmt"
	"os"
	"sync"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

var totalWin int64
var wg sync.WaitGroup

func main() {
	// 生成 100x10000 的矩陣數據
	matrix := generateMatrix(100, 5000)

	// 繪製折線圖並保存為 PNG 文件
	err := plotMatrix(matrix, "line_plot.png")
	if err != nil {
		fmt.Println("Error plotting matrix:", err)
		os.Exit(1)
	}
	fmt.Println("totalWin:", totalWin)
	fmt.Println("Line plot saved as line_plot.png")
}

// generateMatrix 生成 100x10000 的矩陣數據
func generateMatrix(rows, cols int) [][]int64 {
	matrix := make([][]int64, rows)
	for i := range matrix {
		wg.Add(1)
		// 隨機生成 0-100 的數值
		go func() {
			arr := getOneProcess(cols, i)
			pointLock.Lock()
			matrix = append(matrix, arr)
			pointLock.Unlock()
		}()
	}
	wg.Wait()

	return matrix
}

func getOneProcess(cols int, i int) []int64 {
	balance := int64(bet) * 500
	arr := make([]int64, cols)
	for j := 0; j < cols; j++ {
		if j%1000 == 0 {
			fmt.Println(i, j)
		}
		win := getPoint(int32(i))
		balance += win - int64(bet)
		arr[j] = balance

		pointLock.Lock()
		totalWin += win
		pointLock.Unlock()
	}
	wg.Done()
	return arr
}

// plotMatrix 繪製折線圖
func plotMatrix(matrix [][]int64, filename string) error {
	p := plot.New()

	// 設置標題和坐標軸標籤
	p.Title.Text = "Line Plot"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	// 遍歷矩陣的每一行並將其轉換為一條折線
	for i := 0; i < len(matrix); i++ {
		lineData := make(plotter.XYs, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			lineData[j].X = float64(j)
			lineData[j].Y = float64(matrix[i][j])
		}

		// 將每條折線添加到 plot 中
		line, err := plotter.NewLine(lineData)
		if err != nil {
			return err
		}

		p.Add(line)
	}

	// 保存為 PNG 圖片
	if err := p.Save(10*vg.Inch, 6*vg.Inch, filename); err != nil {
		return err
	}

	return nil
}
