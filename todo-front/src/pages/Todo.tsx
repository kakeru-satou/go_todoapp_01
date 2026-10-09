// src/pages/Todo.tsx
import { useNavigate, Link } from "react-router-dom";
import { useEffect, useState, useCallback } from "react";
import axios from "axios";

type TodoType = {
  id: number;
  task: string;
  isDone: boolean;
};

export function Todo() {
  const navigate = useNavigate();
  const [todos, setTodos] = useState<TodoType[]>([]);
  const [error, setError] = useState<string>("");

  const fetchTodos = useCallback(async () => {
    try {
      const token = localStorage.getItem("token");

      if (!token) {
        navigate("/login");
        return;
      }

      const response = await axios.get("http://localhost:8080/api/todoList", {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      setTodos(response.data);
    } catch (err: unknown) {
      if (axios.isAxiosError(err)) {
        setError(err.response?.data?.message || "Todoの取得に失敗しました");
      } else {
        setError("予期せぬエラーが発生しました");
      }
    }
  }, [navigate]);

  useEffect(() => {
    // ⚠️ エラーの先回り：最新のLintルールの誤検知を防ぐコメント
    // eslint-disable-next-line react-hooks/set-state-in-effect
    fetchTodos();
  }, [fetchTodos]);

  // ログアウト処理
  const handleLogout = () => {
    localStorage.removeItem("token");
    navigate("/login");
  };

  return (
    <div style={{ padding: "20px", maxWidth: "600px", margin: "0 auto" }}>
      <div style={{ display: "flex", justifyContent: "space-between", alignItems: "center" }}>
        <h2>📋 Todo一覧</h2>
        <button
          onClick={handleLogout}
          style={{ padding: "5px 10px", backgroundColor: "gray", color: "white" }}
        >
          ログアウト
        </button>
      </div>

      <div style={{ marginBottom: "20px" }}>
        <Link to="/todos/new">
          <button
            style={{ padding: "10px", backgroundColor: "blue", color: "white", width: "100%" }}
          >
            + 新規Todo追加
          </button>
        </Link>
      </div>

      {error && <p style={{ color: "red", fontWeight: "bold" }}>{error}</p>}

      <ul style={{ listStyle: "none", padding: 0 }}>
        {todos.map((todo) => (
          <li
            key={todo.id}
            style={{
              border: "1px solid #ccc",
              margin: "10px 0",
              padding: "15px",
              borderRadius: "8px",
              backgroundColor: "white",
            }}
          >
            <input type="checkbox" checked={todo.isDone} readOnly />
            <h3 style={{ margin: "0 0 10px 0" }}>{todo.task}</h3>
            {/* 第3章でここに編集ボタンと削除ボタンを追加します */}
          </li>
        ))}
      </ul>
    </div>
  );
}
