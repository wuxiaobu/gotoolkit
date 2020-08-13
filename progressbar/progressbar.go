package progressbar

import (
	"fmt"
	"github.com/wuxiaobu/gotoolkit/color"
	"os"
	"strconv"
	"strings"
	"time"
)

type ProgressBar struct {
	BarType string
	count int
	work chan int
	color color.Color
	startTime time.Time
}

func NewPb() *ProgressBar {
	return &ProgressBar{
		work: make(chan int),
		startTime: time.Now(),
	}
}

func (pb *ProgressBar) Bar(describe, BarType string, count int) {
	go pb.creatWorker(describe, BarType, count)
}

func (pb *ProgressBar) creatWorker(describe, BarType string, count int) {
	_, _ = fmt.Fprintln(os.Stdout, describe)
	pb.worker(BarType, count)
	_, _ = fmt.Fprintf(os.Stdout, "\n")
}

func (pb *ProgressBar) worker(BarType string, count int) {
	for i := 0; i <= count; i++ {
		complete := int(i * 100 / count)

		restStr := strings.Repeat("-", 100 - complete)
		completeStr := strings.Repeat(BarType, complete)

		completeStr = pb.color.Green(completeStr)

		//_, _ = fmt.Fprint(os.Stdout, "\r[")
		//_, _ = fmt.Fprintf(os.Stdout, "%s%s]", completeStr, restStr)

		timeAfter := time.Now().Sub(pb.startTime).Seconds()
		timeAfter, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", timeAfter), 64)
		//_, _ = fmt.Fprintf(os.Stdout, "\t%d %%", complete)
		//_, _ = fmt.Fprintf(os.Stdout, "\t耗时%vs", timeAfter)

		out := fmt.Sprintf("\r[%s%s\t%d %%\t耗时%vs", completeStr, restStr, complete, timeAfter)
		_, _ = fmt.Fprintf(os.Stdout, "%s", out)
		pb.work <- i
	}
}

func (pb *ProgressBar) Complete() {
	<-pb.work
}
