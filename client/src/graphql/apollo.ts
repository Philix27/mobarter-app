import { ApolloClient, InMemoryCache, HttpLink, ApolloLink } from '@apollo/client';
import { getOperationAST } from 'graphql';

const cache = new InMemoryCache({
	addTypename: true
});

// const wsLink = new WebSocketLink({
// 	uri: 'wss://localhost:8080/graphql/',
// 	options: {
// 		lazy: true,
// 		reconnect: true
// 	}
// });

const httpLink = new HttpLink({
	uri: 'http://localhost:8080/graphql/'
});

const link = ApolloLink.split(
	(op: any) => {
		// check if it is a subscription
		const operationAST = getOperationAST(op.query, op.operationName);
		return !!operationAST && operationAST.operation === 'subscription';
	},
	// wsLink,
	httpLink
);

export default new ApolloClient({
	cache,
	link,
	connectToDevTools: true
});
