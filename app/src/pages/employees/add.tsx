import { useState } from "react";
import { createEmployee } from "../../client/requests/employees/create";
import { Employee } from "../../client/response_types/employee";

interface IAddFormProps {
  accessToken: string;
  companyId: number;
  addEmployee: (employee: Employee) => void;
}

export const AddForm = (props: IAddFormProps) => {
  const [firstName, setFirstName] = useState("");
  const [lastName, setLastName] = useState("");
  const [email, setEmail] = useState("");
  const [birthDay, setBirthDay] = useState(15);
  const [birthMonth, setBirthMonth] = useState(4);

  const [message, setMessage] = useState<string | null>(null);

  const handleSave = async (e: React.FormEvent) => {
    e.preventDefault();

    const result = await createEmployee({
      accessToken: props.accessToken,
      company_id: props.companyId,
      dto: {
        first_name: firstName,
        last_name: lastName,
        email: email,
        birth_day: birthDay,
        birth_month: birthMonth,
      },
    });

    if (result.kind === "FAILURE") {
      setMessage(result.data.message);
    }

    if (result.kind === "SUCCESS") {
      setFirstName("");
      setLastName("");
      setEmail("");

      props.addEmployee(result.data);
    }
  };

  return (
    <div>
      {message && message}

      <form onSubmit={handleSave}>
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
          <label>Email</label>
          <input value={email} onChange={(e) => setEmail(e.target.value)} />
        </div>

        <div>
          <label>Birth Day</label>
          <input
            value={birthDay}
            onChange={(e) => setBirthDay(parseInt(e.target.value))}
          />
        </div>

        <div>
          <label>Birth Month</label>
          <input
            value={birthMonth}
            onChange={(e) => setBirthMonth(parseInt(e.target.value))}
          />
        </div>

        <div>
          <input value="Save" type="submit" />
        </div>
      </form>
    </div>
  );
};
