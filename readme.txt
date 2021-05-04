CS4513: Project 3 The MapReduce Library
=======================================

Note, this document includes a number of design questions that can help your implementation. We highly recommend that you answer each design question **before** attempting the corresponding implementation.
These questions will help you design and plan your implementation and guide you towards the resources you need.
Finally, if you are unsure how to start the project, we recommend you visit office hours for some guidance on these questions before attempting to implement this project.


Team members
-----------------

1. Tyler Jones (twjones@wpi.edu)
2. Tom Graham (Tgraham@wpi.edu)

Design Questions
------------------

(2 point) 1. If there are n input files, and nReduce number of reduce tasks , how does the the MapReduce Library uniquely name the intermediate files?

  f0-0, ... f0-[nReduce-1]
  f1-0, ... f1-[nReduce-1]
  ...
  f[n-1]-0 ... f[n-1][nReduce-1]

  MapReduce calls reduceName and mergeName in common.go to construct the unique file names.
  reduceName uses the job name, map task and reduce task to create the file name ("mrtmp." + jobName + "-" + strconv.Itoa(mapTask) + "-" + strconv.Itoa(reduceTask)
  mergeName uses the job name and the redeuce task to create the file name. ("mrtmp." + jobName + "-res-" + strconv.Itoa(reduceTask)


(1 point) 2. Following the previous question, for the reduce task r, what are the names of files will it work on?

  f0-r,
  f1-r,
  ...
  f[n-1]-r


(1 point) 3. If the submitted mapreduce job name is "test", what will be the final output file's name?

  mrtmp.test


(2 point) 4. Based on `mapreduce/test_test.go`, when you run the `TestBasic()` function, how many master and workers will be started? And what are their respective addresses and their naming schemes?

  1 master, 2 workers.
  Addresses should look similar to:

  /var/tmp/824-UID/mrPID-master
  /var/tmp/824-UID/mrPID-worker0
  /var/tmp/824-UID/mrPID-worker1

  UID and PID can be any integers.


(4 point) 5. In real-world deployments, when giving a mapreduce job, we often start master and workers on different machines (physical or virtual). Describe briefly the protocol that allows master and workers be aware of each other's existence, and subsequently start working together on completing the mapreduce job. Your description should be grounded on the RPC communications.

  You would use RPC still but the protocol that would allow this to happen is TCP rather than using UNIX-domain sockets. So you would need to use a NFS to share storage across multiple machines, and you would also need to be able to start the processes on all the machines.


(2 point) 6. The current design and implementation uses a number of RPC methods. Can you find out all the RPCs and list their signatures? Briefly describe the criteria
a method needs to satisfy to be considered a RPC method. (Hint: you can look up at: https://golang.org/pkg/net/rpc/)

  In order to be considered an RPC method the method must be structured syntaically like this:
  `fun (t *T) myFunc (argType T1, replyType *T2) error`
  When written like that the t1 and t2 represent encoding and gob. Gob is the interface for encoding values to be used in the GobDecoder.


7. We provide a bash script called `main/test-wc-distributed.sh` that allows testing the distributed MapReduce implementation.

  Both of the test scripts work and provide the expected results.


Errata
------

Describe any known errors, bugs, or deviations from the requirements.

---
