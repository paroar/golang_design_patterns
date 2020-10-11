package main

type IWorkerLauncher interface {
	LaunchWorker(in chan Request)
}
