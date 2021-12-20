import { ApolloClient, InMemoryCache } from '@apollo/client';

const client = new ApolloClient({
  uri:
    typeof window !== 'undefined'
      ? `${window.location.origin}/api/graphql`
      : `http://${process.env.NGINX_HOST || 'localhost'}/api/graphql`,
  cache: new InMemoryCache(),
});

export default client;
