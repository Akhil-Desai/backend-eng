package main

import (
    "fmt"
    "gql-cost-profiler/analyzer"
    "gql-cost-profiler/analyzer/types"
)

func main() {
    SchemaStr := `
        schema {
            query: Query
        }
        type Query {
            user(id: ID!): User
            posts(first: Int): [Post]
        }
        type User {
            id: ID!
            name: String!
            email: String
            posts(limit: Int): [Post]
        }
        type Post {
            id: ID!
            title: String!
            content: String
            author: User
        }
    `

    QueryStr := `
        query TestUserAndPosts {
            user(id: "1") {
                id
                name
                email
                posts(limit: 5) {
                    id
                    title
                }
            }
            posts(first: 10) {
                id
                title
            }
        }
    `

    costCfg := types.CostConfig{
        "Query": {
            "user": types.FieldCost{Base: 2, PerItemArg: "", PerItemCost: 0},
            "posts": types.FieldCost{Base: 1, PerItemArg: "first", PerItemCost: 0.5},
        },
        "User": {
            "id": types.FieldCost{Base: 0.2, PerItemArg: "", PerItemCost: 0},
            "name": types.FieldCost{Base: 0.2, PerItemArg: "", PerItemCost: 0},
            "email": types.FieldCost{Base: 0.2, PerItemArg: "", PerItemCost: 0},
            "posts": types.FieldCost{Base: 1, PerItemArg: "limit", PerItemCost: 0.3},
        },
        "Post": {
            "id": types.FieldCost{Base: 0.1, PerItemArg: "", PerItemCost: 0},
            "title": types.FieldCost{Base: 0.1, PerItemArg: "", PerItemCost: 0},
            "content": types.FieldCost{Base: 0.1, PerItemArg: "", PerItemCost: 0},
            "author": types.FieldCost{Base: 0.5, PerItemArg: "", PerItemCost: 0},
        },
    }

    cost, err := analyzer.ProfileGQLQueryWithChain(SchemaStr, QueryStr, costCfg)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    fmt.Printf("Cost: %f\n", cost)
}
// ...existing code...
