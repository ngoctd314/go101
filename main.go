package main

import (
	"fmt"
	"io/ioutil"
	_ "net/http/pprof"
	"os"
	"strconv"
	"strings"

	"github.com/ngoctd314/go101/go-code/golog"
)

func fn() {
	if _, err := os.Stat("/home/idev/code/go101/localpath/" + "/" + "2022"); os.IsNotExist(err) {
		err = os.MkdirAll("/home/idev/code/go101/localpath/"+"/"+"2022", os.ModePerm)
		if err != nil {
			golog.Error(golog.Server, err.Error())
		}
	}
	// curTime := time.Now().Local()
	// curDay := curTime.Format("2006-01-02")
	// localPath := "/home/idev/code/go101/localpath/" + "/" + curDay
	// nfsPath := "/home/idev/code/go101/nfspath/" + "/" + curDay
	// files, err := ioutil.ReadDir(localPath)
	// if err != nil {
	// 	golog.Error(golog.Server, err.Error())
	// }
	// if _, err := os.Stat(nfsPath); os.IsNotExist(err) {
	// 	_ = os.Mkdir(nfsPath, os.ModePerm)
	// }

	// for _, file := range files {
	// 	err := os.Rename(localPath+"/"+"/"+file.Name(), nfsPath+"/"+file.Name())
	// 	if err != nil {
	// 		golog.Error(golog.Server, err.Error())
	// 	}
	// }
}

func getCPUSample() (idle, total uint64) {
	contents, err := ioutil.ReadFile("/proc/stat")
	if err != nil {
		return
	}
	lines := strings.Split(string(contents), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if fields[0] == "cpu" {
			numFields := len(fields)
			for i := 1; i < numFields; i++ {
				val, err := strconv.ParseUint(fields[i], 10, 64)
				if err != nil {
					fmt.Println("Error: ", i, fields[i], err)
				}
				total += val // tally up all the numbers to get total ticks
				if i == 4 {  // idle is the 5th field in the cpu line
					idle = val
				}
			}
			return
		}
	}
	return
}

func main() {
	idle0, total0 := getCPUSample()
	for i := 0; i < 10e9; i++ {
	}
	idle1, total1 := getCPUSample()

	idleTicks := float64(idle1 - idle0)
	totalTicks := float64(total1 - total0)
	cpuUsage := 100 * (totalTicks - idleTicks) / totalTicks

	fmt.Printf("CPU usage is %f%% [busy: %f, total: %f]\n", cpuUsage, totalTicks-idleTicks, totalTicks)
}
