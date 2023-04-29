import { UnknownError } from "../../defaults/unknown_error";
import { Employee } from "../../response_types/Employee";
import { ErrorMessage } from "../../response_types/error_message";
import { SharedEmployeeDetailParams, SingularEmployeeResponse } from "./types";

// employeeDetail fetches employee detail with given params.
export async function employeeDetail(
  params: SharedEmployeeDetailParams
): Promise<SingularEmployeeResponse> {
  try {
    const resp = await fetch(
      `http://localhost:7755/companies/${params.company_id}/employees/${params.employee_id}`,
      {
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${params.accessToken}`,
        },
      }
    );

    const data = await resp.json();

    if (resp.status === 200) {
      return {
        kind: "SUCCESS",
        data: data as Employee,
      };
    }

    if (resp.status === 401) {
      return {
        kind: "UNAUTHENTICATED",
        data: data as ErrorMessage,
      };
    }

    if (resp.status === 404) {
      return {
        kind: "FAILURE",
        data: data as ErrorMessage,
      };
    }

    return UnknownError;
  } catch (error) {
    return UnknownError;
  }
}
