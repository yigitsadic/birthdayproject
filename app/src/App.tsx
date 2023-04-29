import { useEffect, useState } from "react";
import { LoggedInUser } from "./logged_in_user";
import { refreshToken } from "./client/requests/sessions/refresh";
import { fetchInitialUserData } from "./client/requests/fetch_initial_user_data";

async function fetchUserFromCookie(
  setCurrentUser: React.Dispatch<React.SetStateAction<LoggedInUser | null>>
) {
  const refreshResp = await refreshToken();

  if (refreshResp.kind === "SUCCESS") {
    const userObj = await fetchInitialUserData({
      accessToken: refreshResp.data.access_token,
      company_id: refreshResp.data.company_id,
      user_id: refreshResp.data.user_id,
    });

    setCurrentUser(userObj);
  }
}

function App() {
  const [currentUser, setCurrentUser] = useState<LoggedInUser | null>(null);

  useEffect(() => {
    fetchUserFromCookie(setCurrentUser);
  }, []);

  if (currentUser) {
    return (
      <>
        User is logged in
        <br />
        Welcome, {currentUser.first_name}
      </>
    );
  }

  return <>No user found on session.</>;
}

export default App;
