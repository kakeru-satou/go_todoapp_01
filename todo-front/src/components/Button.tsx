// src/components/Button.tsx
import type { ButtonHTMLAttributes, ReactNode } from "react";
import styles from "./Button.module.css";

// 🌟 ボタンが受け取る「指示書（Props）」の型を厳密に定義する
type ButtonProps = ButtonHTMLAttributes<HTMLButtonElement> & {
  variant?: "primary" | "secondary" | "danger" | "light"; // ボタンの種類（デフォルトはprimary）
  fullWidth?: boolean; // 横幅いっぱいにするか
  children?: ReactNode; // ボタンの中身のテキストやアイコン
};

export function Button({
  variant = "primary",
  fullWidth,
  children,
  className,
  type = "button",
  ...props
}: ButtonProps) {
  // 指示書に従って、適用するCSSクラス名を組み立てる
  const btnClass = `
    ${styles.btn} 
    ${styles[variant]} 
    ${fullWidth ? styles.fullWidth : ""} 
    ${className || ""}
  `.trim();

  return (
    // 組み立てたクラス名と、onClickなどの標準機能(...props)をすべてボタンに渡す
    <button className={btnClass} type={type} {...props}>
      {children}
    </button>
  );
}
