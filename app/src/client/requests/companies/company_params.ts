import { AuthenticationParams } from "../authentication_params";

export interface CompanyParams extends AuthenticationParams {
  company_id: number;
}
