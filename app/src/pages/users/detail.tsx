// import { useContext, useEffect } from "react";
import { AuthContext } from "../../store/auth_context";
import { getUserDetail } from "../../client/requests/users/detail";
import { AuthenticationRequired } from "../general/authentication_required";
import { useContext, useEffect, useState } from "react";
import { User } from "../../client/response_types/user";

export const UserDetailPage = () => {
  const { authStore } = useContext(AuthContext);

  if (!authStore) {
    return <AuthenticationRequired />;
  }

  const [user, setUser] = useState<User | null>(null);

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
        Authenticated user information
        <table>
          <tbody>
            <tr>
              <th>User ID</th>
              <td>{user.id}</td>
            </tr>
            <tr>
              <th>First Name</th>
              <td>{user.first_name}</td>
            </tr>
            <tr>
              <th>Last Name</th>
              <td>{user.last_name}</td>
            </tr>
            <tr>
              <th>Email</th>
              <td>{user.email}</td>
            </tr>
          </tbody>
        </table>
      </div>
    );
  }

  return <>No result found.</>;
};
