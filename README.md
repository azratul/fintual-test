# How to run the program

1. Clone the project


```bash
git clone git@github.com:azratul/fintual-test.git

```

2. CD to directory


```bash
cd fintual-test

```


3. Run the program


```bash
go run *.go

```

4. Help!!! You can run the program with your own dates


```bash
go run *.go -h

```

Example


```bash
go run *.go -start "2021-03-03" -end "2022-01-25"

```


# Options

You can run the program with some basics options, Ex:

1. Populate another JSON file, set a new URL and activate debugging


```bash
export POPULATE=true
export STOCKS_URL="https://my.url-with.stocks"
export DEBUG=true
```


or


```bash
POPULATE=true STOCK_URL="https://my.url-with.stocks" DEBUG=true go run *.go
```

