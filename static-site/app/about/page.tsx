import Link from "next/link";

export default function AboutPage() {
  return (
    <div className="container">
      <h2 className="display-3 text-center">About us</h2>
      <p className="fs-3">
        Welcome to Cakeday.today, a service that helps companies celebrate their
        employees' birthdays with personalized email greetings.
      </p>

      <p className="fs-5">
        At Cakeday.today, we understand the importance of recognizing and
        appreciating your employees. Birthdays are a special occasion that
        deserve to be celebrated, and our service makes it easy for you to do
        just that.
      </p>

      <p className="fs-5">
        Our automated birthday email service allows you to add your employees'
        names, email addresses, and birthdays to our system. We'll take care of
        the rest, sending a personalized birthday greeting to each employee on
        their special day.
      </p>

      <p className="fs-5">
        With our service, you can customize the email greeting with your own
        company logo, message, and branding. You can also choose from a variety
        of pre-designed templates or create your own unique design.
      </p>

      <p className="fs-5">
        At Cakeday.today, we're committed to making your employees feel valued
        and appreciated. Our service is easy to use, affordable, and reliable.
        We take care of the technical details so you can focus on what really
        matters - creating a positive and engaging workplace culture.
      </p>

      <p className="lead">
        Join us today and start celebrating your employees' birthdays in style!
      </p>

      <p className="fs-5 text-center">
        <Link href={"/request-a-demo"} className="btn btn-success btn-lg">
          Request a demo
        </Link>
      </p>
    </div>
  );
}
