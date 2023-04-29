import { UnknownError } from "../../defaults/unknown_error";
import { ErrorMessage } from "../../response_types/error_message";
import { User } from "../../response_types/user";
import { UserParams } from "./user_params";

export type GetUserDetailResponse =
  | {
    kind: "SUCCESS";
    data: User;
    type: "UserResponse";
  }
  | {
    kind: "FAILURE" | "UNAUTHENTICATED";
    data: ErrorMessage;
  };

// getUserDetail fetches user detail with given token and user id.
export async function getUserDetail(
  params: UserParams
): Promise<GetUserDetailResponse> {
  try {
    const resp = await fetch(
      `${import.meta.env.VITE_API_URL}/users/${params.user_id}`,
      {
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${params.accessToken}`,
        },
      }
    );

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

    if (resp.status === 200) {
      return {
        type: "UserResponse",
        kind: "SUCCESS",
        data: body as User,
      };
    }

    return UnknownError;
  } catch (error) {
    return UnknownError;
  }
}
