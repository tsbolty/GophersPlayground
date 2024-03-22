'use server';

import { gql } from '@apollo/client';
import { getClient } from '../lib/fetcher';
import { TestClientComponent } from '../components/test-client-component';

const FIND_ALL_USERS = gql`
  query {
    findAllUsers {
      id
      name
      email
    }
  }
`;

export default async function DashboardPage() {
  const { data, error, loading } = await getClient().query({ query: FIND_ALL_USERS });

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error: {error.message}</p>;
  return (
    <div>
      <h1>Dashboard Server Component</h1>
      <ul>
        {data.findAllUsers.map((user: any) => (
          <li key={user.id}>{user.name}</li>
        ))}
      </ul>
      <TestClientComponent />
    </div>
  );
}
