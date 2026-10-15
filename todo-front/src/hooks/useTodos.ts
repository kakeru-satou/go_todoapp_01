import { useState, useCallback } from "react";
import axios from "axios";
import Cookies from "js-cookie";

export type TodoType = {
  id: number;
  task: string;
  isDone: boolean;
};

export function useTodos() {
  const [todos, setTodos] = useState<TodoType[]>([]);
  const [error, setError] = useState("");

  const getHeaders = () => ({
    headers: { Authorization: `Bearer ${Cookies.get("token")}` },
  });

  const fetchTodos = useCallback(async () => {
    try {
      const res = await axios.get("http://localhost:8080/api/todoList", getHeaders());
      setTodos(res.data);
    } catch (err) {
      if (axios.isAxiosError(err)) setError("データの取得に失敗しました");
    }
  }, []);

  const fetchTodo = useCallback(async (id: string) => {
    try {
      const res = await axios.get(`http://localhost:8080/api/todoList/${id}`, getHeaders());
      return res.data;
    } catch (err) {
      if (axios.isAxiosError(err)) setError("データの取得に失敗しました");
      return null;
    }
  }, []);

  const addTodo = async (task: string) => {
    try {
      await axios.post("http://localhost:8080/api/todoList", { task }, getHeaders());
      fetchTodos();
    } catch (err) {
      if (axios.isAxiosError(err)) setError("Todoの追加に失敗しました");
    }
  };

  const deleteTodo = async (id: number) => {
    if (!window.confirm("本当に削除しますか？")) return;
    try {
      await axios.delete(`http://localhost:8080/api/todoList/${id}`, getHeaders());
      setTodos((prev) => prev.filter((todo) => todo.id !== id));
    } catch (err) {
      if (axios.isAxiosError(err)) setError("Todoの削除に失敗しました");
    }
  };

  const updateTodo = async (id: number, data: Partial<Omit<TodoType, "id">>) => {
    try {
      await axios.patch(`http://localhost:8080/api/todoList/${id}`, data, getHeaders());
      fetchTodos();
    } catch (err) {
      if (axios.isAxiosError(err)) setError("Todoの更新に失敗しました");
    }
  };

  return { todos, error, fetchTodos, fetchTodo, addTodo, deleteTodo, updateTodo };
}
