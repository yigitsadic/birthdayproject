import { UnknownError } from "../../defaults/unknown_error";
import { Employee } from "../../response_types/employee";
import { ErrorMessage } from "../../response_types/error_message";
import { SharedEmployeeListParams } from "./types";

export type EmployeeListResponse =
  | {
    kind: "SUCCESS";
    data: Employee[];
  }
  | {
    kind: "FAILURE" | "UNAUTHENTICATED";
    data: ErrorMessage;
  };

// employeesList fetches list of employees with given company id and access token.
export async function employeesList(
  params: SharedEmployeeListParams
): Promise<EmployeeListResponse> {
  try {
    const resp = await fetch(
      `${import.meta.env.VITE_API_URL}/companies/${params.company_id
      }/employees`,
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
        data: data as Employee[],
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
