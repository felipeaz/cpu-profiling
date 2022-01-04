# cpu-profiling
Basic example of CPU Profiling in Golang which shows the bottlenecks and how much time is 
spent per function

## Generate cpu.prof
Run the following command to create the cpu.prof file

`` go test -cpuprofile cpu.prof -bench .=``

## Analyze functions
First connect to pprof by running the command ``go tool pprof cpu.prof``

Then, run the ``list`` command followed by the function name

``list nonDivisibleSubset``

``list generateSubset``

``list updateSubset``