'use client';

import { useAuth } from '@/context/AuthContext';
import { redirect } from 'next/navigation';

const MainLayout = ({ children }: { children: React.ReactNode }) => {
  const { isAuthenticated, loading } = useAuth();
  if (loading) {
    return null;
  }
  if (!isAuthenticated) {
    redirect('/login');
  }
  return (
    <main className="h-screen w-full flex items-center justify-center">
      {children}
    </main>
  );
};

export default MainLayout;
