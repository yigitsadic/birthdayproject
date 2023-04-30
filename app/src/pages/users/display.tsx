import { User } from "../../client/response_types/user";

interface UserDisplayProps {
  user: User;
  accessToken: string;
  setMode: React.Dispatch<React.SetStateAction<"display" | "edit">>;
}

export const UserDisplay = (props: UserDisplayProps) => {
  return (
    <div>
      <table>
        <tbody>
          <tr>
            <th>User ID</th>
            <td>{props.user.id}</td>
          </tr>
          <tr>
            <th>First Name</th>
            <td>{props.user.first_name}</td>
          </tr>
          <tr>
            <th>Last Name</th>
            <td>{props.user.last_name}</td>
          </tr>
          <tr>
            <th>Email</th>
            <td>{props.user.email}</td>
          </tr>
        </tbody>
      </table>

      <button onClick={() => props.setMode("edit")}>Edit</button>
    </div>
  );
};
