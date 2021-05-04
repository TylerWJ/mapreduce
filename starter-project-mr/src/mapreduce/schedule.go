package mapreduce

// schedule starts and waits for all tasks in the given phase (Map or Reduce).
func (mr *Master) schedule(phase jobPhase) {
	var ntasks int
	var nios int
	switch phase {
	case mapPhase:
		ntasks = len(mr.files) // --> nMap
		nios = mr.nReduce      // --> nReduce
	case reducePhase:
		ntasks = mr.nReduce  // --> nReduce
		nios = len(mr.files) // --> nMap
	}

	debug("Schedule: %v %v tasks (%d I/Os)\n", ntasks, phase, nios)

	var task DoTaskArgs

	for i := 0; i < ntasks; i++ {
		wName := <-mr.registerChannel

		task.JobName = mr.jobName
		task.File = mr.files[i]
		task.Phase = phase
		task.TaskNumber = i
		task.NumOtherPhase = nios

		for !call(wName, "Worker.DoTask", task, nil) {
			wName = <-mr.registerChannel
		}

		go func() {
			mr.registerChannel <- wName
		}()

	}

	// How I think this is gunna work:

	// Register each worker
	// DoTask

	// All ntasks tasks have to be scheduled on workers, and only once all of
	// them have been completed successfully should the function return.
	// Remember that workers may fail, and that any given worker may finish
	// multiple tasks.
	//
	// TODO:

	// hand out the map and reduce tasks to workers
	// return only when tasks have finished
	// use RPC and channels to simulate a distributed environment
	// master -> RPC -> workers
	// worker code: worker.go
	// code to deal with RPC messages: common_rpc.go

	// Three roles:
	// 1. Master: Central coordinator - divides job into tasks and assigns tasks to workers
	// 2. Mapper: A worker executing a map function
	// 3. Reducer: A worker executing a reduce function

	// Master job:
	// 1. Hands out work to the workers and waits for them to finish
	// 2. Coordinate parrallel execution of tasks

	// The master initializes an RPC server
	// The RPC server allows workers to register themselves with the master
	// Each worker should call Register in master.go
	// The master tells the worker about a new task by using Worker.DoTask (in worker.go)
	// The worker, once receiving the DoTaskArgs, will assume either the mapper or the reducer role based on the DoTaskArgs.Phase.
	// each worker knows from which files to read its input and to which files to write its output.

	// Information about the currently running job is in the master struct in master.go
	// Master does not need to know which Map or Reduce functions are being used for the job
	// the workers will take care of executing the right code for Map or Reduce functions (the correct functions are given to them when they are started by main/wc.go)

	debug("Schedule: %v phase done\n", phase)
}
