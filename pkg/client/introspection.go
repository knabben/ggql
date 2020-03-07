package client

import "github.com/graphql-go/graphql"

var (
	introspectionQuery = `
query IntrospectionQuery {
  __schema {
    queryType {
      name
    }
    mutationType {
      name
    }
    subscriptionType {
      name
    }
    types {
      ...FullType
    }
    directives {
      name
      description
      locations
      args {
        ...InputValue
      }
    }
  }
}`

	fragmentFullType = `
fragment FullType on __Type {
  kind
  name
  description
  fields(includeDeprecated: true) {
    name
    description
    args {
      ...InputValue
    }
    type {
      ...TypeRef
    }
    isDeprecated
    deprecationReason
  }
  inputFields {
    ...InputValue
  }
  interfaces {
    ...TypeRef
  }
  enumValues(includeDeprecated: true) {
    name
    description
    isDeprecated
    deprecationReason
  }
  possibleTypes {
    ...TypeRef
  }
}`

	fragmentValue = `
fragment InputValue on __InputValue {
  name
  description
  type {
    ...TypeRef
  }
  defaultValue
}
`
	fragmentTypeRef = `
fragment TypeRef on __Type {
  kind
  name
  ofType {
    kind
    name
    ofType {
      kind
      name
      ofType {
        kind
        name
        ofType {
          kind
          name
          ofType {
            kind
            name
            ofType {
              kind
              name
              ofType {
                kind
                name
              }
            }
          }
        }
      }
    }
  }
}
`
)

type FieldDefinition struct {
	Name              string                 `json:"name"`
	Description       string                 `json:"description"`
	Type              map[string]interface{} `json:"type"`
	Args              []*graphql.Argument    `json:"args"`
	DeprecationReason string                 `json:"deprecationReason"`
}

type Schema struct {
	Schema struct {
		QueryType struct {
			Name string `json:"name"`
		} `json:"queryType"`
		Types []struct {
			Kind   string            `json:"kind"`
			Name   string            `json:"name"`
			Fields []FieldDefinition `json:"fields"`
		} `json:"types"`
	} `json:"__schema"`
}

func BuildIntrospectionQuery() string {
	return introspectionQuery + fragmentFullType + fragmentValue + fragmentTypeRef
}
