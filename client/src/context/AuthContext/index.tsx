'use client';

import api from '@/lib/api';
import React, { createContext, useCallback, useEffect, useMemo } from 'react';
import { toast } from 'sonner';

interface AuthContextProps {
  isAuthenticated: boolean;
  user: User | null;
  setUser: (user: User | null) => void;
  signIn: (email: string, password: string) => Promise<void>;
  signOut: () => Promise<void>;
  signUp: (email: string, password: string) => Promise<void>;
  loading: boolean;
}

const AuthContext = createContext<AuthContextProps | undefined>(undefined);

export const AuthProvider = ({ children }: { children: React.ReactNode }) => {
  const [isAuthenticated, setIsAuthenticated] = React.useState(false);
  const [user, setUser] = React.useState<User | null>(null);
  const [loading, setLoading] = React.useState(true);

  useEffect(() => {
    const fetchUser = async () => {
      try {
        const fetchedUser = await getUserFromStorage();
        if (fetchedUser) {
          setIsAuthenticated(true);
          setUser(fetchedUser);
        }
      } catch (error) {
        console.error(error);
      } finally {
        setLoading(false);
      }
    };

    fetchUser();
  }, [isAuthenticated]);

  console.log(user);

  const signIn = useCallback(
    async (email: string, password: string): Promise<void> => {
      try {
        const response = await api.post('/auth/login', { email, password });
        const { data, success, message } = response.data as ApiResponse;
        if (success) {
          localStorage.setItem('access_token', data);
          setIsAuthenticated(success);
          setUser(await getUserFromStorage());
          toast.success(message);
        }
      } catch (error) {
        console.error(error);
      } finally {
        setLoading(false);
      }
    },
    []
  );

  const signOut = async (): Promise<void> => {
    localStorage.removeItem('access_token');
    setIsAuthenticated(false);
    setUser(null);
  };

  const signUp = useCallback(
    async (email: string, password: string): Promise<void> => {
      try {
        await api.post('/auth/register', { email, password });
        signIn(email, password);
      } catch (error) {
        console.error(error);
      } finally {
        setLoading(false);
      }
    },
    [signIn]
  );

  const getUserFromStorage = async (): Promise<User | null> => {
    const token = localStorage.getItem('access_token');
    if (!token) return null;
    try {
      const response = await api.get('/users/me', {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });
      return response.data;
    } catch (error) {
      localStorage.removeItem('access_token');
      return null;
    }
  };

  const contextValue = useMemo(() => {
    return {
      isAuthenticated,
      user,
      setUser,
      signIn,
      signOut,
      signUp,
      loading,
    };
  }, [isAuthenticated, signIn, signUp, user, loading]);
  return (
    <AuthContext.Provider value={contextValue}>{children}</AuthContext.Provider>
  );
};

export const useAuth = () => {
  const context = React.useContext(AuthContext);
  if (!context) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
};
