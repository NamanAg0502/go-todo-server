import type { Metadata } from 'next';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { IBM_Plex_Sans, Inter } from 'next/font/google';
import './globals.css';
import RootProvider from './provider';
import { Toaster } from '@/components/ui/sonner';

const inter = IBM_Plex_Sans({
  subsets: ['latin'],
  weight: ['200', '300', '400', '500', '600', '700'],
});

export const metadata: Metadata = {
  title: 'Create Next App',
  description: 'Generated by create next app',
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <RootProvider>
        <body className={inter.className}>
          <Toaster closeButton />
          {children}
        </body>
      </RootProvider>
    </html>
  );
}
