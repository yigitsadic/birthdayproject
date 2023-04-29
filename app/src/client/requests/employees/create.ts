import { UnknownError } from "../../defaults/unknown_error";
import { EmployeeDto } from "../../dtos/employee_dto";
import { Employee } from "../../response_types/employee";
import { ErrorMessage } from "../../response_types/error_message";
import { SharedEmployeeListParams, SingularEmployeeResponse } from "./types";

export interface CreateEmployeeParams extends SharedEmployeeListParams {
  dto: EmployeeDto;
}

// createEmployee creates a new employee with given params.
export async function createEmployee(
  params: CreateEmployeeParams
): Promise<SingularEmployeeResponse> {
  try {
    const resp = await fetch(
      `${import.meta.env.VITE_API_URL}/companies/${params.company_id
      }/employees`,
      {
        method: "POST",
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
