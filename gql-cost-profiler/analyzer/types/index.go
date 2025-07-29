package types


type FieldCost struct {
	Base 	  	float32
	PerItemArg 	string
	PerItemCost	float32
}
type CostConfig map[string]map[string]FieldCost

//---------------

type Node struct {
	FieldName string
	FieldArguments   []*FieldArgument
}
type FieldArgument struct {
	Name string
	Value interface{}
}
