import { useState } from "react";
import { useNavigate, Link } from "react-router-dom";
import axios from "axios";

export function Signup() {
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");

  const navigate = useNavigate();

  const handleSignup = async (e: React.SubmitEvent<HTMLFormElement>) => {
    e.preventDefault();
    setError("");

    try {
      const response = await axios.post("http://localhost:8080/api/auth/signup", {
        name,
        email,
        password,
      });

      localStorage.setItem("token", response.data.token);
      alert("登録成功！Todo画面へ移動します。");

      navigate("/todos");
    } catch (err: unknown) {
      if (axios.isAxiosError(err)) {
        setError(err.response?.data?.message || "登録に失敗しました");
      } else {
        setError("予期せぬエラーが発生しました");
      }
    }
  };

  return (
    <div style={{ maxWidth: "400px", margin: "0 auto" }}>
      <h2>🔑 新規登録</h2>

      {error && <p style={{ color: "red", fontWeight: "bold" }}>{error}</p>}

      <form
        onSubmit={handleSignup}
        style={{ display: "flex", flexDirection: "column", gap: "15px" }}
      >
        <div>
          <label>名前</label>
          <br />
          <input
            type="text"
            value={name}
            onChange={(e) => setName(e.target.value)}
            required
            style={{ width: "100%", padding: "8px" }}
          />
        </div>
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
          新規登録
        </button>
      </form>

      <div style={{ marginTop: "20px", textAlign: "center" }}>
        <Link to="/login">ログインはこちら</Link>
      </div>
    </div>
  );
}
