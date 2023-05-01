import { useContext, useState } from "react";
import { AuthContext } from "../../store/auth_context";
import { sessionCreate } from "../../client/requests/sessions/create";
import { redirect } from "react-router-dom";

export const LoginPage = () => {
  const { setAuthStore } = useContext(AuthContext);

  const [message, setMessage] = useState<string | null>(null);
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault();

    const result = await sessionCreate({
      email,
      password,
    });

    if (result.kind === "SUCCESS") {
      setAuthStore({
        company_id: result.data.company_id,
        access_token: result.data.access_token,
        user_id: result.data.user_id,
      });

      console.log("Buraya geldi!");

      redirect("/");
    } else {
      setMessage(result.data.message);
    }
  };

  return (
    <div>
      <h3>login</h3>

      {message && <p>{message}</p>}

      <form onSubmit={handleLogin}>
        <div>
          <label>Email</label>

          <input
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />
        </div>

        <div>
          <label>Password</label>

          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
        </div>

        <div>
          <input type="submit" value="Login" />
        </div>
      </form>
    </div>
  );
};
