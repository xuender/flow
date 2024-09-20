# flow

[![Action][action-svg]][action-url]
[![Report Card][goreport-svg]][goreport-url]
[![codecov][codecov-svg]][codecov-url]
[![Lines of code][lines-svg]][lines-url]
[![godoc][godoc-svg]][godoc-url]
[![License][license-svg]][license-url]

✨ **`xuender/flow` is a streaming programming library based on `iter` for Go 1.23.**

## ❗ Prerequirement

- Go 1.23+
- Not dependent on other libraries

## 🚀 Install

```shell
go get github.com/xuender/flow@latest
```

## 💡 Usage

### Chain

```go
package main

import (
  "fmt"

  "github.com/xuender/flow"
  "github.com/xuender/flow/seq"
)

func main() {
  start := time.Now()

  for num :=range flow.Chain(
    seq.Range(100),
    flow.Filter(func(num int) bool { return num%3 == 0 }),
    flow.Skip[int](5),
    flow.Limit[int](4),
    flow.Peek(func(num int) {
      fmt.Println("peek", num)
      time.Sleep(time.Second)
    }),
    flow.Reverse[int](),
  ) {
    fmt.Println(num)
  }

  fmt.Printf("Chain: %v\n", time.Since(start))
}
```

Output:

```shell
peek 15
peek 18
peek 21
peek 24
24
21
18
15
Chain: 4s
```

