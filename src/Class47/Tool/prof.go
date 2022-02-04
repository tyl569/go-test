package Tool

import (
	"log"
	"math/rand"
	"os"
	"time"

	"runtime/pprof"
)

const (
	ROW = 3
	COL = 3
)

func main() {
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("create cpu.prof failed", err)
	}

	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("start cpu prof failed", err)
	}
	defer pprof.StopCPUProfile()

	x := [ROW][COL]int{}
	fillMatrix(&x)
	calculate(&x)

	f1, err := os.Create("mem.prof")
	if err != nil {
		log.Fatal("create mem.prof failed", err)
	}

	//runtime.GC()
	if err := pprof.WriteHeapProfile(f1); err != nil {
		log.Fatal("cannot write memory profile", err)
	}
	f1.Close()

	f2, err := os.Create("goroutine.prof")
	if err != nil {
		log.Fatal("create goroutine.prof failed", err)
	}
	if gProf := pprof.Lookup("goroutine"); gProf != nil {
		log.Fatal("cannot write goroutine prof", err)
	} else {
		gProf.WriteTo(f2, 0)
	}
	f2.Close()

}

func fillMatrix(m *[ROW][COL]int) {
	s := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < ROW; i++ {
		for j := 0; j < COL; j++ {
			m[i][j] = s.Intn(100000)
		}
	}
}

func calculate(m *[ROW][COL]int) {
	for i := 0; i < ROW; i++ {
		tmp := 0
		for j := 0; j < COL; j++ {
			tmp += m[i][j]
		}
	}
}
