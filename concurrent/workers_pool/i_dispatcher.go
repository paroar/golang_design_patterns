package main

type IDispatcher interface {
	LaunchWorker(w IWorkerLauncher)
	MakeRequest(Request)
	Stop()
}
