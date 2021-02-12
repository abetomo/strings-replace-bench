# strings-replace-bench

```
% go version
go version go1.15.7 darwin/amd64
% go test -bench . -benchmem
goos: darwin
goarch: amd64
pkg: github.com/abetomo/strings-replace-bench
BenchmarkStringsReplace/2_chars/Replace_all-8         	 5196727	       231 ns/op	     320 B/op	       4 allocs/op
BenchmarkStringsReplace/2_chars/Replace_some-8        	 5231845	       230 ns/op	     320 B/op	       4 allocs/op
BenchmarkStringsReplace/2_strings/Replace_all-8       	 3911292	       305 ns/op	     320 B/op	       4 allocs/op
BenchmarkStringsReplace/2_strings/Replace_some-8      	 7506202	       164 ns/op	     160 B/op	       2 allocs/op
BenchmarkStringsReplace/5_chars/Replace_all-8         	 2042359	       589 ns/op	     864 B/op	      10 allocs/op
BenchmarkStringsReplace/5_chars/Replace_some-8        	 2501830	       476 ns/op	     672 B/op	       8 allocs/op
BenchmarkStringsReplace/5_strings/Replace_all-8       	 1536916	       790 ns/op	     896 B/op	      10 allocs/op
BenchmarkStringsReplace/5_strings/Replace_some-8      	 2303527	       528 ns/op	     544 B/op	       6 allocs/op
BenchmarkStringsReplacer/2_chars/Replace_all.-8       	10080566	       119 ns/op	     160 B/op	       2 allocs/op
BenchmarkStringsReplacer/2_chars/Replace_some.-8      	 9946582	       121 ns/op	     160 B/op	       2 allocs/op
BenchmarkStringsReplacer/2_strings/Replace_all.-8     	 3955276	       305 ns/op	     192 B/op	       3 allocs/op
BenchmarkStringsReplacer/2_strings/Replace_some.-8    	 4732401	       254 ns/op	     192 B/op	       3 allocs/op
BenchmarkStringsReplacer/5_chars/Replace_all.-8       	 4766077	       241 ns/op	     192 B/op	       2 allocs/op
BenchmarkStringsReplacer/5_chars/Replace_some.-8      	 5330580	       224 ns/op	     192 B/op	       2 allocs/op
BenchmarkStringsReplacer/5_strings/Replace_all.-8     	 2310986	       520 ns/op	     368 B/op	       4 allocs/op
BenchmarkStringsReplacer/5_strings/Replace_some.-8    	 2884166	       408 ns/op	     368 B/op	       4 allocs/op
PASS
ok  	github.com/abetomo/strings-replace-bench	24.926s
```

## Topic

```
BenchmarkStringsReplace/2_strings/Replace_all-8       	 3911292	       305 ns/op	     320 B/op	       4 allocs/op
BenchmarkStringsReplace/2_strings/Replace_some-8      	 7506202	       164 ns/op	     160 B/op	       2 allocs/op
```

```
BenchmarkStringsReplacer/2_strings/Replace_all.-8     	 3955276	       305 ns/op	     192 B/op	       3 allocs/op
BenchmarkStringsReplacer/2_strings/Replace_some.-8    	 4732401	       254 ns/op	     192 B/op	       3 allocs/op
```
