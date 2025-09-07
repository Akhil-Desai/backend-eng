package analyzer

import (
	"fmt"
	"strconv"
	"gql-cost-profiler/analyzer/types"
)


func applyCost(nodes map[string][]*types.Node, config types.CostConfig) (float32,error) {

	cost := float32(0)

	for parentType,nodes := range nodes {

		for _,node := range nodes {

			fieldCfg,ok := config[parentType][node.FieldName]
			if !ok { continue }

			baseCost := fieldCfg.Base
			cost = baseCost + cost

			perItemArg := fieldCfg.PerItemArg

			perItemCost := fieldCfg.PerItemCost

			for _,arg := range node.FieldArguments {

				if arg.Name == perItemArg {

					argVal,err := convertToFloat32(arg.Value)
					if err != nil { return -1, fmt.Errorf("error converting argument value to float32...💥") }
					cost = (argVal * perItemCost) + cost

				}

			}

		}

	}
	return cost,nil
}

func convertToFloat32(value interface{}) (float32,error) {

	switch t := value.(type) {
	case string:
		newValue, err := strconv.ParseFloat(t, 32)
		if err != nil {
			return float32(-1) , fmt.Errorf("error converting type string to float64 for value... %w 💥", err)
		}
		return float32(newValue),nil

	case int:
		newValue := float32(value.(int))
		return newValue,nil

	case float64:
		return value.(float32),nil

	default:
		return float32(-1) , fmt.Errorf("invalid type for converting to float64 for perItemCost... type was: %T 💥", t)
	}
}
