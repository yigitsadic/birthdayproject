import React from "react";
import ReactDOM from "react-dom/client";
import { RouterProvider, createBrowserRouter } from "react-router-dom";
import { DashboardPage } from "./pages/dashboard.tsx";
import { UserDetailPage } from "./pages/users/detail.tsx";
import { UserEditPage } from "./pages/users/edit.tsx";
import { CompanyDetailPage } from "./pages/companies/detail.tsx";
import { CompanyEditPage } from "./pages/companies/edit.tsx";

const router = createBrowserRouter([
  {
    path: "/",
    element: <DashboardPage />,
    children: [
      {
        path: "users/me",
        element: <UserDetailPage />,
      },
      {
        path: "users/me/edit",
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
]);

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);
