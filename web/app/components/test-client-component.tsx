'use client';

import { gql } from '@apollo/client';
import { useSuspenseQuery } from '@apollo/experimental-nextjs-app-support/ssr';

const FIND_ALL_USERS = gql`
  query {
    findAllUsers {
      id
      name
      email
    }
  }
`;

export function TestClientComponent() {
  const { data, error } = useSuspenseQuery<any>(FIND_ALL_USERS);

  if (error) return <p>Error: {error.message}</p>;

  return (
    <div>
      <h1>Client Componnet</h1>
      <ul>
        {data.findAllUsers.map((user: any) => (
          <li key={user.id}>{user.name}</li>
        ))}
      </ul>
    </div>
  );
  return <div>hi</div>;
}
