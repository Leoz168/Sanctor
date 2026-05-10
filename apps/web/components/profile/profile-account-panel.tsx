import { ShieldAlert, Trash2 } from "lucide-react";
import { profilePanelClassName } from "@/components/profile/profile-panel-styles";

interface ProfileAccountPanelProps {
  email: string;
  username: string;
  isDeleting?: boolean;
  errorMessage?: string;
  onDelete: () => void;
}

export function ProfileAccountPanel({
  email,
  username,
  isDeleting = false,
  errorMessage,
  onDelete,
}: ProfileAccountPanelProps) {
  return (
    <section className={`${profilePanelClassName} overflow-y-auto`}>
      <div className="flex h-14 w-14 items-center justify-center rounded-2xl bg-brand-cream text-brand-orange">
        <ShieldAlert className="h-6 w-6" />
      </div>

      <h1 className="mt-6 text-3xl font-bold tracking-tight text-gray-900">
        Account management
      </h1>
      <p className="mt-3 max-w-2xl text-base font-semibold leading-7 text-gray-500">
        You are signed in as {username || "this user"} with {email || "your account"}.
        Deleting your account is permanent and will remove your profile from Sanctor.
      </p>

      {errorMessage ? (
        <p className="mt-5 rounded-2xl border border-red-200 bg-red-50 px-4 py-3 text-sm font-semibold text-red-700">
          {errorMessage}
        </p>
      ) : null}

      <div className="mt-8 rounded-[1.75rem] border border-red-200 bg-red-50/80 p-6">
        <div className="flex items-start gap-4">
          <div className="flex h-12 w-12 shrink-0 items-center justify-center rounded-2xl bg-white text-red-600 shadow-sm">
            <Trash2 className="h-5 w-5" />
          </div>

          <div className="min-w-0">
            <h2 className="text-lg font-bold text-gray-900">Delete account</h2>
            <p className="mt-2 text-sm font-semibold leading-6 text-gray-600">
              This permanently deletes your user profile. You will be signed out immediately
              and redirected to the login page.
            </p>
            <button
              type="button"
              onClick={onDelete}
              disabled={isDeleting}
              className="mt-5 rounded-2xl bg-red-600 px-6 py-3 text-xs font-bold uppercase tracking-[0.16em] text-white shadow-lg shadow-red-600/25 transition-all hover:-translate-y-0.5 hover:bg-red-700 disabled:cursor-not-allowed disabled:opacity-70"
            >
              {isDeleting ? "Deleting..." : "Delete account"}
            </button>
          </div>
        </div>
      </div>
    </section>
  );
}
