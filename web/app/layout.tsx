import type { Metadata } from 'next';
import { ApolloWrapper } from './lib/apollo-client';

export const metadata: Metadata = {
  title: 'Cool Next.js App',
  description: 'Used for testing Next.js'
};

export default function RootLayout({
  children
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang='en'>
      <body>
        <ApolloWrapper>{children}</ApolloWrapper>
      </body>
    </html>
  );
}
