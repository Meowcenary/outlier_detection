# outlier_detection
An attempt at outlier detection in Go using the isolation forest algorithm

### Installation
This repository is not formed into an executable package and is instead simply three Go programs that are intended to be
compiled separatley. The repository can be cloned with: `git clone git@github.com:Meowcenary/outlier_detection.git` or
`git clone https://github.com/Meowcenary/outlier_detection.git`

### Package selection
The package ["randomForest](https://github.com/malaschitz/randomForest) was selected for testing the ioslation forest
algorithm. The package was selected because it was listed in the curated list
[awesome-go](https://github.com/avelino/awesome-go), had more activity than other packages with similar functionality,
and  had the most stars on Github.

### Results
It was easy enough to install the package as you would install any Go package, but it was challenging to understand how
the results from the Python and R examples provided could be compared to the results from running the randomForest
package. In general the randomForest package was littered with code that was difficult to understand. There were several
examples and tests included that helped, but the variable names selected were so vague that a lot of time was spent just
trying to understand what the examples were returning and even with that effort it wasn't entirely clear how the results
could be compared to the R and Python script results.

Given the relative complexity of working with a language like Go, the low activity on what is ostensibly considered the
best implementation of the isolation forest algorithm in Go, and the difficulty of working with that same package it is
my strong recommendation to work with R, or better yet, Python instead of Go at this time. Python has the benefit of
being a more general purpose langauge than R while still having a relatively simple syntax. It's also faster than R, but
admittedly not as fast as Go. Python also has the benefit of being a popular langauge in the data science field which
makes it easier to hire people who can work in it versus Go.

### Programs
There are three Go files included in this repository that can be compiled into small programs. main.go

- randomforest.go is an annotated version of the random forest example from [randomForest](https://github.com/malaschitz/randomForest/blob/master/examples/isolation.go)
- isolationforest.go is an annotated version of the isolation forest example from [randomForest](https://github.com/malaschitz/randomForest/blob/master/examples/isolation2.go)
- isolatforestcustom.go is an attempt to use randomForest.IsolationForest on the MNIST data set, but it wasn't clear to
me how to actually use the results to calculate an anomaly score for each value. I tried to look through the source code
for the Python library scikit-learn and referened the wikipedia page for isolation forest, but I found it all to be
confusing.
