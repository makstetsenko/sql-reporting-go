# sql reports

App takes a directory with sql scripts as input, executes scripts over provided ms sql databases and makes reports.

## Usage

Just run using `go`
```bash
go run . --connections-config=./connections.example.yaml --sql-scripts=./scripts
```


### --sql-scripts=./scripts-example

Path to directory with sql scripts

```
./scripts
├── select-customers.sql
└── select-products.sql
```

### --connections-config=./connections.yaml

Configs to connect to databases.

App runs scripts over all DB provided in `connections.yaml` file.

Use `connections.example.yaml` as reference.


### Output result

App creates `reports` directory inside directory provided in `--sql-script`.

App creates:
1) CSV report for each sql file
2) Overall merged report in `report.txt`

Output looks like this:

```
./scripts
├── reports
│   └── my-local-env
│       ├── report.txt
│       ├── select-customers.sql.csv
│       └── select-products.sql.csv
├── select-customers.sql
└── select-products.sql
```

## Examples of report files you can find in `_examples` directory.
