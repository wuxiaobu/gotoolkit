package progressbar

import (
	"testing"
	"time"
)

func TestProgressBar(t *testing.T) {
	Pb := NewPb()
	Pb.Bar("测试bar", "◆", 100)
	for i := 0; i < 100; i++ {
		Pb.Complete()
		time.Sleep(time.Second / 10)
	}
}
