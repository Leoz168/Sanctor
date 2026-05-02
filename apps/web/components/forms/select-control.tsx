"use client";

import { ChevronDown } from "lucide-react";

interface SelectControlProps {
  label: string;
  options: string[];
  value?: string;
  onChange?: (value: string) => void;
  className?: string;
  variant?: "default" | "panel" | "hero" | "create";
}

export function SelectControl({
  label,
  options,
  value,
  onChange,
  className = "",
  variant = "default",
}: SelectControlProps) {
  const selectClassName =
    variant === "panel"
      ? "w-full appearance-none bg-white border border-gray-100 rounded-xl px-4 py-3 pr-10 text-sm font-bold text-gray-700 cursor-pointer hover:bg-brand-cream transition-colors focus:outline-none focus:ring-2 focus:ring-brand-orange/20"
      : variant === "hero"
        ? "w-full appearance-none bg-brand-cream border border-gray-100 rounded-2xl px-6 py-5 pr-12 text-base font-bold text-gray-700 cursor-pointer hover:bg-white transition-colors focus:outline-none focus:ring-2 focus:ring-brand-orange/20"
        : variant === "create"
          ? "w-full appearance-none rounded-2xl border border-gray-100 bg-white px-5 py-4 pr-12 text-base font-semibold text-gray-900 shadow-sm outline-none transition-all focus:ring-2 focus:ring-brand-orange/20"
          : "w-full appearance-none bg-brand-cream border border-gray-100 rounded-2xl px-5 py-4 pr-10 text-sm font-bold text-gray-700 cursor-pointer hover:bg-white transition-colors focus:outline-none focus:ring-2 focus:ring-brand-orange/20";

  return (
    <label className={`relative ${className}`}>
      <span className="sr-only">{label}</span>
      <select
        value={value}
        onChange={(event) => onChange?.(event.target.value)}
        className={selectClassName}
      >
        {options.map((option) => (
          <option key={option}>{option}</option>
        ))}
      </select>
      <ChevronDown className="absolute right-4 top-1/2 -translate-y-1/2 text-gray-400 pointer-events-none" size={16} />
    </label>
  );
}
