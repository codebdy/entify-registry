const { ApolloServer, gql } = require('apollo-server');
const { ApolloGateway } = require('@apollo/gateway');
const { readFileSync } = require('fs');

const supergraphSdl = readFileSync('./supergraph.graphql').toString();

// Initialize an ApolloGateway instance and pass it
// the supergraph schema as a string
const gateway = new ApolloGateway({
  supergraphSdl,
});

// Pass the ApolloGateway to the ApolloServer constructor
const server = new ApolloServer({
  gateway,
});

server.listen().then(({ url }) => {
  console.log(`🚀 Server ready at ${url}`);
});