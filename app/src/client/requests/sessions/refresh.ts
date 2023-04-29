import { UnknownError } from "../../defaults/unknown_error";
import { AuthenticationResponse } from "../../response_types/authentication_response";
import { ErrorMessage } from "../../response_types/error_message";

export type RefreshTokenResponse =
  | {
    kind: "SUCCESS";
    data: AuthenticationResponse;
  }
  | {
    kind: "FAILURE";
    data: ErrorMessage;
  };

// refreshToken refreshes token and gets new Access Token and Refresh Token
// from API.
export async function refreshToken(): Promise<RefreshTokenResponse> {
  try {
    const resp = await fetch("http://localhost:7755/sessions/refresh", {
      headers: {
        "Content-Type": "application/json",
      },
    });
    const data = await resp.json();

    if (resp.status === 201) {
      return {
        kind: "SUCCESS",
        data: data as AuthenticationResponse,
      };
    } else {
      return {
        kind: "FAILURE",
        data: data as ErrorMessage,
      };
    }
  } catch (error) {
    return UnknownError;
  }
}
