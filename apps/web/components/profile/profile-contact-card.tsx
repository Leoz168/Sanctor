import { Mail } from "lucide-react";

export function ProfileContactCard() {
  return (
    <div className="flex items-center gap-4 rounded-3xl border border-orange-100 bg-brand-cream px-5 py-4">
      <div className="flex h-11 w-11 shrink-0 items-center justify-center rounded-2xl bg-white text-brand-orange shadow-sm">
        <Mail className="h-5 w-5" />
      </div>

      <div className="min-w-0">
        <p className="text-[11px] font-bold uppercase tracking-[0.22em] text-gray-400">
          Contact email
        </p>
        <p className="mt-1 text-lg font-bold text-gray-900">
          alex.rivera@school.edu
        </p>
      </div>
    </div>
  );
}
