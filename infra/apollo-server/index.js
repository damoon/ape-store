const { ApolloServer } = require('apollo-server');
const { ApolloGateway } = require("@apollo/gateway");

const delay = ms => new Promise(resolve => setTimeout(resolve, ms))
delay(1000) /// waiting 1 second.

const gateway = new ApolloGateway({
    serviceList: [
        { name: 'accounts', url: 'http://accounts:8080/query' },
        { name: 'products', url: 'http://products:8080/query' },
        { name: 'reviews',  url: 'http://reviews:8080/query' }
    ],
});

const server = new ApolloServer({
    gateway,
    subscriptions: false,
});

server.listen().then(({ url }) => {
    console.log(`🚀 Server ready at ${url}`);
});
