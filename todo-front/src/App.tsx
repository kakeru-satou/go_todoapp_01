// src/App.tsx
import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom";
import { Login } from "./pages/Login";
import { Signup } from "./pages/Signup";
import { Todo } from "./pages/Todo";

export function App() {
  return (
    // 🌟 BrowserRouter: この中でルーティングを使うよ！という宣言
    <BrowserRouter>
      <div style={{ padding: "20px", fontFamily: "sans-serif" }}>
        {/* 🌟 Routes: URLとコンポーネントの組み合わせリスト */}
        <Routes>
          <Route path="/login" element={<Login />} />
          <Route path="/signup" element={<Signup />} />
          <Route path="/todos" element={<Todo />} />

          {/* 存在しないURLにアクセスされたら、自動で /login に飛ばす設定 */}
          <Route path="*" element={<Navigate to="/login" replace />} />
        </Routes>
      </div>
    </BrowserRouter>
  );
}

export default App;
