import { ReactNode } from "react";

interface FieldProps {
  label: string;
  children: ReactNode;
}

export function Field({ label, children }: FieldProps) {
  return (
    <label className="block">
      <span className="mb-2 block text-sm font-bold text-gray-700">{label}</span>
      {children}
    </label>
  );
}

export const inputClassName =
  "w-full rounded-2xl border border-gray-100 bg-white px-5 py-4 text-base font-semibold text-gray-900 shadow-sm outline-none transition-all placeholder:text-gray-400 focus:ring-2 focus:ring-brand-orange/20";
