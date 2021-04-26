package mapreduce

import (
	"encoding/json"
	"hash/fnv"
	"io/ioutil"
	"os"
)

// doMap does the job of a map worker: it reads one of the input files
// (inFile), calls the user-defined map function (mapF) for that file's
// contents, and partitions the output into nReduce intermediate files.
func doMap(
	jobName string, // the name of the MapReduce job
	mapTaskNumber int, // which map task this is
	inFile string,
	nReduce int, // the number of reduce task that will be run ("R" in the paper)
	mapF func(file string, contents string) []KeyValue,
) {
	// TODO:
	// You will need to write this function.
	// You can find the filename for this map task's input to reduce task number
	// r using reduceName(jobName, mapTaskNumber, r). The ihash function (given
	// below doMap) should be used to decide which file a given key belongs into.
	//
	// The intermediate output of a map task is stored in the file
	// system as multiple files whose name indicates which map task produced
	// them, as well as which reduce task they are for. Coming up with a
	// scheme for how to store the key/value pairs on disk can be tricky,
	// especially when taking into account that both keys and values could
	// contain newlines, quotes, and any other character you can think of.
	//
	// One format often used for serializing data to a byte stream that the
	// other end can correctly reconstruct is JSON. You are not required to
	// use JSON, but as the output of the reduce tasks *must* be JSON,
	// familiarizing yourself with it here may prove useful. You can write
	// out a data structure as a JSON string to a file using the commented
	// code below. The corresponding decoding functions can be found in
	// common_reduce.go.
	//
	//   enc := json.NewEncoder(file)
	//   for _, kv := ... {
	//     err := enc.Encode(&kv)
	//
	// Remember to close the file after you have written all the values!
	// Use checkError to handle errors.

	content, readErr := ioutil.ReadFile(inFile) // read the content of the file
	checkError(readErr)

	keyValues := mapF(inFile, string(content)) // collect the key values of the file
	size := len(keyValues) / nReduce           // ceil/floor???

	for i := 0; i < nReduce; i++ { // this creates nReduce subfile names for the given file

		fName := reduceName(jobName, mapTaskNumber, i)  // creates the file name
		f, createErr := os.Create("../ofiles/" + fName) // stores the output files in the ofiles folder
		checkError(createErr)

		start := i * size // the starting point in the key/value array for each subfile
		end := start + size
		subContent := keyValues[start:end] // end is exclusive

		enc := json.NewEncoder(f)
		for _, kv := range subContent {
			encErr := enc.Encode(&kv)
			checkError(encErr)
		}

		// Questions:
		// ihash???
		// location of output files???
		// ceil? will cause out of bounds in for loop

	}

	// Tyler's Notes

	// 	Application:
	// master.go creates a master_rpc server for workers to register
	// workers will register using the RPC call Register. The workers will also start up their own RPC servers so that master can dispatch them tasks
	// Workers register using RPC call Register()
	// RPC allows processes to communicate with one another
	// As tasks become available, master.go uses schedule() in scedule.go to assign the different tasks to the different workers (and how to handle worker failure)
	// Each input file = 1 map task
	// Master then makes a call to doMap atleast once for each task, Sequential -> doMap() directly, Distributed -> DoTask() in worker.go to give the task to a worker
	// Each call to do map does:
	// 	1. Read the contents of the input file
	// 	2. Call mapF and passes the file name and the contents of the file - returns an array of key/value pairs for that file
	// 	3. Partitions the output into nReduce files
	// For the ith map task, it will generate a list of files with the following naming pattern: fi-0, fi-1 ... fi-[nReduce-1]
	// So, the total number of files = # of files * nReduce (partition files per file)
	// The master then calls doReduce() atleast once for each reduce task, Sequential -> doMap() directly, Distributed -> DoTask() in worker.go to give the task to a worker
	// For teh jth doReduce() call, doReduce() will go through f0-j, f1-j, ..., f[n-1]-j
	// Basically, doMap goes splits each file into R subfiles, and then doReduce iterates through all of the files R times and works only on a subsection of each file
	// After reduce, master calls mr.merge() in master_splitmerge.go which merges the the nReduce files from the previous step
	// Lastly, master shuts down each worker's RPC and then finally, it's own

	// Where do we store all of these partition files???
	// What is the purpose of ihash?

	// This function accepts jobName, mapTaskNumber, a file, number of reduce tasks that will be run, and the mapF function
	// doMap = map worker: Reads an input file (inFile), calls the mapF function for that file (inputting the file, and the content), partitions the output into nReduce intermediate files
	// Give mapF a file's name and the content and it returns an array of key/value pairs. Key = word, Value = # of times that word appeared in contents, or a list of 1s???

	//	Helper functions:
	// reduceName - constructs the name of the intermediate file which map task - jobName, mapTask, reduceTask
	// ihash function below is used to determine which file a given key belongs into

	// Im guessing you are given a massive list of key/value pairs from mapF. Step three then takes that list and splits it nReduce times.
	// This is also the number of reduce tasks needed for this one file (large list of key/value pairs)
	// After splitting those key value pairs, you then store each partition in a file
	// The files name includes: Which map task produced them and which reduce task are they for. Why do we need the map task???
	// Use JSON to convert key/value data structures to a string and store it in a file.

}

func ihash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
