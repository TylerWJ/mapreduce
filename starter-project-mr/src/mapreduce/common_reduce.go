package mapreduce

import (
	"encoding/json"
	"os"
)

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

	var keyValues []KeyValue

	decoders := make([]*json.Decoder, nMap)
	oFiles := make([]*os.File, nMap)

	for i := 0; i < nMap; i++ { // for each file
		//fmt.Println("Hello")
		var readErr error
		fName := reduceName(jobName, i, reduceTaskNumber) // construct the intermediate file name
		oFiles[i], readErr = os.Open(fName)               // find and open the intermediate file
		checkError(readErr)

		decoders[i] = json.NewDecoder(oFiles[i]) // creates a decoder for the file
		for {
			var keyVal KeyValue
			decErr := decoders[i].Decode(&keyVal)
			if decErr != nil {
				break
			} else {
				keyValues = append(keyValues, keyVal)
			}
		}

		defer oFiles[i].Close()

	}

	kvHash := make(map[string][]string) // Key: String, Value: String array

	for i := 0; i < len(keyValues); i++ {
		kvHash[keyValues[i].Key] = append(kvHash[keyValues[i].Key], keyValues[i].Value)
	}

	fNameMerge := mergeName(jobName, reduceTaskNumber) // create the name of the merge file
	mf, createErr := os.Create(fNameMerge)             // create the merge file
	checkError(createErr)

	enc := json.NewEncoder(mf) // create an encoder for the merge file
	for _, key := range keyValues {
		enc.Encode(KeyValue{key.Key, reduceF(key.Key, kvHash[key.Key])})
	}

	defer mf.Close()

}
