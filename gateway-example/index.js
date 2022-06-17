const { ApolloServer, gql } = require("apollo-server");
const { ApolloGateway, IntrospectAndCompose } = require("@apollo/gateway");
const { readFileSync } = require("fs");
import { GraphQLClient } from "graphql-request";

const port = 8081;
const gqlstr = `
  query{
    services{
      id
      name
      url
    }
  }
`;

const graphQLClient = new GraphQLClient(
  "https://localhost:8080/grahpql",
  {
    mode: "cors",
  }
);
graphQLClient
  .request(gqlstr)
  .then((data) => {
    if (data) {
      console.log(data);
    }
  })
  .catch((err) => {
    console.error(err);
  });
// Initialize an ApolloGateway instance and pass it
// the supergraph schema as a string
const gateway = new ApolloGateway({
  supergraphSdl: new IntrospectAndCompose({
    subgraphs: [
      { name: "accounts", url: "http://localhost:4000/graphql" },
      // { name: "products", url: "http://localhost:4002" },
      // ...additional subgraphs...
    ],
  }),
});

// Pass the ApolloGateway to the ApolloServer constructor
const server = new ApolloServer({
  gateway,
});

server.listen({ port }).then(({ url }) => {
  console.log(`🚀 Server ready at ${url}`);
});
