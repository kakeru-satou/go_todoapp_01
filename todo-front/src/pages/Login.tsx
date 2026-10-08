// src/pages/Login.tsx
import { useState } from "react";
import { useNavigate, Link } from "react-router-dom";
import axios from "axios";

export function Login() {
  // 1. 入力された文字を管理するState
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");

  // 画面遷移を行うためのフック
  const navigate = useNavigate();

  // 2. 「ログイン」ボタンが押された時の処理
  // 🌟 <HTMLFormElement> を指定して、フォーム専用のイベントであることをTypeScriptに教えます
  const handleLogin = async (e: React.SubmitEvent<HTMLFormElement>) => {
    e.preventDefault(); // 🌟 超重要：フォーム送信時の画面リロード（チラつき）を防ぎます
    setError("");

    try {
      // 🌟 axiosを使ってバックエンドにデータを送信！
      // axiosなら JSON.stringify や Headers の設定を自動でやってくれます
      const response = await axios.post("http://localhost:8080/api/auth/signin", {
        email,
        password,
      });

      // 🌟 もらったVIPリストバンド（トークン）をブラウザの金庫に保存！
      localStorage.setItem("token", response.data.token);
      alert("ログイン成功！Todo画面へ移動します。");

      // Todo画面へ自動で移動する
      navigate("/todos");

      // 🌟 any ではなく unknown（中身がわからない安全な箱）としてエラーを受け取ります
    } catch (err: unknown) {
      // 🌟 「これはAxiosが受け取ったエラーだ」と安全確認（型ガード）してからメッセージを取り出します
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

      {/* エラーがある時だけ赤文字で表示 */}
      {error && <p style={{ color: "red", fontWeight: "bold" }}>{error}</p>}

      {/* フォーム全体を <form> で囲み、送信時に handleLogin を動かす */}
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
        {/* Linkコンポーネントを使って画面遷移させる */}
        <Link to="/signup">新規登録はこちら</Link>
      </div>
    </div>
  );
}
