import { UserDto } from "../input_types/user_dto";

export interface User extends UserDto {
  id: number;
  email: string;
}
