import type { Metadata } from 'next';
import './globals.css';

export const metadata: Metadata = {
  title: 'hello-word-20',
  description: 'Centered greeting page',
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
