### Jobs  
---
##### Jobs
```
Jobs
   a) manages one or more pods to execute some operation until it is successful
   b) if a pod fails, then a new pod runs to finish the operation 
   c) jobs and their pods will not be cleared after completion (delete the job to clear all resources)
``` 

##### Example
```yaml
apiVersion: batch/v1
kind: Job
metadata:
  name: factorial5
spec:
  # by default, this value is 1
  # specifies how many successful completions are required 
  completions: 3

  # determines max number of pods to launch to run job in parallel
  # job will not launch more pods than required
  parallelism: 5

  template:
    spec:
      containers:
      - name: factorial5
        image: some-python-image-1.32
        command: ["python",  "-c", "from math import factorial; print(factorial(5))"]
      # job should not restart after completing
      restartPolicy: Never

# here only 3 pods will run in parallel as not all 5 are required
# when a job is done, pods that completed the task will be status "completed"
# run kubectl logs <pod-name> to view output of the job
```

##### Cron Jobs
```
Cron Jobs
   a) jobs that run periodically
   b) each invocation launches a new job object along with corresponding pods
   c) deleting a cron job will delete existing jobs and pods
```

### Probing & Readiness
---
##### Probes
```
Startup Probes
   a) defines what it means for a container to have started
   b) disables liveness and readiness checks until startup succeeds
   c) provides time for the container application to initialize

Liveness Probes
   a) defines what it means for a container to be alive
   b) kubelets have basic restart policies if a container process crashes, but it might not be sufficient

Readiness Probes
   a) defines what it means for a container application to be able to serve requests
   b) container may be initialized, but necessary dependencies might still be or have become unavailable 
   c) when readiness probe fails, containers are not killed but requests will not be received
   d) ensures requests don't flood a pod with requests it cannot process
```

##### Readiness Gates
```
Problems
   a) readiness probes help address pod level readiness, but not an infrastructure level readiness
   b) services, network policies forwarding traffic might not be ready yet

Readiness Gates : provide an extra podSpec to specify set of conditions for when things are ready
```

##### Probe Examples
```yaml
apiVersion: v1
kind: Pod
spec:
  containers:
    - name: app
      image: test-app
      args:
      - /bin/sh
      - -c
      - touch /tmp/healthy; sleep 30; rm -rf /tmp/healthy; sleep 600
      livenessProbe:
        exec:
          command: 
          - cat 
          - /tmp/healthy
        # give time for the pod to initialize
        # kubelet will not check liveness for this period
        initialDelaySeconds: 30
```

##### Init Containers
```
Init Containers
   a) containers that run to completion before all other containers start
   b) can configure such that probes will start only after init containers are started
   c) init containers can provide setup scripts
```