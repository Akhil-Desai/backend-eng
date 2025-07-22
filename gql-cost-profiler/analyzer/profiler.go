package analyzer

import (
	"fmt"
	"gql-cost-profiler/analyzer/types"
)


func ProfileGQLQuery(query string, costcfg types.CostConfig) (float32,error) {
	schema,schemaErr := ParseGQLSchema("")  //Will import schema
	if schemaErr != nil {
		return -1,fmt.Errorf("error parsing schema...💥")
	}
	queryDoc, queryErr := ParseGQLQuery(schema, query)
	if queryErr != nil {
		return -1,fmt.Errorf("error parsing schema...💥")
	}
	queryNodes := ExtractQueryNodes(queryDoc,schema)

	cst,cstErr := applyCost(queryNodes, costcfg)
	if cstErr != nil {
		return -1, fmt.Errorf("error applying cost to query...💥")
	}

	return cst,nil
}
