package analyzer

import (
	"errors"
	"fmt"
	"gql-cost-profiler/analyzer/types"
	"github.com/vektah/gqlparser/v2/ast"
)

type ProfileContext struct {
	SchemaString string
	QueryString string
	Schema *ast.Schema
	QueryDoc *ast.QueryDocument
	QueryNodes map[string][]*types.Node
	Cost float32
	CostCfg types.CostConfig
}

type Handler interface {
	SetNext(Handler)
	Handle (ctx *ProfileContext) error
}

type BaseHandler struct {
	next Handler
}

func (b *BaseHandler) SetNext(f Handler){
	b.next = f
}

func (b *BaseHandler) CallNext(ctx *ProfileContext) error {
	if b.next != nil {
		b.next.Handle(ctx)
	}

	return errors.New("next is nil...💥")
}

//-----------------

type SchemaParser struct { BaseHandler }
func (h *SchemaParser) Handle(ctx *ProfileContext) error {
	schema,err := ParseGQLSchema(ctx.SchemaString)
	if err != nil {
		return err
	}
	ctx.Schema = schema
	h.CallNext(ctx)
	return nil
}

type QueryParser struct { BaseHandler }
func (h *QueryParser) Handle(ctx *ProfileContext) error {

	queryDoc,err := ParseGQLQuery(ctx.Schema, ctx.QueryString)
	if err != nil {
		return err
	}
	ctx.QueryDoc = queryDoc
	h.CallNext(ctx)
	return nil
}

type NodeExtractor struct{ BaseHandler }
func (h *NodeExtractor) Handle(ctx *ProfileContext) error {
    if ctx.QueryDoc == nil {
        return fmt.Errorf("query doc missing...💥");
    }
    nodes := ExtractQueryNodes(ctx.QueryDoc, ctx.Schema)
    ctx.QueryNodes = nodes
    return h.CallNext(ctx)
}

type CostApplier struct{ BaseHandler }
func (h *CostApplier) Handle(ctx *ProfileContext) error {
    if ctx.QueryNodes == nil {
        return fmt.Errorf("query nodes missing")
    }
    cost,_ := applyCost(ctx.QueryNodes,ctx.CostCfg)
    ctx.Cost = cost
    return h.CallNext(ctx)
}

func ProfileGQLQueryWithChain(schemaStr string, query string, costCfg types.CostConfig) (float32,error) {
	ctx := &ProfileContext{
		SchemaString: schemaStr,
		QueryString:  query,
		CostCfg: costCfg,
	}

	schemaParser := &SchemaParser{}
    queryParser := &QueryParser{}
    nodeExtractor := &NodeExtractor{}
    costApplier := &CostApplier{}

	schemaParser.SetNext(queryParser)
	queryParser.SetNext(nodeExtractor)
    nodeExtractor.SetNext(costApplier)

	err := schemaParser.Handle(ctx)
	if err != nil {
		return float32(-1), err
	}

	return ctx.Cost, nil

}
