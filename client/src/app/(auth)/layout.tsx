'use client';

import { useAuth } from '@/context/AuthContext';
import { redirect } from 'next/navigation';

const AuthLayout = ({ children }: { children: React.ReactNode }) => {
  const { isAuthenticated, loading } = useAuth();
  if (loading) {
    return null;
  }
  if (isAuthenticated) {
    redirect('/');
  }
  return (
    <main className="h-screen w-full flex items-center justify-center">
      {children}
    </main>
  );
};

export default AuthLayout;
