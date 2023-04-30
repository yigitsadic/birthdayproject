import { useState } from "react";
import { Employee } from "../../client/response_types/employee";
import { EmployeeListItemDisplay } from "./display";
import { EmployeeListItemEdit } from "./edit";

interface EmployeeListItemInterface {
  employee: Employee;
  accessToken: string;
  companyId: number;
  removeFromList: (id: number) => void;
}

export const EmployeeListItemWrapper = (props: EmployeeListItemInterface) => {
  const [mode, setMode] = useState<"display" | "edit">("display");
  const [employee, setEmployee] = useState(props.employee);

  if (mode === "edit") {
    return (
      <EmployeeListItemEdit
        key={employee.id}
        employee={employee}
        setMode={setMode}
        accessToken={props.accessToken}
        companyId={props.companyId}
        setEmployee={setEmployee}
      />
    );
  }

  return (
    <EmployeeListItemDisplay
      removeFromList={props.removeFromList}
      key={employee.id}
      employee={employee}
      setMode={setMode}
      accessToken={props.accessToken}
      companyId={props.companyId}
    />
  );
};
