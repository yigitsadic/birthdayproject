import { LoggedInUser } from "../../logged_in_user";
import { companyDetail } from "./companies/detail";
import { getUserDetail } from "./users/detail";

export interface fetchInitialUserDataParams {
  user_id: number;
  company_id: number;
  accessToken: string;
}

// fetchInitialUserData fetches user information with company information after login.
export async function fetchInitialUserData(
  params: fetchInitialUserDataParams
): Promise<LoggedInUser> {
  try {
    const { user_id, company_id, accessToken } = params;

    const userDataPromise = getUserDetail({
      user_id,
      accessToken,
    });

    const companyDataPromise = companyDetail({
      accessToken,
      company_id,
    });

    const userObj: Partial<LoggedInUser> = { access_token: accessToken };
    await Promise.all([userDataPromise, companyDataPromise])
      .then((results) => {
        results.forEach((result) => {
          if (result.kind === "SUCCESS") {
            if (result.type === "CompanyDetailResponse") {
              userObj.logged_in = true;
              userObj.company_id = result.data.id;
              userObj.company_name = result.data.name;
            }

            if (result.type === "UserResponse") {
              userObj.logged_in = true;
              userObj.user_id = result.data.id;
              userObj.first_name = result.data.first_name;
              userObj.last_name = result.data.last_name;
              userObj.email = result.data.email;
            }
          }
        });
      })
      .catch(() => {
        return { logged_in: false };
      });

    return userObj as LoggedInUser;
  } catch (error) {
    return { logged_in: false };
  }
}
