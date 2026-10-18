// src/pages/Todo.tsx
import { useState } from "react";
import { useNavigate, Navigate } from "react-router-dom"; // 🌟 ログアウト後の画面移動用
import { useAuth } from "../hooks/useAuth";
import { useTodos, type TodoType } from "../hooks/useTodos";

import {
  FaPenToSquare,
  FaListUl,
  FaPlus,
  FaArrowRightFromBracket,
  FaTrash,
  FaArrowLeft,
  FaSpinner,
} from "react-icons/fa6";
import { Button } from "../components/Button";

import styles from "./Todo.module.css";

export function TodoModal() {
  const { userName, isAuthenticated, logout } = useAuth();
  const navigate = useNavigate();
  const { todos, isLoading, isError, addTodo, deleteTodo, updateTodo } = useTodos();

  // 🌟 モーダルの開閉状態を切り替えるためのState（最初はどちらも非表示の false）
  const [isAddModalOpen, setIsAddModalOpen] = useState(false);
  const [editingTodo, setEditingTodo] = useState<TodoType | null>(null); // nullなら編集モーダルは閉じている状態

  // 🌟 フォームへの入力内容を一時保存するState
  const [formTask, setFormTask] = useState("");

  // 🌟 4-5と同じ：トークンの破棄は Context、画面移動はこの画面の担当
  const handleLogout = () => {
    logout();
    navigate("/login");
  };

  // 新規追加モーダルの「追加する」ボタンが押された時の処理
  const handleAddSubmit = async (e: React.SubmitEvent<HTMLFormElement>) => {
    e.preventDefault();
    await addTodo({ task: formTask });
    setIsAddModalOpen(false); // 登録が終わったらモーダルをスッと閉じる
    setFormTask("");
  };

  const handleDelete = async (id: number) => {
    if (window.confirm("本当に削除しますか？")) {
      try {
        await deleteTodo(id);
      } catch (error) {
        console.error(error);
      }
    }
  };

  // 編集モーダルの「更新する」ボタンが押された時の処理
  const handleEditSubmit = async (e: React.SubmitEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (!editingTodo) return;
    const task = formTask;
    await updateTodo({ id: editingTodo.id, task });
    setEditingTodo(null); // 更新が終わったらモーダルを閉じる
  };

  // 既存のTodoの「編集」ボタンを押した時の処理
  const openEditModal = (todo: TodoType) => {
    // 💡 新たにAPI通信をせず、手元にあるデータをそのままフォームにセットしてモーダルを開く！
    setFormTask(todo.task);
    setEditingTodo(todo);
  };

  if (!isAuthenticated) {
    return <Navigate to="/login" replace />;
  }

  if (isLoading) {
    return (
      <div
        style={{
          display: "flex",
          justifyContent: "center",
          alignItems: "center",
          height: "100vh",
          fontSize: "2rem",
          color: "var(--primary-color)",
        }}
      >
        <FaSpinner style={{ animation: "spin 1s linear infinite" }} />
      </div>
    );
  }

  if (isError) {
    return (
      <div style={{ color: "red", textAlign: "center", marginTop: "50px" }}>
        データの取得に失敗しました。
      </div>
    );
  }

  return (
    <div className={styles.appContainer}>
      {/* 🌟 サイドバー */}
      <aside className={styles.sidebar}>
        <div className={styles.sidebarHeader}>
          <h2>
            <FaPenToSquare /> MY TODOS
          </h2>
        </div>
        <nav className={styles.sidebarNav}>
          <ul>
            <li className={styles.active}>
              <FaListUl /> Todo一覧
            </li>
            <li
              onClick={() => {
                setIsAddModalOpen(true);
                setFormTask("");
              }}
            >
              <FaPlus /> 新規Todo追加
            </li>
          </ul>
        </nav>
        <div className={styles.sidebarFooter}>
          {/* 🌟 自作したButtonコンポーネントを使用！ */}
          <Button onClick={handleLogout} fullWidth>
            ログアウト
          </Button>
        </div>
      </aside>

      {/* 🌟 メインコンテンツ */}
      <main className={styles.mainContent}>
        <header className={styles.topHeader}>
          {/* 🌟 APIから取得した最新のユーザー名を表示！ */}
          <h1>Hello, {userName}!</h1>

          <div className={styles.userProfile}>
            {/* variant="danger" を指定するだけでスタイルが適用される！ */}
            <Button
              variant="danger"
              onClick={handleLogout}
              className={styles.mobileLogoutBtn}
              title="ログアウト"
            >
              <FaArrowRightFromBracket />
            </Button>
          </div>
        </header>

        <section className={styles.todoSection}>
          <div className={styles.sectionHeader}>
            <h2>TODO LIST</h2>
            <Button
              onClick={() => {
                setIsAddModalOpen(true);
                setFormTask("");
              }}
            >
              <FaPlus /> 新規Task
            </Button>
          </div>

          {/* {error && <p style={{ color: "red" }}>{error}</p>} */}

          <div className={styles.todoList}>
            {todos.length > 0 && (
              <div className={styles.todoHeader}>
                <span>タイトル</span>
                <span style={{ textAlign: "right" }}>アクション</span>
              </div>
            )}

            {todos.map((todo) => (
              <div key={todo.id} className={styles.todoItem}>
                <div className={styles.todoTitle}>{todo.task}</div>
                <div className={styles.todoActions}>
                  <input
                    type="checkbox"
                    checked={todo.isDone}
                    onChange={() => updateTodo({ id: todo.id, isDone: !todo.isDone })}
                  ></input>
                  <Button variant="secondary" onClick={() => openEditModal(todo)}>
                    編集
                  </Button>
                  <Button variant="danger" onClick={() => handleDelete(todo.id)}>
                    <FaTrash />
                  </Button>
                </div>
              </div>
            ))}
          </div>
        </section>
      </main>

      {/* ここから下に、モーダルのコード（演習課題）が入ります！ */}
      {isAddModalOpen && (
        <div className={styles.modalOverlay}>
          <div className={styles.modalContent}>
            <button className={styles.btnClose} onClick={() => setIsAddModalOpen(false)}>
              <FaArrowLeft />
            </button>

            <h3 className={styles.modalTitle}>✨ 新規Todo追加</h3>
            <form onSubmit={handleAddSubmit}>
              <div className={styles.formGroup}>
                <label>タイトル</label>
                <input
                  value={formTask}
                  onChange={(e) => setFormTask(e.target.value)}
                  placeholder="タスクのタイトル"
                  required
                />
              </div>
              <div className={styles.modalActions}>
                <Button type="submit">追加する</Button>
                <Button variant="light" onClick={() => setIsAddModalOpen(false)}>
                  キャンセル
                </Button>
              </div>
            </form>
          </div>
        </div>
      )}
      {editingTodo && (
        <div className={styles.modalOverlay}>
          <div className={styles.modalContent}>
            <button className={styles.btnClose} onClick={() => setEditingTodo(null)}>
              <FaArrowLeft />
            </button>

            <h3 className={styles.modalTitle}>✏️ Todo編集</h3>
            <form onSubmit={handleEditSubmit}>
              <div className={styles.formGroup}>
                <label>タイトル</label>
                <input
                  value={formTask}
                  onChange={(e) => setFormTask(e.target.value)}
                  placeholder="タスクのタイトル"
                  required
                />
              </div>
              <div className={styles.modalActions}>
                <Button type="submit">更新する</Button>
                <Button variant="light" onClick={() => setEditingTodo(null)}>
                  キャンセル
                </Button>
              </div>
            </form>
          </div>
        </div>
      )}
    </div>
    // <div style={{ padding: "20px", maxWidth: "800px", margin: "0 auto" }}>
    //   <div style={{ display: "flex", justifyContent: "space-between", alignItems: "center" }}>
    //     <h2>📋 Todo一覧</h2>
    //     <button
    //       onClick={handleLogout}
    //       style={{ padding: "5px 10px", backgroundColor: "gray", color: "white" }}
    //     >
    //       ログアウト
    //     </button>
    //   </div>

    //   {/* クリックされたら、追加用モーダルのフラグを true にしてフォームを空にする */}
    //   <button
    //     onClick={() => {
    //       setIsAddModalOpen(true);
    //       setFormTask("");
    //     }}
    //     style={{ margin: "20px 0", padding: "10px", backgroundColor: "blue", color: "white" }}
    //   >
    //     + 新規Todo追加
    //   </button>

    //   {error && <p style={{ color: "red" }}>{error}</p>}

    //   {/* Todoリストの表示 */}
    //   <ul style={{ listStyle: "none", padding: 0 }}>
    //     {todos.map((todo) => (
    //       <li
    //         key={todo.id}
    //         style={{
    //           border: "1px solid #ccc",
    //           margin: "10px 0",
    //           padding: "15px",
    //           backgroundColor: "white",
    //         }}
    //       >
    //         <h3>{todo.task}</h3>
    //         <div style={{ display: "flex", gap: "10px", marginTop: "10px" }}>
    //           <input
    //             type="checkbox"
    //             checked={todo.isDone}
    //             onChange={() => updateTodo(todo.id, { isDone: !todo.isDone })}
    //           />
    //           <button onClick={() => openEditModal(todo)}>編集</button>
    //           <button onClick={() => deleteTodo(todo.id)} style={{ color: "red" }}>
    //             削除
    //           </button>
    //         </div>
    //       </li>
    //     ))}
    //   </ul>

    //   {/* ==================================================
    //       🚀 新規追加モーダル（isAddModalOpen が true の時だけ画面に現れる）
    //       ================================================== */}
    //   {isAddModalOpen && (
    //     <div
    //       style={{
    //         position: "fixed",
    //         top: 0,
    //         left: 0,
    //         width: "100vw",
    //         height: "100vh",
    //         backgroundColor: "rgba(0,0,0,0.5)",
    //         display: "flex",
    //         justifyContent: "center",
    //         alignItems: "center",
    //       }}
    //     >
    //       <div
    //         style={{
    //           backgroundColor: "white",
    //           padding: "30px",
    //           borderRadius: "10px",
    //           width: "400px",
    //         }}
    //       >
    //         <h3 style={{ textAlign: "center" }}>✨ 新規Todo追加</h3>
    //         <form
    //           onSubmit={handleAddSubmit}
    //           style={{ display: "flex", flexDirection: "column", gap: "15px", marginTop: "20px" }}
    //         >
    //           <div>
    //             <label style={{ fontWeight: "bold" }}>タイトル</label>
    //             <input
    //               value={formTask}
    //               onChange={(e) => setFormTask(e.target.value)}
    //               placeholder="タスクのタイトル"
    //               required
    //               style={{ width: "100%", padding: "10px", marginTop: "5px" }}
    //             />
    //           </div>
    //           <div style={{ display: "flex", gap: "10px" }}>
    //             <button
    //               type="submit"
    //               style={{
    //                 flex: 1,
    //                 backgroundColor: "blue",
    //                 color: "white",
    //                 padding: "10px",
    //                 border: "none",
    //                 borderRadius: "5px",
    //               }}
    //             >
    //               追加する
    //             </button>
    //             <button
    //               type="button"
    //               onClick={() => setIsAddModalOpen(false)}
    //               style={{
    //                 flex: 1,
    //                 padding: "10px",
    //                 border: "1px solid #ccc",
    //                 borderRadius: "5px",
    //               }}
    //             >
    //               キャンセル
    //             </button>
    //           </div>
    //         </form>
    //       </div>
    //     </div>
    //   )}

    //   {/* ==================================================
    //       🚀 編集モーダル（editingTodo にデータが入っている時だけ画面に現れる）
    //       ================================================== */}
    //   {editingTodo && (
    //     <div
    //       style={{
    //         position: "fixed",
    //         top: 0,
    //         left: 0,
    //         width: "100vw",
    //         height: "100vh",
    //         backgroundColor: "rgba(0,0,0,0.5)",
    //         display: "flex",
    //         justifyContent: "center",
    //         alignItems: "center",
    //       }}
    //     >
    //       <div
    //         style={{
    //           backgroundColor: "white",
    //           padding: "30px",
    //           borderRadius: "10px",
    //           width: "400px",
    //         }}
    //       >
    //         <h3 style={{ textAlign: "center" }}>✏️ Todo編集</h3>
    //         <form
    //           onSubmit={handleEditSubmit}
    //           style={{ display: "flex", flexDirection: "column", gap: "15px", marginTop: "20px" }}
    //         >
    //           <div>
    //             <label style={{ fontWeight: "bold" }}>タイトル</label>
    //             <input
    //               value={formTask}
    //               onChange={(e) => setFormTask(e.target.value)}
    //               placeholder="タスクのタイトル"
    //               required
    //               style={{ width: "100%", padding: "10px", marginTop: "5px" }}
    //             />
    //           </div>
    //           <div style={{ display: "flex", gap: "10px" }}>
    //             <button
    //               type="submit"
    //               style={{
    //                 flex: 1,
    //                 backgroundColor: "green",
    //                 color: "white",
    //                 padding: "10px",
    //                 border: "none",
    //                 borderRadius: "5px",
    //               }}
    //             >
    //               更新する
    //             </button>
    //             <button
    //               type="button"
    //               onClick={() => setEditingTodo(null)}
    //               style={{
    //                 flex: 1,
    //                 padding: "10px",
    //                 border: "1px solid #ccc",
    //                 borderRadius: "5px",
    //               }}
    //             >
    //               キャンセル
    //             </button>
    //           </div>
    //         </form>
    //       </div>
    //     </div>
    //   )}
    // </div>
  );
}
