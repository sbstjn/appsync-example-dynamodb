# AppSync Example w/ DynamoDB

[![MIT License](https://badgen.now.sh/badge/License/MIT/blue)](http://github.com/sbstjn/appsync-example-dynamodb/blob/master/LICENSE.md)

> Deploy a GraphQL API using AWS AppSync, Serverless Application Model, and DynamoDB. \
> Based on [appsync-example-lambda](https://github.com/superluminar-io/appsync-example-lambda/) and [appsync-resolvers](https://sbstjn.com/serverless-graphql-with-appsync-and-lambda.html).

## Schema

```graphql
type Person {
	id: ID!
	name: String!
	age: Int!
	birthday: String!

	friends: [Person!]!
}

type Query {
	people: [Person!]!
	person(id: ID!): Person
}

type Mutation {
	personCreate(name: String!, birthday: String!): Person!
	personFriendsAdd(person: ID, friend: ID!): Person!
}

schema {
	query: Query
	mutation: Mutation
}
```

## Usage

### Deployment

```bash
# Create S3 Bucket for CloudFormation Artifacts
$ > AWS_PROFILE=your-profile-name \
    make configure

# Build, Package, and Deploy the CloudFormation Stack
$ > AWS_PROFILE=your-profile-name \
    make build package deploy
```

### API Access

```bash
# Print GraphQL API Endoint
$ > AWS_PROFILE=your-profile-name \
    make outputs-GraphQL

https://tdk6mhrty7ii.appsync-api.eu-central-1.amazonaws.com/graphql

# Print AppSync API Key
$ > AWS_PROFILE=your-profile-name \
    make outputs-APIKey

da2-1jdf4nmbwpsdr4vfxcxfza
```

### Example

#### Query

```bash
$ > curl \
    -XPOST https://tdk6mhrty7ii.appsync-api.eu-central-1.amazonaws.com/graphql \
    -H "Content-Type:application/graphql" \
    -H "x-api-key:da2-1jdf4nmbwpsdr4vfxcxfza" \
    -d '{ "query": "query { people { name } }" }' | jq
```

#### Mutation

```bash
$ > curl \
    -XPOST https://tdk6mhrty7ii.appsync-api.eu-central-1.amazonaws.com/graphql \
    -H "Content-Type:application/graphql" \
    -H "x-api-key:da2-1jdf4nmbwpsdr4vfxcxfza" \
    -d '{ "query": "mutation { personCreate(name:\"Gregory Valdes\", birthday:\"1975-10-04T00:00:00Z\") { id } }" }' | jq

$ > curl \
    -XPOST https://tdk6mhrty7ii.appsync-api.eu-central-1.amazonaws.com/graphql \
    -H "Content-Type:application/graphql" \
    -H "x-api-key:da2-1jdf4nmbwpsdr4vfxcxfza" \
    -d '{ "query": "mutation { personCreate(name:\"Alejandro Manno\", birthday:\"1962-03-23T00:00:00Z\") { id } }" }' | jq

$ > curl \
    -XPOST https://tdk6mhrty7ii.appsync-api.eu-central-1.amazonaws.com/graphql \
    -H "Content-Type:application/graphql" \
    -H "x-api-key:da2-1jdf4nmbwpsdr4vfxcxfza" \
    -d '{ "query": "mutation { personFriendsAdd( person: "b5b2d08e …", friend: "467551bb …" ) { id friends { id } } }" }' | jq
```

## Resolvers

* `Query.people`
* `Query.person`
* `Field.person.age (Lambda)`
* `Field.person.friends`
* `Mutation.personCreate`
* `Mutation.personFriendsAdd`

## License

Feel free to use the code, it's released using the [MIT license](LICENSE.md).

## Contribution

You are welcome to contribute to this project! 😘 

To make sure you have a pleasant experience, please read the [code of conduct](CODE_OF_CONDUCT.md). It outlines core values and beliefs and will make working together a happier experience.
