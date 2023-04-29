import { UnknownError } from "../../defaults/unknown_error";
import { ErrorMessage } from "../../response_types/error_message";
import { SharedEmployeeDetailParams } from "./types";

export type EmployeeDestroyResponse =
  | {
    kind: "SUCCESS";
  }
  | {
    kind: "FAILURE" | "UNAUTHENTICATED";
    data: ErrorMessage;
  };

// destroyEmployee deletes given employee.
export async function employeeDestroy(
  params: SharedEmployeeDetailParams
): Promise<EmployeeDestroyResponse> {
  try {
    const resp = await fetch(
      `${import.meta.env.VITE_API_URL}/companies/${params.company_id
      }/employees/${params.employee_id}`,
      {
        method: "DELETE",
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
