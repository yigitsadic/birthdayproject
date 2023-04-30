import { Company } from "../../client/response_types/company";

interface CompanyDisplayParams {
  company: Company;
  setMode: React.Dispatch<React.SetStateAction<"display" | "edit">>;
}

export const CompanyDisplay = (props: CompanyDisplayParams) => {
  return (
    <div>
      <table>
        <tbody>
          <tr>
            <th>Company ID</th>
            <td>{props.company.id}</td>
          </tr>
          <tr>
            <th>Name</th>
            <td>{props.company.name}</td>
          </tr>
        </tbody>
      </table>

      <button onClick={() => props.setMode("edit")}>Edit</button>
    </div>
  );
};
