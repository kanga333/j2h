# j2h

j2h is a tool to convert json to hive ddl.

## Installation

```sh
$ go get github.com/kanga333/j2h
```
or  

Download the binary directly from the release page.

## Usage

```sh
$ j2h -help
j2h is a tool to convert json to hive ddl

Usage: j2h <option>
  -json-path string
        Path of json file.
  -version
        Print version information.
```

## Example

```json:test.json
{
  "foo": {
    "bar": [
      10,
      21,
      20
    ],
    "baz": [
      [
        1.1,
        1.2
      ],
      [
        1.3,
        1.4
      ]
    ],
    "hoge": "string"
  },
  "piyo": true
}
```

```sh
$ j2h -path test.json
create external table json_data(
  foo struct<
    bar:array<int>,
    baz:array<
      array<double>
    >,
    hoge:string
  >,
  piyo boolean
)
```

## Restrictions

- Hive Reserved words are output in lowercase letters.
- It does not correspond to the output that converts json to map of hive.
- All integers are output as int type.
- All decimals are output as double type.
- The null type of json is converted to the binary type of hive.
- If the array element type of json is mixed, it is converted to binary type of hive.

