import React, { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { useAuth } from "../hooks/useAuth";
import { useTodos } from "../hooks/useTodos";

export function EditTodo() {
  const navigate = useNavigate();
  const { id } = useParams<{ id: string }>();

  const [task, setTask] = useState("");
  const { isAuthenticated } = useAuth();
  const { error, fetchTodo, updateTodo } = useTodos();

  // const fetchSingleTodo = useCallback(async () => {
  //   try {
  //     if (!isAuthenticated) return navigate("/login");

  //     const response = fetchTodo(id);

  //     setTask(response.data.task);
  //   } catch (err: unknown) {
  //     if (axios.isAxiosError(err)) {
  //       setError(err.response?.data?.message || "データの取得に失敗しました");
  //     } else {
  //       setError("予期せぬエラーが発生しました");
  //     }
  //   }
  // }, [id, navigate]);

  useEffect(() => {
    const fetchSingleTodo = async () => {
      if (!isAuthenticated) return navigate("/login");
      if (id == null) return;

      const response = await fetchTodo(id);
      if (response == null) return;
      setTask(response.task);
    };

    fetchSingleTodo();
  }, [isAuthenticated, id, fetchTodo, setTask, navigate]);

  const handleUpdate = async (e: React.SubmitEvent<HTMLFormElement>) => {
    e.preventDefault();
    // setError("");

    if (!isAuthenticated) return navigate("/login");

    if (id == null) {
      return;
    }
    await updateTodo(Number(id), { task });
    alert("更新しました！");
    navigate("/todos");
  };

  return (
    <div style={{ padding: "20px", maxWidth: "400px", margin: "0 auto" }}>
      <h2>✏️ Todo編集</h2>
      {error && <p style={{ color: "red", fontWeight: "bold" }}>{error}</p>}

      <form
        onSubmit={handleUpdate}
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
          style={{
            padding: "10px",
            backgroundColor: "#4CAF50",
            color: "white",
            fontSize: "16px",
            border: "none",
            cursor: "pointer",
          }}
        >
          更新する
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
