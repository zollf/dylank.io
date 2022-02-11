import { ApolloClient, DefaultOptions, InMemoryCache } from '@apollo/client';

const defaultOptions: DefaultOptions = {
  watchQuery: {
    fetchPolicy: 'no-cache',
  },
  query: {
    fetchPolicy: 'no-cache',
  },
};

const client = new ApolloClient({
  uri:
    typeof window !== 'undefined'
      ? `${window.location.origin}/api/graphql`
      : `http://${process.env.NGINX_HOST || 'localhost'}/api/graphql`,
  cache: new InMemoryCache(),
  defaultOptions: defaultOptions,
});

export default client;
