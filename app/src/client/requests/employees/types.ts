import { Employee } from "../../response_types/employee";
import { ErrorMessage } from "../../response_types/error_message";
import { AuthenticationParams } from "../authentication_params";

export type SingularEmployeeResponse =
  | {
    kind: "SUCCESS";
    data: Employee;
  }
  | {
    kind: "FAILURE" | "UNAUTHENTICATED";
    data: ErrorMessage;
  };

export interface SharedEmployeeListParams extends AuthenticationParams {
  company_id: number;
}

export interface SharedEmployeeDetailParams extends SharedEmployeeListParams {
  employee_id: number;
}
