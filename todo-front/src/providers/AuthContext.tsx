import { createContext, useState } from "react";
import type { ReactNode } from "react";
import Cookies from "js-cookie";
import { jwtDecode } from "jwt-decode";

export type AuthContextType = {
  isAuthenticated: boolean;
  login: (token: string) => void;
  logout: () => void;
};

// eslint-disable-next-line react-refresh/only-export-components
export const AuthContext = createContext<AuthContextType | null>(null);

export function AuthProvider({ children }: { children: ReactNode }) {
  const [isAuthenticated, setIsAuthenticated] = useState<boolean>(!!Cookies.get("token"));

  const login = (token: string) => {
    const decoded = jwtDecode(token);
    const expiresAt = decoded.exp == null ? undefined : new Date(decoded.exp * 1000);
    Cookies.set("token", token, { expires: expiresAt });
    setIsAuthenticated(true);
  };

  const logout = () => {
    Cookies.remove("token");
    setIsAuthenticated(false);
  };

  return (
    <AuthContext.Provider value={{ isAuthenticated, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
}
