## Overview
CS 4513 - Distributed Systems Project to create a MapReduce Library in Go, focusing on RPC and fault tolerance through four phases.

## Phases

1. **Sequential MapReduce Library**:
   - Implement `doMap()` and `doReduce()`.
   - Validate with Go test suite.

2. **Word Count Application**:
   - Develop `mapF()` and `reduceF()` functions for word count.
   - Verify with tests.

3. **Distributed MapReduce Library**:
   - Modify `schedule()` for task distribution.
   - Test functionality.

4. **Handling Worker Failures**:
   - Enhance `schedule.go` for fault tolerance.
   - Confirm resilience with tests.

## Key Concepts

- **Architecture**:
  - **Master**: Manages tasks and workers.
  - **Workers**: Execute tasks and communicate via RPC.

- **Workflow**:
  - Initialization: Master setup and worker registration.
  - Task Distribution: Assign and monitor tasks.
  - Map Phase: Process files and store results.
  - Reduce Phase: Aggregate data.
  - Completion: Merge results, shut down.

- **Fault Tolerance**:
  - Detect and reassign tasks from failed workers.
  - Ensure idempotent execution.

Experience distributed systems by building a scalable and fault-tolerant MapReduce library.
