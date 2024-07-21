package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/hi6okuni/writing-interpreter-in-go/repl"
)

func main() {

	// CPUプロファイリングの開始
	// cpuProfile, err := os.Create("cpu_profile.pprof")
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "could not create CPU profile: %v\n", err)
	// 	return
	// }
	// if err := pprof.StartCPUProfile(cpuProfile); err != nil {
	// 	fmt.Fprintf(os.Stderr, "could not start CPU profile: %v\n", err)
	// 	return
	// }
	// defer pprof.StopCPUProfile()

	// // メモリプロファイリングの設定
	// memProfile, err := os.Create("mem_profile.pprof")
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "could not create memory profile: %v\n", err)
	// 	return
	// }
	// defer memProfile.Close()
	// runtime.GC() // メモリプロファイルを取る前にGCを実行
	// if err := pprof.WriteHeapProfile(memProfile); err != nil {
	// 	fmt.Fprintf(os.Stderr, "could not write memory profile: %v\n", err)
	// 	return
	// }

	// // HTTP serverの起動（pprofのデバッグ情報を提供）
	// go func() {
	// 	fmt.Println(http.ListenAndServe("localhost:6060", nil))
	// }()

	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n",
		user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
