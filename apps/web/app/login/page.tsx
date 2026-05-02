import { Lock, Mail } from "lucide-react";
import { AuthCard } from "@/components/auth/auth-card";
import { AuthPageShell } from "@/components/auth/auth-page-shell";

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
