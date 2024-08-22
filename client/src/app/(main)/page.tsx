'use client';

import Link from 'next/link';
import {
  DropdownMenu,
  DropdownMenuTrigger,
  DropdownMenuContent,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuItem,
} from '@/components/ui/dropdown-menu';
import { Button } from '@/components/ui/button';
import { Avatar, AvatarImage, AvatarFallback } from '@/components/ui/avatar';
import { Card, CardContent } from '@/components/ui/card';
import {
  Table,
  TableHeader,
  TableRow,
  TableHead,
  TableBody,
  TableCell,
} from '@/components/ui/table';
import { Checkbox } from '@/components/ui/checkbox';
import { Badge } from '@/components/ui/badge';
import {
  CheckIcon,
  ClockIcon,
  FilePenIcon,
  ListIcon,
  TrashIcon,
} from 'lucide-react';
import { useTodos } from '@/context/TodoContext';
import { useEffect } from 'react';
import Loading from '../loading';

export default function MainPage() {
  const { fetchTodos, todos, loading, addTodo, updateTodo, deleteTodo } =
    useTodos();
  useEffect(() => {
    fetchTodos();
  }, [fetchTodos]);

  if (loading || !todos) return <Loading />;

  console.log(todos);

  return (
    <div className="flex min-h-screen flex-col bg-background w-full">
      <header className="sticky top-0 z-30 flex h-14 items-center gap-4 border-b bg-background px-4 sm:static sm:h-auto sm:border-0 sm:bg-transparent sm:px-6">
        <Link
          href="#"
          className="flex items-center gap-2 font-semibold"
          prefetch={false}
        >
          <CheckIcon className="h-6 w-6" />
          <span className="text-lg">Todo App</span>
        </Link>
        <div className="flex items-center justify-end gap-2 max-w-5xl mx-auto w-full">
          <DropdownMenu>
            <DropdownMenuTrigger asChild>
              <Button
                variant="ghost"
                size="icon"
                className="overflow-hidden rounded-full"
              >
                <Avatar className="h-8 w-8 border">
                  <AvatarImage src="/placeholder-user.jpg" alt="@shadcn" />
                  <AvatarFallback>JP</AvatarFallback>
                </Avatar>
              </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent align="end">
              <DropdownMenuLabel>Jared Palmer</DropdownMenuLabel>
              <DropdownMenuSeparator />
              <DropdownMenuItem>Profile</DropdownMenuItem>
              <DropdownMenuItem>Settings</DropdownMenuItem>
              <DropdownMenuSeparator />
              <DropdownMenuItem>Logout</DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>
        </div>
      </header>
      <div className="flex flex-1">
        <div className="hidden border-r bg-muted/40 sm:block">
          <div className="flex h-full max-h-screen flex-col gap-2">
            <div className="flex h-[60px] items-center border-b px-6">
              <span className="text-lg font-semibold">Menu</span>
            </div>
            <div className="flex-1 overflow-auto py-2">
              <nav className="grid items-start px-4 text-sm font-medium">
                <Link
                  href="#"
                  className="flex items-center gap-3 rounded-lg px-3 py-2 text-muted-foreground transition-all hover:text-primary"
                  prefetch={false}
                >
                  <ListIcon className="h-4 w-4" />
                  All Tasks
                </Link>
                <Link
                  href="#"
                  className="flex items-center gap-3 rounded-lg px-3 py-2 text-muted-foreground transition-all hover:text-primary"
                  prefetch={false}
                >
                  <CheckIcon className="h-4 w-4" />
                  Completed
                </Link>
                <Link
                  href="#"
                  className="flex items-center gap-3 rounded-lg bg-muted px-3 py-2 text-primary transition-all hover:text-primary"
                  prefetch={false}
                >
                  <ClockIcon className="h-4 w-4" />
                  Pending
                </Link>
              </nav>
            </div>
          </div>
        </div>
        <div className="flex flex-1 flex-col max-w-5xl mx-auto">
          <main className="flex-1 p-4 sm:p-6">
            <div className="grid gap-4">
              <div className="flex items-center justify-between">
                <h1 className="text-2xl font-bold">Tasks</h1>
                <Button size="sm">Add Task</Button>
              </div>
              <Card>
                <CardContent>
                  <Table>
                    <TableHeader>
                      <TableRow>
                        <TableHead>Task</TableHead>
                        <TableHead>Status</TableHead>
                        <TableHead>Due Date</TableHead>
                        <TableHead>Actions</TableHead>
                      </TableRow>
                    </TableHeader>
                    <TableBody>
                      {todos.map((task) => (
                        <TableRow key={task.id}>
                          <TableCell>
                            <div className="flex items-center gap-3">
                              <Checkbox />
                              <div className="font-medium">{task.title}</div>
                            </div>
                          </TableCell>
                          <TableCell>
                            <Badge variant="outline">
                              {!task.isCompleted ? 'pending' : 'completed'}
                            </Badge>
                          </TableCell>
                          <TableCell>
                            {new Date(task.createdAt).toLocaleString()}
                          </TableCell>
                          <TableCell>
                            <div className="flex items-center gap-2">
                              <Button variant="ghost" size="icon">
                                <FilePenIcon className="h-4 w-4" />
                                <span className="sr-only">Edit</span>
                              </Button>
                              <Button variant="ghost" size="icon">
                                <TrashIcon className="h-4 w-4" />
                                <span className="sr-only">Delete</span>
                              </Button>
                            </div>
                          </TableCell>
                        </TableRow>
                      ))}
                    </TableBody>
                  </Table>
                </CardContent>
              </Card>
            </div>
          </main>
        </div>
      </div>
    </div>
  );
}
