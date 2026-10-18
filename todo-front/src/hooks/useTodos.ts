import axios from "axios";
import Cookies from "js-cookie";

import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { useAuth } from "./useAuth";

export type TodoType = {
  id: number;
  task: string;
  isDone: boolean;
};

type UpdateTodoInput = Partial<TodoType> & { id: number };

const getHeaders = () => ({
  headers: { Authorization: `Bearer ${Cookies.get("token")}` },
});

export function useTodos() {
  const queryClient = useQueryClient();
  const { isAuthenticated } = useAuth();

  const {
    data: todos = [],
    isLoading,
    isError,
  } = useQuery({
    queryKey: ["todos"],
    queryFn: async () => {
      const res = await axios.get("http://localhost:8080/api/todoList", getHeaders());
      return res.data as TodoType[];
    },
    enabled: isAuthenticated,
  });

  // const fetchTodos = useCallback(async () => {
  //   try {
  //     const res = await axios.get("http://localhost:8080/api/todoList", getHeaders());
  //     setTodos(res.data);
  //   } catch (err) {
  //     if (axios.isAxiosError(err)) setError("データの取得に失敗しました");
  //   }
  // }, []);

  // const fetchTodo = useCallback(async (id: string) => {
  //   try {
  //     const res = await axios.get(`http://localhost:8080/api/todoList/${id}`, getHeaders());
  //     return res.data;
  //   } catch (err) {
  //     if (axios.isAxiosError(err)) setError("データの取得に失敗しました");
  //     return null;
  //   }
  // }, []);

  const addTodoMutation = useMutation({
    mutationFn: async (newTodo: { task: string }) => {
      await axios.post("http://localhost:8080/api/todoList", newTodo, getHeaders());
    },
    onSuccess: () => queryClient.invalidateQueries({ queryKey: ["todos"] }),
  });

  // const addTodo = async (task: string) => {
  //   try {
  //     await axios.post("http://localhost:8080/api/todoList", { task }, getHeaders());
  //     fetchTodos();
  //   } catch (err) {
  //     if (axios.isAxiosError(err)) setError("Todoの追加に失敗しました");
  //   }
  // };

  const deleteTodoMutation = useMutation({
    mutationFn: async (id: number) => {
      await axios.delete(`http://localhost:8080/api/todoList/${id}`, getHeaders());
    },
    onSuccess: () => queryClient.invalidateQueries({ queryKey: ["todos"] }),
  });

  // const deleteTodo = async (id: number) => {
  //   if (!window.confirm("本当に削除しますか？")) return;
  //   try {
  //     await axios.delete(`http://localhost:8080/api/todoList/${id}`, getHeaders());
  //     setTodos((prev) => prev.filter((todo) => todo.id !== id));
  //   } catch (err) {
  //     if (axios.isAxiosError(err)) setError("Todoの削除に失敗しました");
  //   }
  // };

  const updateTodoMutation = useMutation({
    mutationFn: async (variables: UpdateTodoInput) => {
      const task = variables.task;
      const isDone = variables.isDone;
      await axios.patch(
        `http://localhost:8080/api/todoList/${variables.id}`,
        { task, isDone },
        getHeaders(),
      );
    },
    onSuccess: () => queryClient.invalidateQueries({ queryKey: ["todos"] }),
  });

  // const updateTodo = async (id: number, data: Partial<Omit<TodoType, "id">>) => {
  //   try {
  //     await axios.patch(`http://localhost:8080/api/todoList/${id}`, data, getHeaders());
  //     fetchTodos();
  //   } catch (err) {
  //     if (axios.isAxiosError(err)) setError("Todoの更新に失敗しました");
  //   }
  // };

  // return { todos, error, fetchTodos, fetchTodo, addTodo, deleteTodo, updateTodo };
  return {
    todos,
    isLoading,
    isError,
    addTodo: addTodoMutation.mutateAsync,
    deleteTodo: deleteTodoMutation.mutateAsync,
    updateTodo: updateTodoMutation.mutateAsync,
  };
}
