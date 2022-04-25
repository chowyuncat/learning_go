The first problem with the naive benchmarking approach is that the time to read the input file adds significant noise to the benchmark, and as a result the number of iterations of the code under test may be higher than necessary.

```
func main() {

	bytes, err := ioutil.ReadFile(os.Args[2]) // this line 
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	text := string(bytes)

	var fp func(s string) WordCountMap

	switch os.Args[1] {
	case "naive":
		fp = WordCountNaive
	default:
		fp = WordCountWithStringFields
	}

	for i := 0; i < 1000; i++ {
		fp(text)
	}
```

+ go run wordcount.go fields romeo_and_juliet.txt
real	0m1.681s
user	0m1.513s
sys	0m0.150s
+ go run wordcount.go naive romeo_and_juliet.txt

real	0m4.321s
user	0m4.258s
sys	0m0.151s
csz@bathtowel learn_benchmark % ./bench.sh
+ SOURCE_FILE=wordcount.go
+ INPUT_FILE=romeo_and_juliet.txt
+ go run wordcount.go fields romeo_and_juliet.txt

real	0m1.511s
user	0m1.490s
sys	0m0.142s
+ go run wordcount.go naive romeo_and_juliet.txt

real	0m4.320s
user	0m4. adds significant
noise to the benchmark:

unc main() {

	bytes, err := ioutil.ReadFile(os.Args[2])
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	text := string(bytes)

	var fp func(s string) WordCountMap

	switch os.Args[1] {
	case "naive":
		fp = WordCountNaive
	default:
		fp = WordCountWithStringFields
	}

	for i := 0; i < 1000; i++ {
		fp(text)
	}
sys	0m0.131s