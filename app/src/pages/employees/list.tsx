import { useContext, useEffect, useState } from "react";
import { AuthContext } from "../../store/auth_context";
import { AuthenticationRequired } from "../general/authentication_required";
import { Employee } from "../../client/response_types/employee";
import { employeesList } from "../../client/requests/employees/list";
import { EmployeeListItemWrapper } from "./listItemWrapper";
import { AddForm } from "./add";

export const EmployeesListPage = () => {
  const { authStore } = useContext(AuthContext);

  if (!authStore) {
    return <AuthenticationRequired />;
  }

  const [employees, setEmployees] = useState<Employee[]>([]);
  const [showNewForm, setShowNewForm] = useState(false);

  const fetchFromAPI = async () => {
    const result = await employeesList({
      accessToken: authStore.access_token,
      company_id: authStore.company_id,
    });

    if (result.kind === "SUCCESS") {
      setEmployees(result.data);
    }
  };

  const addEmployee = (employee: Employee) => {
    if (employees) {
      const newEmloyees = [employee, ...employees];
      setEmployees(newEmloyees);
    } else {
      setEmployees([employee]);
    }
  };

  const removeFromList = (id: number) => {
    const newEmployees = employees.filter((em) => em.id !== id);

    setEmployees(newEmployees);
  };

  useEffect(() => {
    fetchFromAPI();
  }, [authStore.access_token]);

  return (
    <div>
      <h3>Employees</h3>

      {showNewForm ? (
        <div>
          <button onClick={() => setShowNewForm(false)}>Close Form</button>

          <AddForm
            addEmployee={addEmployee}
            accessToken={authStore.access_token}
            companyId={authStore.company_id}
          />
        </div>
      ) : (
        <div>
          <button onClick={() => setShowNewForm(true)}>Add New</button>
        </div>
      )}

      <table className="table">
        <thead>
          <tr>
            <th>First Name</th>
            <th>Last Name</th>
            <th>Email</th>
            <th>Birthday</th>
            <th>Actions</th>
          </tr>
        </thead>

        <tbody>
          {employees &&
            employees.map((employee) => {
              return (
                <EmployeeListItemWrapper
                  removeFromList={removeFromList}
                  key={employee.id}
                  employee={employee}
                  companyId={authStore.company_id}
                  accessToken={authStore.access_token}
                />
              );
            })}
        </tbody>
      </table>
    </div>
  );
};
