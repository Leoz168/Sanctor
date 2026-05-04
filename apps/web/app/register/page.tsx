import { Lock, Mail, User } from "lucide-react";
import { AuthCard } from "@/components/auth/auth-card";
import { AuthPageShell } from "@/components/auth/auth-page-shell";

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
      >
        <label className="flex items-start gap-3 text-sm font-medium leading-relaxed text-gray-500">
          <input type="checkbox" className="mt-1 h-5 w-5 rounded border-gray-300 accent-brand-orange" />
          I agree to the terms and want updates about student housing opportunities.
        </label>
      </AuthCard>
    </AuthPageShell>
  );
}
