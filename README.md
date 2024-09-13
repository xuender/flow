# flow

[![Action][action-svg]][action-url]
[![Report Card][goreport-svg]][goreport-url]
[![Lines of code][lines-svg]][lines-url]
[![godoc][godoc-svg]][godoc-url]
[![License][license-svg]][license-url]

✨ **`xuender/flow` is a ...**

## 🚀 Install

```shell
go install gitee.com/xuender/flow@latest
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

### Flow Functions

* Chain
* Parallel

### Intermediate Functions

* Filter
* Limit
* Peek
* Reduce
* Reverse
* Skip
* Sort
* SortFunc
* Unique

### Terminal Functions

* seq.Count
* seq.Each
* seq.First
* seq.Max
* seq.Min
* seq.Sum
* seq.Join
* seq.AnyMatch
* seq.AllMatch
* seq.NoneMatch

### Seq Functions

* seq.Chan
* seq.Emit
* seq.Range
* seq.Filter
* seq.FlatMap
* seq.Limit
* seq.Map
* seq.Peek
* seq.Reduce
* seq.Reverse
* seq.Skip
* seq.Sort
* seq.SortFunc
* seq.ToChans
* seq.Unique


## 👤 Contributors

![Contributors][contributors-svg]

## 📝 License

© ender, 2024~time.Now

[MIT LICENSE][license-url]

[action-url]: https://gitee.com/xuender/flow/actions
[action-svg]: https://gitee.com/xuender/flow/workflows/Go/badge.svg

[goreport-url]: https://goreportcard.com/report/gitee.com/xuender/flow
[goreport-svg]: https://goreportcard.com/badge/gitee.com/xuender/flow

[godoc-url]: https://godoc.org/gitee.com/xuender/flow
[godoc-svg]: https://godoc.org/gitee.com/xuender/flow?status.svg

[license-url]: https://gitee.com/xuender/flow/blob/master/LICENSE
[license-svg]: https://img.shields.io/badge/license-MIT-blue.svg

[contributors-svg]: https://contrib.rocks/image?repo=xuender/flow

[lines-svg]: https://sloc.xyz/gitee/xuender/flow
[lines-url]: https://github.com/boyter/scc
