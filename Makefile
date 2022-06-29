.PHONY: bench-default
bench-default:
	go test -bench=BenchmarkWSDefault --count=30 > default.txt

.PHONY: bench-custom
bench-custom:
	go test -bench=BenchmarkWSCustom --count=30 > custom.txt

.PHONY: bench-join
bench-join:
	go test -bench=BenchmarkJoin --count=30 > join.txt

.PHONY: bench-parallel-default
bench-parallel-default:
	go test -bench=BenchmarkParallelWSDefault --count=30 > parallel_default.txt

.PHONY: bench-parallel-custom
bench-parallel-custom:
	go test -bench=BenchmarkParallelWSCustom --count=30 > parallel_custom.txt

.PHONY: bench-parallel-join
bench-parallel-join:
	go test -bench=BenchmarkParallelJoin --count=30 > parallel_join.txt

