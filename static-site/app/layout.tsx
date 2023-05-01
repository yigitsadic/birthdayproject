import { Footer } from "./Footer";
import { Header } from "./Header";
import "./product.css";

export const metadata = {
  title: "Cakeday Today",
  description:
    "Celebrate Your Employees' Birthdays with Personalized Emails. Make every birthday special and show your employees how much you appreciate them with our automated birthday email service. Sign up today and let us help you create a more positive and engaging workplace culture.",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en" data-bs-theme="dark">
      <head>
        <link
          href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0-alpha3/dist/css/bootstrap.min.css"
          rel="stylesheet"
          integrity="sha384-KK94CHFLLe+nY2dmCWGMq91rCGa5gtU4mk92HdvYe+M/SXH301p5ILy+dN9+nJOZ"
          crossOrigin="anonymous"
        />
        <script
          src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0-alpha3/dist/js/bootstrap.bundle.min.js"
          integrity="sha384-ENjdO4Dr2bkBIFxQpeoTz1HIcje39Wm4jDKdf19U8gI4ddQ3GYNS7NTKfAdVQSZe"
          crossOrigin="anonymous"
        ></script>
      </head>

      <body>
        <Header />

        <main>{children}</main>

        <Footer />
      </body>
    </html>
  );
}
