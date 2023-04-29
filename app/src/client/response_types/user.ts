import { UserDto } from "../dtos/user_dto";

export interface User extends UserDto {
  id: number;
  email: string;
}
