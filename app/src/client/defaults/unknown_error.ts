import { ErrorMessage } from "../response_types/error_message";

export const UnknownError: {
  kind: "FAILURE";
  data: ErrorMessage;
} = {
  kind: "FAILURE",
  data: {
    message: "Unknown error occurred",
  },
};
