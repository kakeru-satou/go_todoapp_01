import { createContext, useState } from "react";
import { useQuery, useQueryClient } from "@tanstack/react-query";
import type { ReactNode } from "react";
import Cookies from "js-cookie";
import { jwtDecode } from "jwt-decode";
import axios from "axios";

export type AuthContextType = {
  isAuthenticated: boolean;
  userName: string;
  login: (token: string) => void;
  logout: () => void;
};

// eslint-disable-next-line react-refresh/only-export-components
export const AuthContext = createContext<AuthContextType | null>(null);

const fetchUserProfile = async (token: string): Promise<string> => {
  const res = await axios.get("http://localhost:8080/api/users/me", {
    headers: { Authorization: `Bearer ${token}` },
  });
  return res.data.name;
};

export function AuthProvider({ children }: { children: ReactNode }) {
  const [token, setToken] = useState<string | undefined>(Cookies.get("token"));
  const queryClient = useQueryClient();
  // const [isAuthenticated, setIsAuthenticated] = useState<boolean>(!!Cookies.get("token"));
  // const [userName, setUserName] = useState<string>("ゲスト");

  // useEffect(() => {
  //   const token = Cookies.get("token");
  //   if (token) {
  //     fetchUserProfile(token);
  //   }
  // }, [fetchUserProfile]);

  const { data: userName = "ゲスト" } = useQuery({
    queryKey: ["userProfile", token],
    queryFn: () => fetchUserProfile(token!),
    enabled: !!token,
  });

  const login = (newToken: string) => {
    const decoded = jwtDecode(newToken);
    const expiresAt = decoded.exp == null ? undefined : new Date(decoded.exp * 1000);
    Cookies.set("token", newToken, { expires: expiresAt });
    setToken(newToken);
  };

  const logout = () => {
    Cookies.remove("token");
    setToken(undefined);
    queryClient.clear();
  };

  return (
    <AuthContext.Provider value={{ isAuthenticated: !!token, userName, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
}
