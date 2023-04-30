import React from "react";
import ReactDOM from "react-dom/client";
import { RouterProvider, createBrowserRouter } from "react-router-dom";
import { DashboardPage } from "./pages/dashboard.tsx";
import { UserDetailPage } from "./pages/users/detail.tsx";
import { CompanyDetailPage } from "./pages/companies/detail.tsx";
import { LoginPage } from "./pages/login/login.tsx";
import { EmployeesListPage } from "./pages/employees/list.tsx";

const router = createBrowserRouter([
  {
    path: "/",
    element: <DashboardPage />,
    children: [
      {
        path: "me",
        element: <UserDetailPage />,
      },
      {
        path: "company",
        element: <CompanyDetailPage />,
      },
      {
        path: "employees",
        element: <EmployeesListPage />,
      },
    ],
  },
  {
    path: "/login",
    element: <LoginPage />,
  },
]);

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);
