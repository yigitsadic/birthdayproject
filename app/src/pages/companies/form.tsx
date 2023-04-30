import { useState } from "react";
import { updateCompany } from "../../client/requests/companies/update";
import { Company } from "../../client/response_types/company";

interface CompanyFormParams {
  company: Company;
  accessToken: string;
  setMode: React.Dispatch<React.SetStateAction<"display" | "edit">>;
  setCompany: React.Dispatch<React.SetStateAction<Company | null>>;
}

export const CompanyForm = (props: CompanyFormParams) => {
  const [name, setName] = useState(props.company.name);
  const [message, setMessage] = useState<string | null>(null);

  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault();

    const resp = await updateCompany({
      accessToken: props.accessToken,

      company_id: props.company.id,
      dto: {
        name: name,
      },
    });

    if (resp.kind === "FAILURE") {
      setMessage(resp.data.message);
    }

    if (resp.kind === "SUCCESS") {
      props.setCompany({
        name: resp.data.name,
        id: resp.data.id,
      });

      props.setMode("display");
    }
  };

  return (
    <div>
      {message && message}

      <form onSubmit={handleSubmit}>
        <input value={name} onChange={(e) => setName(e.target.value)} />

        <input type="submit" value="Save" />
      </form>

      <button onClick={() => props.setMode("display")}>Display</button>
    </div>
  );
};
