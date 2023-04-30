import { employeeDestroy } from "../../client/requests/employees/destroy";
import { Employee } from "../../client/response_types/employee";

interface IEmployeeListItemDisplay {
  employee: Employee;
  accessToken: string;
  companyId: number;
  setMode: React.Dispatch<React.SetStateAction<"display" | "edit">>;
  removeFromList: (id: number) => void;
}

export const EmployeeListItemDisplay = (props: IEmployeeListItemDisplay) => {
  const handleDelete = async (id: number) => {
    await employeeDestroy({
      company_id: props.companyId,
      accessToken: props.accessToken,
      employee_id: id,
    });

    props.removeFromList(id);
  };

  return (
    <tr key={props.employee.id}>
      <td>{props.employee.first_name}</td>
      <td>{props.employee.last_name}</td>
      <td>{props.employee.email}</td>
      <td>
        {props.employee.birth_day} / {props.employee.birth_month}
      </td>
      <td>
        <button onClick={() => props.setMode("edit")}>Edit</button>
        <button onClick={() => handleDelete(props.employee.id)}>Delete</button>
      </td>
    </tr>
  );
};
