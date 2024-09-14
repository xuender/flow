# flow

[![Action][action-svg]][action-url]
[![Report Card][goreport-svg]][goreport-url]
[![codecov][codecov-svg]][codecov-url]
[![Lines of code][lines-svg]][lines-url]
[![godoc][godoc-svg]][godoc-url]
[![License][license-svg]][license-url]

✨ **`xuender/flow` is a streaming programming library based on iterators for Go 1.23.**

## ❗ Prerequirement

- Go 1.23+

## 🚀 Install

```shell
go get github.com/xuender/flow@latest
```

## 💡 Usage

```go
func main() {
	seq.Each(flow.Chain(
		seq.Range(100),
		flow.Filter(func(num int) bool { return num%3 == 0 }),
		flow.Skip[int](5),
		flow.Limit[int](4),
		flow.Reverse[int](),
	), func(num int) bool {
		fmt.Println(num)

		return true
	})
}

```

```shell
24
21
18
15
```

### Flow Functions

- Chain
- Parallel

### Intermediate Functions

- Distinct
- Filter
- Limit
- Peek
- Reverse
- Skip
- Sort
- SortFunc

### Terminal Functions

- seq.Count
- seq.Each
- seq.First
- seq.Max
- seq.Min
- seq.Reduce
- seq.Sum
- seq.Join
- seq.AnyMatch
- seq.AllMatch
- seq.NoneMatch

### Seq Functions

- seq.Chan
- seq.Distinct
- seq.Emit
- seq.Range
- seq.Filter
- seq.FlatMap
- seq.Limit
- seq.Map
- seq.Peek
- seq.Reduce
- seq.Reverse
- seq.Skip
- seq.Sort
- seq.SortFunc
- seq.ToChans


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
