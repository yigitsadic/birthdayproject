import { UnknownError } from "../../defaults/unknown_error";
import { ErrorMessage } from "../../response_types/error_message";
import { User } from "../../response_types/user";

export type GetUserDetailResponse =
  | {
    kind: "SUCCESS";
    data: User;
  }
  | {
    kind: "FAILURE" | "UNAUTHENTICATED";
    data: ErrorMessage;
  };

// getUserDetail fetches user detail with given token and user id.
export async function getUserDetail(
  token: string,
  user_id: number
): Promise<GetUserDetailResponse> {
  try {
    const resp = await fetch(`http://localhost:7755/users/${user_id}`, {
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
    });

    const body = await resp.json();

    if (resp.status === 401) {
      return {
        kind: "UNAUTHENTICATED",
        data: body as ErrorMessage,
      };
    }

    if (resp.status === 404) {
      return {
        kind: "FAILURE",
        data: body as ErrorMessage,
      };
    }

    return {
      kind: "SUCCESS",
      data: body as User,
    };
  } catch (error) {
    return UnknownError;
  }
}
