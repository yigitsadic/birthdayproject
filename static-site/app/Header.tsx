import Link from "next/link";

export const Header = () => {
  return (
    <header className="site-header sticky-top py-1">
      <nav className="container d-flex flex-column flex-md-row justify-content-between">
        <a className="py-2" href="#" aria-label="Product">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="24"
            height="24"
            fill="none"
            stroke="currentColor"
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            className="d-block mx-auto"
            role="img"
            viewBox="0 0 24 24"
          >
            <title>cakeday.today</title>
            <circle cx="12" cy="12" r="10" />
            <path d="M14.31 8l5.74 9.94M9.69 8h11.48M7.38 12l5.74-9.94M9.69 16L3.95 6.06M14.31 16H2.83m13.79-4l-5.74 9.94" />
          </svg>
        </a>

        <Link className="py-2 d-none d-md-inline-block" href={"/"}>
          Home
        </Link>

        <Link
          className="py-2 d-none d-md-inline-block"
          href={"/request-a-demo"}
        >
          Request a demo
        </Link>

        <Link className="py-2 d-none d-md-inline-block" href={"/pricing"}>
          Pricing
        </Link>

        <Link className="py-2 d-none d-md-inline-block" href={"/about"}>
          About
        </Link>

        <Link className="py-2 d-none d-md-inline-block" href={"/app"}>
          App Dashboard
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="16"
            height="16"
            fill="currentColor"
            className="bi bi-arrow-up-right-square"
            viewBox="0 0 16 16"
          >
            <path
              fill-rule="evenodd"
              d="M15 2a1 1 0 0 0-1-1H2a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V2zM0 2a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V2zm5.854 8.803a.5.5 0 1 1-.708-.707L9.243 6H6.475a.5.5 0 1 1 0-1h3.975a.5.5 0 0 1 .5.5v3.975a.5.5 0 1 1-1 0V6.707l-4.096 4.096z"
            />
          </svg>
        </Link>
      </nav>
    </header>
  );
};
