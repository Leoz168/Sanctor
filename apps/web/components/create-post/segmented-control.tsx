"use client";

import { useState } from "react";

interface SegmentedControlProps {
  label: string;
  options: string[];
  defaultValue: string;
  size?: "md" | "sm";
}

export function SegmentedControl({
  label,
  options,
  defaultValue,
  size = "md",
}: SegmentedControlProps) {
  const [selected, setSelected] = useState(defaultValue);
  const isSmall = size === "sm";

  return (
    <div>
      <p
        className={
          isSmall
            ? "mb-2 text-sm font-bold uppercase tracking-[0.16em] text-gray-400"
            : "mb-2 text-xl font-bold text-gray-700"
        }
      >
        {label}
      </p>
      <div
        className={`grid border border-gray-100 bg-white p-1 shadow-sm ${
          isSmall ? "rounded-xl" : "rounded-2xl"
        }`}
        style={{ gridTemplateColumns: `repeat(${options.length}, minmax(0, 1fr))` }}
      >
        {options.map((option) => (
          <button
            key={option}
            type="button"
            onClick={() => setSelected(option)}
            className={`font-bold uppercase transition-all ${
              selected === option
                ? "bg-brand-orange text-white shadow-lg shadow-brand-orange/20"
                : "text-gray-400 hover:text-brand-orange"
            } ${isSmall ? "rounded-lg px-3 py-2.5 text-xs" : "rounded-xl px-4 py-3 text-sm"}`}
          >
            {option}
          </button>
        ))}
      </div>
    </div>
  );
}
