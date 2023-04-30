// import { useContext, useEffect } from "react";
import { AuthContext } from "../../store/auth_context";
import { getUserDetail } from "../../client/requests/users/detail";
import { AuthenticationRequired } from "../general/authentication_required";
import { useContext, useEffect, useState } from "react";
import { User } from "../../client/response_types/user";
import { UserDisplay } from "./display";
import { UserForm } from "./form";

export const UserDetailPage = () => {
  const { authStore } = useContext(AuthContext);

  if (!authStore) {
    return <AuthenticationRequired />;
  }

  const [user, setUser] = useState<User | null>(null);
  const [mode, setMode] = useState<"display" | "edit">("display");

  const fetchFromAPI = async () => {
    const result = await getUserDetail({
      accessToken: authStore.access_token,
      user_id: authStore.user_id,
    });

    if (result.kind === "SUCCESS") {
      setUser(result.data);
    }
  };

  useEffect(() => {
    fetchFromAPI();
  }, [authStore.access_token]);

  if (user) {
    return (
      <div>
        <h3>User Information</h3>

        {mode === "display" ? (
          <UserDisplay
            setMode={setMode}
            user={user}
            accessToken={authStore.access_token}
          />
        ) : (
          <UserForm
            setMode={setMode}
            setUser={setUser}
            accessToken={authStore.access_token}
            user={user}
          />
        )}
      </div>
    );
  }

  return <>No result found.</>;
};
