// src/pages/NewTodo.tsx
import { useState } from "react";
import { useNavigate } from "react-router-dom";
import axios from "axios";

export function NewTodo() {
  const navigate = useNavigate();
  const [task, setTask] = useState("");
  const [error, setError] = useState("");

  // フォームが送信（Submit）された時の処理
  const handleSubmit = async (e: React.SubmitEvent<HTMLFormElement>) => {
    // ⚠️ エラーの先回り：画面リロードを防ぐための1行
    e.preventDefault();
    setError("");

    const token = localStorage.getItem("token");
    if (!token) {
      navigate("/login");
      return;
    }

    try {
      await axios.post(
        "http://localhost:8080/api/todoList",
        { task },
        { headers: { Authorization: `Bearer ${token}` } },
      );

      // 成功したら一覧画面に戻る
      navigate("/todos");
    } catch (err: unknown) {
      if (axios.isAxiosError(err)) {
        setError(err.response?.data?.message || "Todoの追加に失敗しました");
      } else {
        setError("予期せぬエラーが発生しました");
      }
    }
  };

  return (
    <div style={{ padding: "20px", maxWidth: "400px", margin: "0 auto" }}>
      <h2>✨ 新規Todo追加</h2>
      {error && <p style={{ color: "red", fontWeight: "bold" }}>{error}</p>}

      <form
        onSubmit={handleSubmit}
        style={{ display: "flex", flexDirection: "column", gap: "15px" }}
      >
        <div>
          <label>タイトル</label>
          <br />
          <input
            type="text"
            value={task}
            onChange={(e) => setTask(e.target.value)}
            required
            style={{ width: "100%", padding: "8px" }}
          />
        </div>

        <button
          type="submit"
          style={{ padding: "10px", backgroundColor: "blue", color: "white", fontSize: "16px" }}
        >
          追加する
        </button>
      </form>

      <div style={{ marginTop: "15px", textAlign: "center" }}>
        <button
          onClick={() => navigate("/todos")}
          style={{
            border: "none",
            background: "none",
            color: "blue",
            cursor: "pointer",
            textDecoration: "underline",
          }}
        >
          キャンセルして戻る
        </button>
      </div>
    </div>
  );
}
