import { UnknownError } from "../../defaults/unknown_error";
import { Company } from "../../response_types/company";
import { ErrorMessage } from "../../response_types/error_message";
import { CompanyParams } from "./company_params";

export type CompanyDetailResponse =
  | {
    kind: "SUCCESS";
    data: Company;
    type: "CompanyDetailResponse";
  }
  | {
    kind: "FAILURE" | "UNAUTHENTICATED";
    data: ErrorMessage;
  };

// companyDetail fetches company details with given access token and company id.
export async function companyDetail(
  params: CompanyParams
): Promise<CompanyDetailResponse> {
  try {
    const resp = await fetch(
      `${import.meta.env.VITE_API_URL}/companies/${params.company_id}`,
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
        data: data as Company,
        type: "CompanyDetailResponse",
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
