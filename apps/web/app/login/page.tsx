"use client";

import { FormEvent, useState } from "react";
import { useRouter } from "next/navigation";
import { Lock, Mail } from "lucide-react";
import { AuthCard } from "@/components/auth/auth-card";
import { AuthPageShell } from "@/components/auth/auth-page-shell";
import { loginWithPassword, persistAuthToken } from "@/lib/auth-client";

const fields = [
  {
    label: "Email",
    name: "email",
    type: "email",
    placeholder: "you@school.edu",
    icon: Mail,
  },
  {
    label: "Password",
    name: "password",
    type: "password",
    placeholder: "Enter your password",
    icon: Lock,
  },
];

export default function LoginPage() {
  const router = useRouter();
  const [values, setValues] = useState({
    email: "",
    password: "",
  });
  const [error, setError] = useState("");
  const [isSubmitting, setIsSubmitting] = useState(false);

  const handleSubmit = async (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    setError("");

    if (!values.email.trim() || !values.password) {
      setError("Enter both your email and password.");
      return;
    }

    setIsSubmitting(true);
    try {
      const response = await loginWithPassword({
        email: values.email.trim(),
        password: values.password,
      });

      persistAuthToken(response.token);
      router.push("/communities");
    } catch (err) {
      setError(err instanceof Error ? err.message : "Login failed.");
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <AuthPageShell>
      <AuthCard
        title="Login"
        description="Use your account to manage saved homes and community posts."
        fields={fields}
        submitLabel="Login"
        googleLabel="Sign in with Google"
        footerText="New to Rentling?"
        footerLinkLabel="Create an account"
        footerHref="/register"
        values={values}
        onFieldChange={(name, value) => setValues((current) => ({ ...current, [name]: value }))}
        onSubmit={handleSubmit}
        error={error}
        isSubmitting={isSubmitting}
      >
        <div className="flex items-center justify-between gap-4 text-sm font-semibold sm:text-base">
          <label className="flex items-center gap-3 text-gray-500">
            <input type="checkbox" className="h-5 w-5 rounded border-gray-300 accent-brand-orange" />
            Remember me
          </label>
          <a href="#" className="text-brand-orange hover:text-orange-600">
            Forgot password?
          </a>
        </div>
      </AuthCard>
    </AuthPageShell>
  );
}
