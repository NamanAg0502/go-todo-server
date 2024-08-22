'use client';
import api from '@/lib/api';
import { AxiosResponse } from 'axios';
import React, {
  createContext,
  useCallback,
  useEffect,
  useMemo,
  useReducer,
} from 'react';
import { toast } from 'sonner';

interface TodoContextProps {
  fetchTodos: () => Promise<void>;
  updateTodo: (id: string) => Promise<void>;
  deleteTodo: (id: string) => Promise<void>;
  addTodo: (title: string) => Promise<void>;
  todos: Todo[] | undefined;
  loading: boolean;
}

interface TodoState {
  todos: Todo[] | undefined;
  loading: boolean;
}

type TodoAction =
  | { type: 'SET_TODOS'; payload: Todo[] }
  | { type: 'SET_LOADING'; payload: boolean };

const todoReducer = (state: TodoState, action: TodoAction): TodoState => {
  switch (action.type) {
    case 'SET_TODOS':
      return { ...state, todos: action.payload };
    case 'SET_LOADING':
      return { ...state, loading: action.payload };
    default:
      return state;
  }
};

const TodoContext = createContext<TodoContextProps | undefined>(undefined);

export const TodoProvider = ({ children }: { children: React.ReactNode }) => {
  const [state, dispatch] = useReducer(todoReducer, {
    todos: undefined,
    loading: false,
  });
  const { todos, loading } = state;

  const apiRequest = async (request: () => Promise<AxiosResponse>) => {
    try {
      dispatch({ type: 'SET_LOADING', payload: true });
      const response = await request();
      const { data, success, message } = response.data as ApiResponse;
      if (success) {
        return data;
      } else {
        toast.error(message);
        return null;
      }
    } catch (error) {
      toast.error('An error occurred');
      console.error(error);
      return null;
    } finally {
      dispatch({ type: 'SET_LOADING', payload: false });
    }
  };

  const fetchTodos = useCallback(async () => {
    const data = await apiRequest(() => api.get('/todos'));
    if (data) dispatch({ type: 'SET_TODOS', payload: data as Todo[] });
  }, []);

  const contextValue = useMemo(() => {
    return {
      fetchTodos,
      updateTodo: async (id: string) =>
        apiRequest(() => api.put(`/todos/${id}`, { isCompleted: true })),
      deleteTodo: async (id: string) =>
        apiRequest(() => api.delete(`/todos/${id}`)),
      addTodo: async (title: string, isCompleted = false) =>
        apiRequest(() => api.post('/todos', { title, isCompleted })),
      todos,
      loading,
    };
  }, [fetchTodos, todos, loading]);
  return (
    <TodoContext.Provider value={contextValue}>{children}</TodoContext.Provider>
  );
};

export const useTodos = () => {
  const context = React.useContext(TodoContext);
  if (!context) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
};
