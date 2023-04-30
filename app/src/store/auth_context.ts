import { createContext } from "react";
import { AuthStore } from "./store";

interface IAuthContext {
  authStore: AuthStore | null;
  setAuthStore: React.Dispatch<React.SetStateAction<AuthStore | null>>;
}

export const AuthContext = createContext<IAuthContext>({
  authStore: null,
  setAuthStore: () => { },
});
