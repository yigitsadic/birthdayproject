import { useState } from "react";
import { User } from "../../client/response_types/user";
import { updateUser } from "../../client/requests/users/update";

interface UserFormProps {
  user: User;
  accessToken: string;
  setMode: React.Dispatch<React.SetStateAction<"display" | "edit">>;
  setUser: React.Dispatch<React.SetStateAction<User | null>>;
}

export const UserForm = (props: UserFormProps) => {
  const [message, setMessage] = useState<string | null>(null);
  const [firstName, setFirstName] = useState(props.user.first_name);
  const [lastName, setLastName] = useState(props.user.last_name);

  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault();

    const resp = await updateUser({
      accessToken: props.accessToken,
      user_id: props.user.id,

      dto: {
        first_name: firstName,
        last_name: lastName,
      },
    });

    if (resp.kind === "FAILURE") {
      setMessage(resp.data.message);
    }

    if (resp.kind === "SUCCESS") {
      props.setUser({
        id: props.user.id,
        first_name: firstName,
        last_name: lastName,
        email: props.user.email,
      });
    }

    props.setMode("display");
  };

  return (
    <div>
      {message && <p>{message}</p>}

      <form onSubmit={handleSubmit}>
        <div>
          <label>First Name</label>
          <input
            value={firstName}
            onChange={(e) => setFirstName(e.target.value)}
          />
        </div>

        <div>
          <label>Last Name</label>
          <input
            value={lastName}
            onChange={(e) => setLastName(e.target.value)}
          />
        </div>

        <div>
          <input type="submit" value="Save" />
        </div>
      </form>

      <button onClick={() => props.setMode("display")}>Display</button>
    </div>
  );
};
