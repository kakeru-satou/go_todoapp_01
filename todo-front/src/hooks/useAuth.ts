import { useContext } from "react";
import { AuthContext } from "../providers/AuthContext";

export function useAuth() {
  const result = useContext(AuthContext);

  if (result == null) {
    throw new Error("useAuth は AuthProvider の内側で使ってください");
  }

  return result;
}
