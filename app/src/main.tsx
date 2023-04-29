import React from "react";
import ReactDOM from "react-dom/client";
import { RouterProvider, createBrowserRouter } from "react-router-dom";
import { DashboardPage } from "./pages/dashboard.tsx";
import { UserDetailPage } from "./pages/users/detail.tsx";
import { UserEditPage } from "./pages/users/edit.tsx";
import { CompanyDetailPage } from "./pages/companies/detail.tsx";
import { CompanyEditPage } from "./pages/companies/edit.tsx";
import { LoginPage } from "./pages/login/login.tsx";

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
        path: "me/edit",
        element: <UserEditPage />,
      },
      {
        path: "company",
        element: <CompanyDetailPage />,
      },
      {
        path: "company/edit",
        element: <CompanyEditPage />,
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
