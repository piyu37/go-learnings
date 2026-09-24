1. Go Module Manipulation
Question:

You want to call a certain piece of code from a newly created Go module. When doing so, you want to replace the package containing your module - sample.com/new with ../new.

Which of these commands should you run from the command prompt to achieve this?

Options:

$ go mod edit -replace sample.com/new = ../new

replace sample.com/new => ../new

Either Choice 1 or 2

None of these

✅ Correct Answer:

Either Choice 1 or 2

2. Redeclarations and Reassignments in Go
Question:

You observe that in a := declaration a variable v appears even if it has already been declared. Which of the following is true in the given context?

The declaration is in the same scope as the existing declaration of v.

The corresponding value in the initialization is NOT assignable to v.

Options:

Only 1

Only 2

Both 1 and 2

Neither 1 nor 2

✅ Correct Answer:

Only 2

3. Creating Modules in Go
Question:

You want to call a certain piece of code from a newly created Go module. When doing so, you want to replace the package containing your module - sample.com/new with ../new.

What is the purpose of doing this?

Options:

It helps in locating the dependency

It helps in discriminating the fields in the code from the method

It helps in repurposing the code

Both Choice 1 and 2

Both Choice 1 and 3

✅ Correct Answer:

It helps in locating the dependency

4. Handling Errors in Golang
Question:

When handling errors in golang, you observe that there is some missing information about the invalid argument returned by the error. To tackle this issue, you decide to format the string according to the printf rules and return it as an error created by errors.New.

Which of these functions should you use to successfully implement this?

Options:

fmt package's Errorf function

The errors.errorString function

fmt package's Println function

fmt package's log.fatal function

✅ Correct Answer:

fmt package's Errorf function

5. Handling file errors
Question:

What does the following code return when it fails to open a file?

func Open(name string) (file *File, err error)

Options:

It returns a non-nil error

It returns an error of type errorString

It returns an internal server error

None of these

✅ Correct Answer:

It returns a non-nil error

6. Working with Goroutines
Question:

Analyze the following snippet and choose the valid inferences:

var xyz = make(chan int, 3)

func main() {
    for _, w := range work {
        go func() {
            xyz <- 1
            w()
            <-xyz
        }()
    }
    select {}
}


Statements:

The program starts a goroutine for every entry in the work list.

The program ensures that at most three goroutines are running work functions at a time.

Options:

Only 1

Only 2

Both 1 and 2

Neither 1 nor 2

✅ Correct Answer:

Both 1 and 2

(Buffered channel acts as concurrency limiter)

7. Optimize Concurrent Map Access
Question:

Analyze this Go snippet for concurrent map updates causing contention:

var mu sync.Mutex
var m = make(map[int]int)

func add(k, v int) {
    mu.Lock()
    for i := 0; i < 100000; i++ {
        m[k+i] = v
    }
    mu.Unlock()
}

Options:

Replace sync.Mutex with sync.RWMutex so multiple goroutines can write concurrently.

Serialize all writes using a single writer goroutine and send update requests over a channel.

Minimize the locked section: do heavy computation outside the mutex and lock only to apply small batches or use sharded maps.

Increase GOMAXPROCS to reduce mutex contention automatically.

✅ Correct Answer:

Minimize the locked section: do heavy computation outside the mutex and lock only to apply small batches or use sharded maps.

