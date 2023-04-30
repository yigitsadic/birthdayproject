import { useState } from "react";
import { Employee } from "../../client/response_types/employee";
import { employeeUpdate } from "../../client/requests/employees/update";

interface IEmployeeListItemEdit {
  employee: Employee;
  accessToken: string;
  companyId: number;
  setMode: React.Dispatch<React.SetStateAction<"display" | "edit">>;
  setEmployee: React.Dispatch<React.SetStateAction<Employee>>;
}

export const EmployeeListItemEdit = (props: IEmployeeListItemEdit) => {
  const [firstName, setFirstName] = useState(props.employee.first_name);
  const [lastName, setLastName] = useState(props.employee.last_name);
  const [email, setEmail] = useState(props.employee.email);
  const [birthDay, setBirthDay] = useState(props.employee.birth_day);
  const [birthMonth, setBirthMonth] = useState(props.employee.birth_month);

  const [message, setMessage] = useState<string | null>(null);

  const handleSave = async () => {
    const result = await employeeUpdate({
      accessToken: props.accessToken,
      employee_id: props.employee.id,
      company_id: props.companyId,
      dto: {
        first_name: firstName,
        last_name: lastName,
        email: email,
        birth_day: birthDay,
        birth_month: birthMonth,
      },
    });

    if (result.kind === "SUCCESS") {
      props.setEmployee(result.data);
      props.setMode("display");
    }

    if (result.kind === "FAILURE") {
      setMessage(result.data.message);
    }
  };

  return (
    <>
      {message && (
        <tr>
          <td colSpan={5}>{message}</td>
        </tr>
      )}

      <tr key={props.employee.id}>
        <td>
          <input
            value={firstName}
            onChange={(e) => setFirstName(e.target.value)}
          />
        </td>

        <td>
          <input
            value={lastName}
            onChange={(e) => setLastName(e.target.value)}
          />
        </td>

        <td>
          <input value={email} onChange={(e) => setEmail(e.target.value)} />
        </td>

        <td>
          <input
            value={birthDay}
            onChange={(e) => setBirthDay(parseInt(e.target.value))}
          />

          <input
            value={birthMonth}
            onChange={(e) => setBirthMonth(parseInt(e.target.value))}
          />
        </td>

        <td>
          <button onClick={handleSave}>Save</button>
          <button onClick={() => props.setMode("display")}>Display</button>
        </td>
      </tr>
    </>
  );
};
