import { useEffect, useState } from "react";
import { Link, Outlet } from "react-router-dom";
import { LoggedInUser } from "../logged_in_user";
import { refreshToken } from "../client/requests/sessions/refresh";
import { fetchInitialUserData } from "../client/requests/fetch_initial_user_data";
import { AuthStore } from "../store/store";
import { AuthContext } from "../store/auth_context";

async function fetchUserFromCookie(
  setCurrentUser: React.Dispatch<React.SetStateAction<LoggedInUser>>,
  setAuthStore: React.Dispatch<React.SetStateAction<AuthStore | null>>
) {
  const refreshResp = await refreshToken();

  if (refreshResp.kind === "SUCCESS") {
    const userObj = await fetchInitialUserData({
      accessToken: refreshResp.data.access_token,
      company_id: refreshResp.data.company_id,
      user_id: refreshResp.data.user_id,
    });

    setAuthStore({
      user_id: refreshResp.data.user_id,
      company_id: refreshResp.data.company_id,
      access_token: refreshResp.data.access_token,
    });

    setCurrentUser(userObj);
  }
}
export const DashboardPage = () => {
  const [currentUser, setCurrentUser] = useState<LoggedInUser>({
    logged_in: false,
  });

  const [authStore, setAuthStore] = useState<AuthStore | null>(null);

  useEffect(() => {
    fetchUserFromCookie(setCurrentUser, setAuthStore);
  }, []);

  if (currentUser?.logged_in) {
    return (
      <div>
        Hello {currentUser.first_name} {currentUser.last_name}, <br />
        <table className="table">
          <tbody>
            <tr>
              <td>
                <Link to={"/"}>Dashboard</Link>
              </td>
              <td>
                <Link to={"/me"}>Me</Link>
              </td>
              <td>
                <Link to={"/company"}>Company</Link>
              </td>
            </tr>
          </tbody>
        </table>
        <hr />
        <AuthContext.Provider value={{ authStore, setAuthStore }}>
          <Outlet />
        </AuthContext.Provider>
      </div>
    );
  } else {
    return (
      <div>
        Hello unknown user!, <br />
        <Outlet />
      </div>
    );
  }
};
