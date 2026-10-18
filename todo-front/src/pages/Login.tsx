// src/pages/Login.tsx
import { useState } from "react";
import { useNavigate, Link } from "react-router-dom";
import axios from "axios";
import { useAuth } from "../hooks/useAuth";

export function Login() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const { login } = useAuth();

  const navigate = useNavigate();
  const handleLogin = async (e: React.SubmitEvent<HTMLFormElement>) => {
    e.preventDefault();
    setError("");

    try {
      const response = await axios.post("http://localhost:8080/api/auth/signin", {
        email,
        password,
      });

      // localStorage.setItem("token", response.data.token);
      // Cookies.set("token", response.data.token, { expires: 7 });
      login(response.data.token);
      alert("ログイン成功！Todo画面へ移動します。");

      navigate("/todos");
    } catch (err: unknown) {
      if (axios.isAxiosError(err)) {
        setError(err.response?.data?.message || "ログインに失敗しました");
      } else {
        setError("予期せぬエラーが発生しました");
      }
    }
  };

  return (
    <div style={{ maxWidth: "400px", margin: "0 auto" }}>
      <h2>🔑 ログイン</h2>

      {error && <p style={{ color: "red", fontWeight: "bold" }}>{error}</p>}

      <form
        onSubmit={handleLogin}
        style={{ display: "flex", flexDirection: "column", gap: "15px" }}
      >
        <div>
          <label>メールアドレス</label>
          <br />
          <input
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
            style={{ width: "100%", padding: "8px" }}
          />
        </div>
        <div>
          <label>パスワード</label>
          <br />
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
            style={{ width: "100%", padding: "8px" }}
          />
        </div>
        <button
          type="submit"
          style={{ padding: "10px", backgroundColor: "blue", color: "white", fontSize: "16px" }}
        >
          ログイン
        </button>
      </form>

      <div style={{ marginTop: "20px", textAlign: "center" }}>
        <Link to="/signup">新規登録はこちら</Link>
      </div>
    </div>
  );
}