[[play](https://go.dev/play/p/-qmIigTfHXt)]

### Parallel

```go
package main

import (
  "fmt"

  "github.com/xuender/flow"
  "github.com/xuender/flow/seq"
)

func main() {
  start := time.Now()

  for num := range flow.Parallel(3,
    seq.Range(100),
    flow.Filter(func(num int) bool { return num%3 == 0 }),
    flow.Skip[int](5),
    flow.Limit[int](4),
    flow.Peek(func(num int) {
      fmt.Println("peek", num)
      time.Sleep(time.Second)
    }),
    flow.Reverse[int](),
  ) {
    fmt.Println(num)
  }

  fmt.Printf("Parallel: %v\n", time.Since(start))
}
```

Output:

```shell
peek 15
peek 18
peek 21
peek 24
24
21
18
15
Parallel: 2s
```

[[play](https://go.dev/play/p/1E76NdEwaoP)]

### seq.Range

```go
package main

import (
  "fmt"

  "github.com/xuender/flow/seq"
)

func main() {
  fmt.Println("seq.Range(3)")

  for num := range seq.Range(3) {
    fmt.Println(num)
  }

  fmt.Println("\nseq.Range(-2, -6)")

  for num := range seq.Range(-2, -6) {
    fmt.Println(num)
  }

  fmt.Println("\nseq.Range(2, 10, 3)")

  for num := range seq.Range(2, 10, 3) {
    fmt.Println(num)
  }

  fmt.Println("\nseq.Range(3, 7, 0)")

  for num := range seq.Range(3, 7, 0) {
    fmt.Println(num)
  }
}
```

Output:

```shell
seq.Range(3)
0
1
2

seq.Range(-2, -6)
-2
-3
-4
-5

seq.Range(2, 10, 3)
2
5
8

seq.Range(3, 7, 0)
3
3
3
3
```

[[play](https://go.dev/play/p/-FZJfetngJY)]

### Flow Functions

| iter.Seq[V] | iter.Seq2[K,V] |
| ----------- | -------------- |
| Chain       | Chain2         |
| Parallel    | Parallel2      |

### Intermediate Functions

| iter.Seq[V]       | iter.Seq2[K,V]  | Note                         |
| ----------------- | --------------- | ---------------------------- |
| Append            | Append2         |                              |
| CenteredMovingAvg |                 |                              |
| Distinct          | Distinct2       |                              |
| DropWhile         | DropWhile2      |                              |
| Filter            | Filter2         | Parallel                     |
| Limit             | Limit2          |                              |
| Map               | Map2            | Parallel, Cannot change type |
| Merge             | Merge2          |                              |
| MovingAvg         |                 |                              |
| Peek              | Peek2           | Parallel                     |
| Prepend           | Prepend2        |                              |
| Repeat            | Repeat2         |                              |
| Reverse           | Reverse2        | Collecting                   |
| Shuffle           | Shuffle2        | Collecting                   |
| Skip              | Skip2           |                              |
| Sort              | Sort2           | Collecting                   |
| SortFunc          | SortFunc2       | Collecting                   |
| SortStableFunc    | SortStableFunc2 | Collecting                   |
| TakeWhile         | TakeWhile2      |                              |

### Terminal Functions

| iter.Seq[V]   | iter.Seq2[K,V] |
| ------------- | -------------- |
| seq.AllMatch  | seq.AllMatch2  |
| seq.AnyMatch  | seq.AnyMatch2  |
| seq.Clone     | seq.Clone2     |
| seq.Chunk     | seq.Chunk2     |
| seq.Count     | seq.Count2     |
| seq.Each      | seq.Each2      |
| seq.Emit      | seq.Emit2      |
| seq.First     | seq.First2     |
| seq.Last      | seq.Last2      |
| seq.Join      |                |
| seq.Max       | seq.Max2       |
| seq.MaxFunc   | seq.MaxFunc2   |
| seq.Min       | seq.Min2       |
| seq.MinFunc   | seq.MinFunc2   |
| seq.NoneMatch | seq.NoneMatch2 |
| seq.Reduce    | seq.Reduce2    |
| seq.Sum       |                |

### Other Seq Functions

| iter.Seq[V]           | iter.Seq2[K,V]        | Note                     |
| --------------------- | --------------------- | ------------------------ |
| seq.Append            | seq.Append2           |                          |
| seq.CenteredMovingAvg |                       | Boundary effect at end   |
| seq.Chan              | seq.Chan2             |                          |
| seq.Collect           | seq.Collect2          |                          |
| seq.Distinct          | seq.Distinct2         |                          |
| seq.DropWhile         | seq.DropWhile2        |                          |
| seq.Filter            | seq.Filter2           |                          |
| seq.FlatMap           |                       |                          |
|                       | seq.Keys              |                          |
| seq.Limit             | seq.Limit2            |                          |
| seq.Map               | seq.Map2              |                          |
| seq.Merge             | seq.Merge2            |                          |
| seq.MovingAvg         |                       | Boundary effect at start |
| seq.Partition         | seq.Partition2        | Block                    |
| seq.Peek              | seq.Peek2             |                          |
| seq.Prepend           | seq.Prepend2          |                          |
| seq.Range             | seq.Range2            |                          |
| seq.Reduce            | seq.Reduce2           |                          |
| seq.Repeat            | seq.Repeat2           |                          |
| seq.Reverse           | seq.Reverse2          |                          |
| seq.Send              | seq.Send2             |                          |
| seq.Shuffle           | seq.Shuffle2          |                          |
| seq.Skip              | seq.Skip2             |                          |
| seq.Sorted            | seq.Sorted2           |                          |
| seq.SortedFunc        | seq.SortedFunc2       |                          |
| seq.SortedStableFunc  | seq.SortedStableFunc2 |                          |
| seq.TakeWhile         | seq.TakeWhile2        |                          |
| seq.ToChans           | seq.ToChans2          |                          |
|                       | seq.Tuples            |                          |
|                       | seq.Values            |                          |
| seq.Window            | seq.Window2           |                          |
| seq.Zip               |                       |                          |

### Custom Step[V] / Step2[K,V] function

Refer to `Append`, `Sored`, `Filter`, etc.

ps: Attention Parallel.

## 👤 Contributors

![Contributors][contributors-svg]

## 📝 License

© ender, 2024~time.Now

[MIT LICENSE][license-url]

[action-url]: https://github.com/xuender/flow/actions
[action-svg]: https://github.com/xuender/flow/workflows/Go/badge.svg

[goreport-url]: https://goreportcard.com/report/github.com/xuender/flow
[goreport-svg]: https://goreportcard.com/badge/github.com/xuender/flow

[codecov-url]: https://codecov.io/gh/xuender/flow
[codecov-svg]: https://codecov.io/gh/xuender/flow/graph/badge.svg?token=1VAC5OJJZR

[godoc-url]: https://godoc.org/github.com/xuender/flow
[godoc-svg]: https://godoc.org/github.com/xuender/flow?status.svg

[license-url]: https://github.com/xuender/flow/blob/master/LICENSE
[license-svg]: https://img.shields.io/badge/license-MIT-blue.svg

[contributors-svg]: https://contrib.rocks/image?repo=xuender/flow

[lines-svg]: https://sloc.xyz/gitee/xuender/flow
[lines-url]: https://github.com/boyter/scc
