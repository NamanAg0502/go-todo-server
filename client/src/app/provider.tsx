'use client';

import { AuthProvider } from '@/context/AuthContext';
import { TodoProvider } from '@/context/TodoContext';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';

export default function RootProvider({
  children,
}: {
  children: React.ReactNode;
}) {
  const queryClient = new QueryClient();

  return (
    <QueryClientProvider client={queryClient}>
      <AuthProvider>
        <TodoProvider>{children}</TodoProvider>
      </AuthProvider>
    </QueryClientProvider>
  );
}
