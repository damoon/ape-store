const { ApolloServer } = require('apollo-server');
const { ApolloGateway } = require("@apollo/gateway");

const gateway = new ApolloGateway({
    serviceList: [
        { name: 'accounts', url: 'http://accounts:8080/query' }
//        { name: 'products', url: 'http://products:4002/query' },
//        { name: 'reviews', url: 'http://reviews:4003/query' }
    ],
});

const server = new ApolloServer({
    gateway,
});

server.listen().then(({ url }) => {
    console.log(`🚀 Server ready at ${url}`);
});
