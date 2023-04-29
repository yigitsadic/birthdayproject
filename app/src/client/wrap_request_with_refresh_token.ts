import { AuthenticationParams } from "./requests/authentication_params";
import { refreshToken } from "./requests/sessions/refresh";

// WrapRequestWithRefreshToken makes given request. When it gets unauthenticated response, retries with
// refresh token.
export async function WrapRequestWithRefreshToken<
  RReturnType extends { kind: string },
  RParams extends AuthenticationParams
>(func: Function, params: RParams): Promise<RReturnType> {
  const result = (await func(params)) as RReturnType;
  if (result.kind === "SUCCESS" || result.kind === "FAILURE") {
    // If request resulted with success or failure (404 or 500) return with result.

    return result;
  } else {
    // If request resulted with 401, that means it is unauthenticated.
    // Try to re-authenticate via refresh token.
    const refreshTokenResp = await refreshToken();

    if (refreshTokenResp.kind === "SUCCESS") {
      // Refresh token resulted with success. That means we can try to make request again.
      params.accessToken = refreshTokenResp.data.access_token;
      const newResult = (await func(params)) as RReturnType;

      // If this request resulted with success or failure return that.
      if (newResult.kind === "SUCCESS" || newResult.kind === "FAILURE") {
        return newResult;
      }

      // Otherwise, user need to login.
      return new Promise(() => {
        return {
          kind: "UNAUTHENTICATED",
          data: { message: "Unauthenticated" },
        };
      });
    } else {
      // Refresh token resulted with unauthenticated. That means user need to login.
      return new Promise(() => {
        return {
          kind: "UNAUTHENTICATED",
          data: { message: "Unauthenticated" },
        };
      });
    }
  }
}
