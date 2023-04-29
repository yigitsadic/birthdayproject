import { AuthenticationParams } from "../authentication_params";

export interface UserParams extends AuthenticationParams {
  user_id: number;
}
