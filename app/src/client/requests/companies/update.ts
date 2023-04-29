import { UnknownError } from "../../defaults/unknown_error";
import { CompanyDto } from "../../dtos/company_dto";
import { Company } from "../../response_types/company";
import { ErrorMessage } from "../../response_types/error_message";
import { CompanyParams } from "./company_params";

export interface UpdateCompanyParams extends CompanyParams {
  dto: CompanyDto;
}

export type UpdateCompanyResponse =
  | {
    kind: "SUCCESS";
    data: Company;
  }
  | {
    kind: "FAILURE" | "UNAUTHENTICATED";
    data: ErrorMessage;
  };

export async function updateCompany(
  params: UpdateCompanyParams
): Promise<UpdateCompanyResponse> {
  try {
    const resp = await fetch(
      `http://localhost:7755/companies/${params.company_id}`,
      {
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${params.accessToken}`,
        },
        body: JSON.stringify(params.dto),
        method: "PUT",
      }
    );

    const data = await resp.json();

    if (resp.status === 200) {
      return {
        kind: "SUCCESS",
        data: data as Company,
      };
    }

    if (resp.status === 422) {
      return {
        kind: "FAILURE",
        data: data as ErrorMessage,
      };
    }

    if (resp.status === 401) {
      return {
        kind: "UNAUTHENTICATED",
        data: data as ErrorMessage,
      };
    }

    return UnknownError;
  } catch (error) {
    return UnknownError;
  }
}
