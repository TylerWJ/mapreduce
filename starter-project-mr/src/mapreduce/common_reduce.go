package mapreduce

// doReduce does the job of a reduce worker: it reads the intermediate
// key/value pairs (produced by the map phase) for this task, sorts the
// intermediate key/value pairs by key, calls the user-defined reduce function
// (reduceF) for each key, and writes the output to disk.
func doReduce(
	jobName string, // the name of the whole MapReduce job
	reduceTaskNumber int, // which reduce task this is
	nMap int, // the number of map tasks that were run ("M" in the paper)
	reduceF func(key string, values []string) string,
) {
	// TODO:
	// You will need to write this function.
	// You can find the intermediate file for this reduce task from map task number
	// m using reduceName(jobName, m, reduceTaskNumber).
	// Remember that you've encoded the values in the intermediate files, so you
	// will need to decode them. If you chose to use JSON, you can read out
	// multiple decoded values by creating a decoder, and then repeatedly calling
	// .Decode() on it until Decode() returns an error.
	//
	// You should write the reduced output in as JSON encoded KeyValue
	// objects to a file named mergeName(jobName, reduceTaskNumber). We require
	// you to use JSON here because that is what the merger than combines the
	// output from all the reduce tasks expects. There is nothing "special" about
	// JSON -- it is just the marshalling format we chose to use. It will look
	// something like this:
	//
	// enc := json.NewEncoder(mergeFile)
	// for key in ... {
	// 	enc.Encode(KeyValue{key, reduceF(...)})
	// }
	// file.Close()
	//
	// Use checkError to handle errors.

	// Tyler's Notes

	// This function accepts jobName (the name of the whole map reduce job), reduceTaskNumber, the number of map tasks that were run, and the reduceF function
	// doMap = map worker: Reads an input file (inFile), calls the mapF function for that file (inputting the file, and the content), partitions the output into nReduce intermediate files
	// Give mapF a file's name and the content and it returns an array of key/value pairs. Key = word, Value = # of times that word appeared in contents, or a list of 1s???
	// reduceName - constructs the name of the intermediate file which map task - jobName, mapTask, reduceTask

	// Steps:
	// 	1. Read the contents of the input file
	// 	2. Call mapF and pass the file name and the contents of the file - returns an array of key/value pairs for that file
	// 	3. Partitions the output into Reduce intermediate files

	// I have no clue what step 3 means
	// Im guessing you are given a massive list of key/value pairs from mapF. Step three then takes that list and splits it nReduce times.
	// This is also the number of reduce tasks needed for this one file (large list of key/value pairs)
	// After splitting those key value pairs, you then store each partition in a file
	// The files name includes: Which map task produced them and which reduce task are they for. Why do we need the map task???
	// Use JSON to convert key/value data structures to a string and store it in a file.

	// What is mapTask and reduceTask?
	// ihash function below is used to determine which file a given key belongs into

}
