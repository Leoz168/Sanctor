"use client";

import { useState } from "react";

interface SegmentedControlProps {
  label: string;
  options: string[];
  defaultValue: string;
}

export function SegmentedControl({ label, options, defaultValue }: SegmentedControlProps) {
  const [selected, setSelected] = useState(defaultValue);

  return (
    <div>
      <p className="mb-2 text-xl font-bold text-gray-700">{label}</p>
      <div className="grid rounded-2xl border border-gray-100 bg-white p-1 shadow-sm" style={{ gridTemplateColumns: `repeat(${options.length}, minmax(0, 1fr))` }}>
        {options.map((option) => (
          <button
            key={option}
            type="button"
            onClick={() => setSelected(option)}
            className={`rounded-xl px-4 py-3 text-sm font-bold uppercase transition-all ${
              selected === option
                ? "bg-brand-orange text-white shadow-lg shadow-brand-orange/20"
                : "text-gray-400 hover:text-brand-orange"
            }`}
          >
            {option}
          </button>
        ))}
      </div>
    </div>
  );
}
