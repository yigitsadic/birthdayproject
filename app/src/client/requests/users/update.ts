import { UnknownError } from "../../defaults/unknown_error";
import { UserDto } from "../../input_types/user_dto";
import { ErrorMessage } from "../../response_types/error_message";
import { User } from "../../response_types/user";
import { UserParams } from "./user_params";

export type UpdateUserResponse =
  | {
    kind: "SUCCESS";
    data: User;
  }
  | {
    kind: "FAILURE" | "UNAUTHENTICATED";
    data: ErrorMessage;
  };

export interface UpdateUserParams extends UserParams {
  dto: UserDto;
}

export async function updateUser(
  params: UpdateUserParams
): Promise<UpdateUserResponse> {
  try {
    const resp = await fetch(`http://localhost:7755/users/${params.user_id}`, {
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${params.accessToken}`,
      },
      body: JSON.stringify(params.dto),
      method: "PUT",
    });

    const body = await resp.json();

    if (resp.status === 200) {
      return {
        kind: "SUCCESS",
        data: body as User,
      };
    }

    if (resp.status === 401) {
      return {
        kind: "UNAUTHENTICATED",
        data: body as ErrorMessage,
      };
    }

    return UnknownError;
  } catch (error) {
    return UnknownError;
  }
}