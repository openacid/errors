2019 Feb 19

```
BenchmarkErrors/openacid/errors-stack-10-8               1000000              1083 ns/op             320 B/op          3 allocs/op
BenchmarkErrors/errors-stack-10-8                       30000000                51.2 ns/op            16 B/op          1 allocs/op
BenchmarkErrors/openacid/errors-stack-100-8              1000000              2111 ns/op             320 B/op          3 allocs/op
BenchmarkErrors/errors-stack-100-8                       5000000               384 ns/op              16 B/op          1 allocs/op
BenchmarkErrors/openacid/errors-stack-1000-8              200000              8204 ns/op             320 B/op          3 allocs/op
BenchmarkErrors/errors-stack-1000-8                       500000              3638 ns/op              16 B/op          1 allocs/op
BenchmarkStackFormatting/%s-stack-10-8                  10000000               169 ns/op               8 B/op          1 allocs/op
BenchmarkStackFormatting/%v-stack-10-8                  10000000               173 ns/op               8 B/op          1 allocs/op
BenchmarkStackFormatting/%+v-stack-10-8                   200000              9977 ns/op            1938 B/op         19 allocs/op
BenchmarkStackFormatting/%s-stack-30-8                  10000000               169 ns/op               8 B/op          1 allocs/op
BenchmarkStackFormatting/%v-stack-30-8                  10000000               171 ns/op               8 B/op          1 allocs/op
BenchmarkStackFormatting/%+v-stack-30-8                   100000             20522 ns/op            4368 B/op         33 allocs/op
BenchmarkStackFormatting/%s-stack-60-8                  10000000               172 ns/op               8 B/op          1 allocs/op
BenchmarkStackFormatting/%v-stack-60-8                  10000000               172 ns/op               8 B/op          1 allocs/op
BenchmarkStackFormatting/%+v-stack-60-8                   100000             20397 ns/op            4368 B/op         33 allocs/op
BenchmarkStackFormatting/%s-stacktrace-10-8               500000              3207 ns/op             240 B/op          2 allocs/op
BenchmarkStackFormatting/%v-stacktrace-10-8               200000              6059 ns/op             300 B/op          5 allocs/op
BenchmarkStackFormatting/%+v-stacktrace-10-8              200000              7718 ns/op            1836 B/op          5 allocs/op
BenchmarkStackFormatting/%s-stacktrace-30-8               200000              6408 ns/op             512 B/op          2 allocs/op
BenchmarkStackFormatting/%v-stacktrace-30-8               100000             12076 ns/op             608 B/op          2 allocs/op
BenchmarkStackFormatting/%+v-stacktrace-30-8              100000             15934 ns/op            4133 B/op          2 allocs/op
BenchmarkStackFormatting/%s-stacktrace-60-8               200000              6344 ns/op             512 B/op          2 allocs/op
BenchmarkStackFormatting/%v-stacktrace-60-8               100000             12158 ns/op             608 B/op          2 allocs/op
BenchmarkStackFormatting/%+v-stacktrace-60-8              100000             15533 ns/op            4132 B/op          2 allocs/op
```
