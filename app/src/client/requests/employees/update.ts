import { UnknownError } from "../../defaults/unknown_error";
import { EmployeeDto } from "../../dtos/employee_dto";
import { Employee } from "../../response_types/employee";
import { ErrorMessage } from "../../response_types/error_message";
import { SharedEmployeeDetailParams, SingularEmployeeResponse } from "./types";

export interface EmployeeUpdateParams extends SharedEmployeeDetailParams {
  dto: EmployeeDto;
}

// employeeUpdate updates given employee with params.
export async function employeeUpdate(
  params: EmployeeUpdateParams
): Promise<SingularEmployeeResponse> {
  try {
    const resp = await fetch(
      `${import.meta.env.VITE_API_URL}/companies/${params.company_id
      }/employees/${params.employee_id}`,
      {
        method: "PUT",
        body: JSON.stringify(params.dto),
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

    if (resp.status === 422) {
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
