package cli


import (
	"flag"
	"fmt"
	"os"
	"gql-cost-profiler/analyzer"
)

func CostCLI() (float32,error) {

	costConfigPath := flag.String("cost-cfg", "", "Path to the cost config file")
	queryStringPath := flag.String("query", "", "Path to the query file")

	flag.Parse()


	_,err := os.ReadFile(*costConfigPath) //going to have to figure out how to transform this into our type // maybe enforce it as a JSON
	if err != nil {
		return -1,fmt.Errorf("error reading contents from the cost config file...💥")
	}

	queryString,err := os.ReadFile(*queryStringPath)
	if err != nil {
		return -1,fmt.Errorf("error reading contents from the query file...💥")
	}

	cost,err := analyzer.ProfileGQLQuery(string(queryString),nil)

	if err != nil {
		return -1,fmt.Errorf("error profiling cost of query...💥")
	}

	return cost,nil
}
