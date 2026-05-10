import { Mail } from "lucide-react";
import { FieldLabel } from "@/components/profile/profile-field";

interface ProfileContactCardProps {
  email: string;
  onChange: (value: string) => void;
}

export function ProfileContactCard({ email, onChange }: ProfileContactCardProps) {
  return (
    <label className="block">
      <FieldLabel>Contact email</FieldLabel>
      <div className="flex h-12 items-center gap-3 rounded-2xl border border-orange-100/80 bg-brand-cream/80 px-4 shadow-inner shadow-orange-900/5">
        <div className="flex h-9 w-9 shrink-0 items-center justify-center rounded-xl bg-white text-brand-orange shadow-sm">
          <Mail className="h-4 w-4" />
        </div>

        <input
          type="email"
          value={email}
          onChange={(event) => onChange(event.target.value)}
          className="min-w-0 flex-1 bg-transparent text-sm font-bold text-gray-900 outline-none placeholder:text-gray-400"
          placeholder="name@school.edu"
        />
      </div>
    </label>
  );
}
