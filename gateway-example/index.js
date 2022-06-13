const { ApolloServer, gql } = require("apollo-server");
const { ApolloGateway } = require("@apollo/gateway");
const { readFileSync } = require("fs");

// Initialize an ApolloGateway instance and pass it
// the supergraph schema as a string
const gateway = new ApolloGateway({
  serviceList: [
    { name: "people", url: "http://localhost:4001" },
    { name: "films", url: "http://localhost:4002" },
  ],
});

// Pass the ApolloGateway to the ApolloServer constructor
const server = new ApolloServer({
  gateway,
});

server.listen({ port: 8081 }).then(({ url }) => {
  console.log(`🚀 Server ready at ${url}`);
});
