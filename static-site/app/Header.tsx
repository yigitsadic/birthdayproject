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
        </Link>
      </nav>
    </header>
  );
};
