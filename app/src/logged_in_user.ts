export type LoggedInUser =
  | {
    logged_in: true;
    user_id: number;
    first_name: string;
    last_name: string;
    email: string;
    company_name: string;
    company_id: number;
    access_token: string;
  }
  | {
    logged_in: false;
  };
