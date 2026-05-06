"use client";

import { FormEvent, useState } from "react";
import { useRouter } from "next/navigation";
import { Lock, Mail, User } from "lucide-react";
import { AuthCard } from "@/components/auth/auth-card";
import { AuthPageShell } from "@/components/auth/auth-page-shell";
import {
  buildUsername,
  persistAuthToken,
  registerWithPassword,
  splitFullName,
} from "@/lib/auth-client";

const fields = [
  {
    label: "Full name",
    name: "name",
    type: "text",
    placeholder: "Your name",
    icon: User,
  },
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
    placeholder: "Create a password",
    icon: Lock,
  },
];

export default function RegisterPage() {
  const router = useRouter();
  const [values, setValues] = useState({
    name: "",
    email: "",
    password: "",
  });
  const [agreedToTerms, setAgreedToTerms] = useState(false);
  const [error, setError] = useState("");
  const [isSubmitting, setIsSubmitting] = useState(false);

  const handleSubmit = async (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    setError("");

    if (!values.name.trim() || !values.email.trim() || !values.password) {
      setError("Complete your name, email, and password first.");
      return;
    }

    if (!agreedToTerms) {
      setError("Please agree to the terms before creating an account.");
      return;
    }

    const { firstName, lastName } = splitFullName(values.name);
    if (!firstName) {
      setError("Enter your full name.");
      return;
    }

    setIsSubmitting(true);
    try {
      const response = await registerWithPassword({
        email: values.email.trim(),
        password: values.password,
        username: buildUsername(values.name, values.email),
        firstName,
        lastName,
      });

      persistAuthToken(response.token);
      router.push("/communities");
    } catch (err) {
      setError(err instanceof Error ? err.message : "Registration failed.");
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <AuthPageShell>
      <AuthCard
        title="Register"
        description="Create your account and personalize your campus housing search."
        fields={fields}
        submitLabel="Create account"
        googleLabel="Sign up with Google"
        footerText="Already have an account?"
        footerLinkLabel="Login"
        footerHref="/login"
        values={values}
        onFieldChange={(name, value) => setValues((current) => ({ ...current, [name]: value }))}
        onSubmit={handleSubmit}
        error={error}
        isSubmitting={isSubmitting}
      >
        <label className="flex items-start gap-3 text-sm font-medium leading-relaxed text-gray-500">
          <input
            type="checkbox"
            checked={agreedToTerms}
            onChange={(event) => setAgreedToTerms(event.target.checked)}
            className="mt-1 h-5 w-5 rounded border-gray-300 accent-brand-orange"
          />
          I agree to the terms and want updates about student housing opportunities.
        </label>
      </AuthCard>
    </AuthPageShell>
  );
}
