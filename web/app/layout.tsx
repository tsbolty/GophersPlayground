import type { Metadata } from "next";

export const metadata: Metadata = {
  title: "Cool Next.js App",
  description: "Used for testing Next.js",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  );
}
