# Unique

Unique is a Go package used to determine the count of unique rows within a CSV. It is made to be used on the command line.

## Installation

Using Go, and if you have `$GOPATH/bin` in your path, you can install using:

```bash
go get -u github.com/jasonmccallister/unique
```

## Usage

The only required argument for `unique` is the path to the file.

```bash
unique orders.csv
```

Output:

```bash
CSV processed!
Here are your results:

Using Column:    0
Total Uniques:   2
Total Rows:      4
```

> The above command will default to using the first column (`0`) of the CSV as the row.

### Specifying the Column

If you need to specify the column, you can pass that as a flag.

```bash
unique orders.csv -column=2
```

_**Note**: If you have a large dataset, its best to limit the number of columns. For example, export only the primary key and the value you need to find the uniques._

## Security

If you discover any security related issues, please email themccallister@gmail.com instead of using the issue tracker.

## Credits

- [Jason McCallister](https://github.com/jasonmccallister)
- [All Contributors](../../contributors)

## License

The MIT License (MIT). Please see [License File](LICENSE.md) for more information.

Brought to you by [Jason McCallister](https://mccallister.io)
