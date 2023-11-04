count = 10
bench:
	echo "Running benchmark"
	go test -benchmem -count=$(count) -bench . -run ^$$ > ./outputs/bench_result.txt
	benchstat -col "/engine@(templ native)" ./outputs/bench_result.txt > ./outputs/benchstat_result.txt
	cat ./outputs/benchstat_result.txt