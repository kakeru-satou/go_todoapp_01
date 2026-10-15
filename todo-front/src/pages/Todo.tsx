import { useNavigate, Link } from "react-router-dom";
import { useEffect } from "react";
import { useAuth } from "../hooks/useAuth";
import { useTodos } from "../hooks/useTodos";

export function Todo() {
  const navigate = useNavigate();
  // const [todos, setTodos] = useState<TodoType[]>([]);
  // const [error, setError] = useState<string>("");
  const { logout, isAuthenticated } = useAuth();
  const { todos, error, fetchTodos, deleteTodo } = useTodos();

  // const fetchTodos = useCallback(async () => {
  //   try {
  //     const token = Cookies.get("token");

  //     if (!token) {
  //       navigate("/login");
  //       return;
  //     }

  //     const response = await axios.get("http://localhost:8080/api/todoList", {
  //       headers: {
  //         Authorization: `Bearer ${token}`,
  //       },
  //     });

  //     setTodos(response.data);
  //   } catch (err: unknown) {
  //     if (axios.isAxiosError(err)) {
  //       setError(err.response?.data?.message || "Todoの取得に失敗しました");
  //     } else {
  //       setError("予期せぬエラーが発生しました");
  //     }
  //   }
  // }, [navigate]);

  useEffect(() => {
    if (isAuthenticated) {
      fetchTodos();
    } else {
      navigate("/login");
      return;
    }
  }, [isAuthenticated, fetchTodos, navigate]);

  // ログアウト処理
  const handleLogout = () => {
    // localStorage.removeItem("token");
    // Cookies.remove("token");
    logout();
    navigate("/login");
  };

  // const handleDelete = async (id: number) => {
  //   if (!window.confirm("本当に削除しますか？")) return;

  //   try {
  //     const token = Cookies.get("token");
  //     if (!token) {
  //       navigate("/login");
  //       return;
  //     }
  //     await axios.delete(`http://localhost:8080/api/todoList/${id}`, {
  //       headers: {
  //         Authorization: `Bearer ${token}`,
  //       },
  //     });

  //     setTodos((prev) => prev.filter((todo) => todo.id !== id));
  //   } catch (err: unknown) {
  //     if (axios.isAxiosError(err)) {
  //       setError(err.response?.data?.message || "Todoの削除に失敗しました");
  //     } else {
  //       setError("予期せぬエラーが発生しました");
  //     }
  //   }
  // };

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
            <div style={{ display: "flex", gap: "10px", justifyContent: "flex-start" }}>
              <input type="checkbox" checked={todo.isDone} readOnly />
            </div>
            <h3 style={{ margin: "0 0 10px 0" }}>{todo.task}</h3>
            <div style={{ display: "flex", gap: "10px", justifyContent: "flex-end" }}>
              <Link to={`/todos/${todo.id}/edit`}>
                <button
                  style={{
                    padding: "5px 15px",
                    backgroundColor: "#4CAF50",
                    color: "white",
                    border: "none",
                    borderRadius: "4px",
                    cursor: "pointer",
                  }}
                >
                  編集
                </button>
              </Link>

              <button
                onClick={() => deleteTodo(todo.id)}
                style={{
                  padding: "5px 15px",
                  backgroundColor: "#f44336",
                  color: "white",
                  border: "none",
                  borderRadius: "4px",
                  cursor: "pointer",
                }}
              >
                削除
              </button>
            </div>
          </li>
        ))}
      </ul>
    </div>
  );
}
