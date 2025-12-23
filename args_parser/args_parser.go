package args_parser

import "flag"

type Args struct {
	SqlScriptsPath      string
	DbConnectionsConfig string
}

func Parse() Args {

	sqlScriptsPath := flag.String("sql-scripts", "./sql", "Path where sql scripts are located. There will be placed reports in 'reports' dir")
	dbConnectionsConfigYamlPath := flag.String("connections-config", "./connections.yaml", "Path to YAML file with connections configs")
	flag.Parse()

	return Args{SqlScriptsPath: *sqlScriptsPath, DbConnectionsConfig: *dbConnectionsConfigYamlPath}
}
