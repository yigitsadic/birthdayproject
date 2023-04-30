import { useContext, useEffect, useState } from "react";
import { AuthenticationRequired } from "../general/authentication_required";
import { Company } from "../../client/response_types/company";
import { companyDetail } from "../../client/requests/companies/detail";
import { AuthContext } from "../../store/auth_context";
import { CompanyDisplay } from "./display";
import { CompanyForm } from "./form";

export const CompanyDetailPage = () => {
  const { authStore } = useContext(AuthContext);

  if (!authStore) {
    return <AuthenticationRequired />;
  }

  const [company, setCompany] = useState<Company | null>(null);
  const [mode, setMode] = useState<"display" | "edit">("display");

  const fetchFromAPI = async () => {
    const result = await companyDetail({
      accessToken: authStore.access_token,
      company_id: authStore.company_id,
    });

    if (result.kind === "SUCCESS") {
      setCompany(result.data);
    }
  };

  useEffect(() => {
    fetchFromAPI();
  }, [authStore.access_token]);

  if (company) {
    return (
      <div>
        <h3>Company Information</h3>
        {mode === "display" ? (
          <CompanyDisplay company={company} setMode={setMode} />
        ) : (
          <CompanyForm
            company={company}
            accessToken={authStore.access_token}
            setCompany={setCompany}
            setMode={setMode}
          />
        )}
      </div>
    );
  }

  return <div>No content found.</div>;
};
