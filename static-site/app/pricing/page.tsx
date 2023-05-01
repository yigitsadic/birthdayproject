import Link from "next/link";

export default function PricingPage() {
  return (
    <div className="container">
      <h2 className="display-3 text-center">Pricing</h2>

      <p className="fs-3">
        At Cakeday.today, we offer a range of pricing plans to suit businesses
        of all sizes. Our plans are designed to be flexible, affordable, and
        scalable, so you can choose the option that best fits your needs and
        budget.
      </p>

      <p className="fs-5">
        Our Small plan is completely free for businesses with up to 10
        employees. This plan includes all the basic features of our service,
        including automated birthday emails and customizable templates.
      </p>
      <p className="fs-5">
        For businesses with 10 to 25 employees, we offer our Medium plan, which
        is priced at just $5 per month. This plan includes all the features of
        our Small plan, as well as weekly reports on upcoming birthdays and the
        ability to customize your email templates.
      </p>
      <p className="fs-5">
        For larger businesses with more than 25 employees, we offer our
        Enterprise plan, which is priced at $19.99 per month. This plan includes
        all the features of our Medium plan, as well as advanced customization
        options, priority support, and dedicated account management.
      </p>
      <p className="fs-5">
        We understand that every business is unique, and we're happy to work
        with you to find a pricing plan that meets your specific needs. If you
        have special requirements or would like to discuss custom pricing,
        please don't hesitate to contact us.
      </p>
      <p className="fs-5">
        At Cakeday.today, our goal is to help businesses create a more positive
        and engaging workplace culture by celebrating their employees'
        birthdays. Choose the plan that's right for you and start showing your
        employees how much you value them today!
      </p>
      <p className="fs-5 text-center">
        <Link href={"/request-a-demo"} className="btn btn-success btn-lg">
          Request a demo
        </Link>
      </p>
    </div>
  );
}
