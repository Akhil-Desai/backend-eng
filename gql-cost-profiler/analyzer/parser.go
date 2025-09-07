package analyzer

import (
	"fmt"
	"gql-cost-profiler/analyzer/types"
	"github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/validator"
)



func ParseGQLSchema(schemaStr string) (*ast.Schema, error) {
	schema,err := gqlparser.LoadSchema(&ast.Source{Name: "schema.graphql", Input: schemaStr})
	if err != nil {
		return nil, fmt.Errorf("failed to parse schmea: %w 💥", err)
	}
	return schema, nil
}

func ParseGQLQuery(schema *ast.Schema, queryStr string) (*ast.QueryDocument, error) {
	query, err := gqlparser.LoadQuery(schema, queryStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse query: %w 💥", err)
	}
	errs := validator.Validate(schema, query)
	if len(errs) > 0 {
		return nil, fmt.Errorf("query validaton errors: %v 💥", errs)
	}
	return query, nil
}

func ExtractQueryNodes(doc *ast.QueryDocument, schema *ast.Schema)(map[string][]*types.Node) {

	nodes := make(map[string][]*types.Node)
	visitedFragments := make(map[string]struct{})


	var walkSelections func(selections ast.SelectionSet, parentType string)
	walkSelections = func(selections ast.SelectionSet, parentType string) {

		for _,sel := range selections{
			switch s := sel.(type){

			case *ast.Field:

				fieldDef := schema.Types[parentType].Fields.ForName(s.Name)
				if fieldDef == nil {
					continue
				}

				var newFieldArgs []*types.FieldArgument
				for _,arg := range s.Arguments {
					newFieldArgs = append(newFieldArgs,
						&types.FieldArgument{
						Name: arg.Name,
						Value : arg.Value.Raw,
					})
				}
				nodes[parentType] = append(nodes[parentType], &types.Node{FieldName: s.Name, FieldArguments: newFieldArgs})

				walkSelections(s.SelectionSet, fieldDef.Type.Name())

			case *ast.FragmentSpread:

				if _, ok := visitedFragments[s.Name]; ok {
					continue
				}
				visitedFragments[s.Name] = struct{}{}
				frag := doc.Fragments.ForName(s.Name)
				if frag != nil {
					walkSelections(frag.SelectionSet, frag.TypeCondition)
				}
			case *ast.InlineFragment:
				walkSelections(s.SelectionSet, s.TypeCondition )
			}
		}
	}

	for _,op := range doc.Operations {
		var rootType string
		switch op.Operation {
		case ast.Query:
			rootType = "Query"
		case ast.Mutation:
			rootType = "Mutation"
		case ast.Subscription:
			rootType = "Subscription"

		}

		walkSelections(op.SelectionSet, rootType)
	}

	return nodes
}
