'use client';

import { ApolloLink, createHttpLink } from '@apollo/client';
import {
  NextSSRApolloClient,
  ApolloNextAppProvider,
  NextSSRInMemoryCache,
  SSRMultipartLink
} from '@apollo/experimental-nextjs-app-support/ssr';
import { setContext } from '@apollo/link-context';

const httpLink = createHttpLink({
  uri: process.env.NEXT_PUBLIC_API_URL + '/api/graphql'
});

const authLink = setContext(async (_, { headers }) => {
  let accessToken;
  if (typeof window !== 'undefined') {
    accessToken = localStorage.getItem('accessToken');
  }

  return {
    headers: {
      ...headers,
      authorization: accessToken ? `Bearer ${accessToken}` : ''
    }
  };
});

function makeClient() {
  return new NextSSRApolloClient({
    cache: new NextSSRInMemoryCache(),
    link:
      typeof window === 'undefined'
        ? ApolloLink.from([
            new SSRMultipartLink({
              stripDefer: true
            }),
            authLink.concat(httpLink)
          ])
        : authLink.concat(httpLink)
  });
}

export function ApolloWrapper({ children }: React.PropsWithChildren) {
  return <ApolloNextAppProvider makeClient={makeClient}>{children}</ApolloNextAppProvider>;
}
