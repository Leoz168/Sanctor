import { Mail } from "lucide-react";

export function ProfileContactCard() {
  return (
    <div className="flex items-center gap-3 rounded-2xl border border-orange-100 bg-brand-cream px-4 py-3">
      <div className="flex h-9 w-9 shrink-0 items-center justify-center rounded-xl bg-white text-brand-orange shadow-sm">
        <Mail className="h-4 w-4" />
      </div>

      <div className="min-w-0">
        <p className="text-[10px] font-bold uppercase tracking-[0.2em] text-gray-400">
          Contact email
        </p>
        <p className="mt-0.5 text-base font-bold text-gray-900">
          alex.rivera@school.edu
        </p>
      </div>
    </div>
  );
}
