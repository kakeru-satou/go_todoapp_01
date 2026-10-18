import { createContext, useState, useEffect, useCallback } from "react";
import type { ReactNode } from "react";
import Cookies from "js-cookie";
import { jwtDecode } from "jwt-decode";
import axios from "axios";

export type AuthContextType = {
  isAuthenticated: boolean;
  userName: string;
  login: (token: string) => Promise<void>;
  logout: () => void;
};

// eslint-disable-next-line react-refresh/only-export-components
export const AuthContext = createContext<AuthContextType | null>(null);

export function AuthProvider({ children }: { children: ReactNode }) {
  const [isAuthenticated, setIsAuthenticated] = useState<boolean>(!!Cookies.get("token"));
  const [userName, setUserName] = useState<string>("ゲスト");

  const fetchUserProfile = useCallback(async (token: string) => {
    try {
      const res = await axios.get("http://localhost:8080/api/users/me", {
        headers: { Authorization: `Bearer ${token}` },
      });
      setUserName(res.data.name);
    } catch (error) {
      console.error("プロフィール取得失敗", error);
    }
  }, []);

  useEffect(() => {
    const token = Cookies.get("token");
    if (token) {
      // eslint-disable-next-line react-hooks/set-state-in-effect
      fetchUserProfile(token);
    }
  }, [fetchUserProfile]);

  const login = async (token: string) => {
    const decoded = jwtDecode(token);
    const expiresAt = decoded.exp == null ? undefined : new Date(decoded.exp * 1000);
    Cookies.set("token", token, { expires: expiresAt });
    setIsAuthenticated(true);
    await fetchUserProfile(token);
  };

  const logout = () => {
    Cookies.remove("token");
    setIsAuthenticated(false);
    setUserName("ゲスト");
  };

  return (
    <AuthContext.Provider value={{ isAuthenticated, userName, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
}
