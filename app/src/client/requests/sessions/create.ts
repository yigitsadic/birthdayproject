import { UnknownError } from "../../defaults/unknown_error";
import { SessionDto } from "../../input_types/session_dto";
import { AuthenticationResponse } from "../../response_types/authentication_response";
import { ErrorMessage } from "../../response_types/error_message";

export type SessionCreateResponse =
  | {
    kind: "SUCCESS";
    data: AuthenticationResponse;
  }
  | {
    kind: "FAILURE";
    data: ErrorMessage;
  };

// sessionCreate sends request to API for creating AccessToken and RefreshToken.
export async function sessionCreate(
  dto: SessionDto
): Promise<SessionCreateResponse> {
  try {
    const resp = await fetch("http://localhost:7755/sessions/create", {
      method: "POST",
      body: JSON.stringify(dto),
      headers: {
        "Content-Type": "application/json",
      },
    });
    const body = await resp.json();

    if (resp.status === 201) {
      return {
        kind: "SUCCESS",
        data: body as AuthenticationResponse,
      };
    } else {
      return {
        kind: "FAILURE",
        data: body as ErrorMessage,
      };
    }
  } catch (error) {
    return UnknownError;
  }
}
