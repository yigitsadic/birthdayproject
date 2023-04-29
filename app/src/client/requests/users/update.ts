import { UnknownError } from "../../defaults/unknown_error";
import { UserDto } from "../../dtos/user_dto";
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

    const data = await resp.json();

    if (resp.status === 200) {
      return {
        kind: "SUCCESS",
        data: data as User,
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
