"use client";

import { useState } from "react";

interface SegmentedControlProps {
  label: string;
  options: string[];
  defaultValue: string;
  value?: string;
  onChange?: (value: string) => void;
  size?: "md" | "sm";
}

export function SegmentedControl({
  label,
  options,
  defaultValue,
  value,
  onChange,
  size = "md",
}: SegmentedControlProps) {
  const [internalSelected, setInternalSelected] = useState(defaultValue);
  const selected = value ?? internalSelected;
  const setSelected = (nextValue: string) => {
    if (value === undefined) {
      setInternalSelected(nextValue);
    }
    onChange?.(nextValue);
  };
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
