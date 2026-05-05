import { Mail } from "lucide-react";
import { FieldLabel } from "@/components/profile/profile-field";

export function ProfileContactCard() {
  return (
    <div>
      <FieldLabel>Contact email (optional)</FieldLabel>
      <div className="flex h-12 items-center gap-3 rounded-2xl border border-orange-100/80 bg-brand-cream/80 px-4 shadow-inner shadow-orange-900/5">
        <div className="flex h-9 w-9 shrink-0 items-center justify-center rounded-xl bg-white text-brand-orange shadow-sm">
          <Mail className="h-4 w-4" />
        </div>

        <p className="min-w-0 truncate text-sm font-bold text-gray-900">
          alex.rivera@school.edu
        </p>
      </div>
    </div>
  );
}
